package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"mmm/auth"
	"mmm/middleware"
	"mmm/util"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/xid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type pagination struct {
	Query      *gorm.DB    `json:"-"`
	Items      interface{} `json:"items"`
	TotalItems int64       `json:"totalItems"`
	Size       int64       `json:"size"`
	Page       int64       `json:"page"`
	UrlQuery   string      `json:"urlQuery"`
	TotalPages int64       `json:"totalPages"`
}

//接受 string int 两种传入值
func (p *pagination) Paginate() *gorm.DB {
	p.Query.Count(&p.TotalItems)
	if p.TotalItems == 0 {
		return p.Query
	}
	p.TotalPages = int64(math.Ceil(float64(p.TotalItems) / float64(p.Size)))
	query := p.Query.Offset(int((p.Page - 1) * p.Size)).Limit(int(p.Size))
	return query
}

func NewPagination(page, size int64) *pagination {
	return &pagination{Page: page, Size: size}
}

type User struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index,omitempty"`
	Name      string         `json:"name"`
	Intro     string         `json:"intro"`
}

type Dummy struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index,omitempty"`
	Name      string         `json:"name"`
	Intro     string         `json:"intro"`
}

type APIResponse struct {
	ErrorCode    int         `json:"error_code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()

	// Migrate the schema
	db.AutoMigrate(&User{}, &Dummy{})
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	JWTKey = []byte("jwtsecret")
	db     *gorm.DB
)

func main() {
	engine := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true
	engine.Use(cors.New(corsConfig))

	authMiddleware := middleware.JWT(JWTKey, &auth.JWTClaims{}, func(c *gin.Context, err error) interface{} {
		resp := &APIResponse{}
		if err != nil {
			resp.ErrorCode = 401
			resp.ErrorMessage = err.Error()
		}
		return resp
	}, func(c *gin.Context, token *jwt.Token) {
		if claims, ok := token.Claims.(*auth.JWTClaims); ok {
			// ... compare with ....
			c.Set("identity", claims.Audience)
		} else {
			resp := &APIResponse{}
			resp.ErrorCode = 401
			resp.ErrorMessage = "invalid jwt claims"
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
	},
	)

	authGroup := engine.Group("/")

	authGroup.Use(authMiddleware)

	engine.GET("/test", authMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": "ok",
		})
	})

	engine.POST("/auth", func(c *gin.Context) {
		// persiter jwt for the user, also check with jwt.id(unique and exsited)
		up := &LoginPayload{}

		resp := &APIResponse{}

		err := json.NewDecoder(c.Request.Body).Decode(up)
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer c.Request.Body.Close()

		user := &User{}
		err = db.Model(user).Where("name = ? and intro = ?", up.Username, up.Password).First(&user).Error
		if err != nil {
			resp.ErrorCode = 2
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}

		claims := auth.NewJWTClaims(fmt.Sprintf("%d", user.ID), xid.New().String())
		token, err := util.GenerateJWT(JWTKey, jwt.SigningMethodHS256, claims)
		if err != nil {
			resp.ErrorCode = 3
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = gin.H{
			"token":      token,
			"expired_at": time.Now().Add(time.Hour * 72).Unix(),
		}

		c.JSON(http.StatusOK, resp)
	})

	authGroup.POST("/users", func(c *gin.Context) {
		resp := &APIResponse{}

		user := &User{}
		err := json.NewDecoder(c.Request.Body).Decode(user)
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer c.Request.Body.Close()
		err = db.Model(user).Create(user).Error
		if err != nil {
			resp.ErrorCode = 2
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		c.JSON(http.StatusOK, user)
	})

	authGroup.GET("/users", func(c *gin.Context) {
		resp := &APIResponse{}
		users := make([]User, 0)

		sizeParam := c.Query("size")
		pageParam := c.Query("page")
		size, _ := strconv.ParseInt(sizeParam, 10, 0)
		page, _ := strconv.ParseInt(pageParam, 10, 0)
		if size <= 0 {
			size = 2
		}
		if page <= 0 {
			page = 1
		}
		pagination := NewPagination(page, size)
		pagination.Items = &users
		query := db.Model(&User{}).Order("id desc")
		pagination.Query = query

		urlQ := c.Request.URL.Query()
		urlQ.Del("page")
		urlQ.Del("size")

		err := pagination.Paginate().Find(pagination.Items).Error
		pagination.UrlQuery = urlQ.Encode()

		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = pagination
		c.JSON(http.StatusOK, resp)
	})

	authGroup.DELETE("/users", func(c *gin.Context) {
		resp := &APIResponse{}
		items := make([]int, 0)
		users := make([]User, 0)
		err := json.NewDecoder(c.Request.Body).Decode(&items)
		fmt.Println(items)

		if err != nil {
			fmt.Println(err.Error())
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer c.Request.Body.Close()

		err = db.Model(&User{}).Where("id in ?", items).Delete(&users).Error
		if err != nil {
			resp.ErrorCode = 2
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = users
		c.JSON(http.StatusOK, resp)
	})

	authGroup.DELETE("/users/:id", func(c *gin.Context) {
		resp := &APIResponse{}

		user := &User{}
		err := db.Model(user).Delete(user, c.Param("id")).Error
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = user
		c.JSON(http.StatusOK, resp)
	})

	authGroup.PUT("/users/:id", func(c *gin.Context) {
		resp := &APIResponse{}

		user := &User{}
		err := json.NewDecoder(c.Request.Body).Decode(user)
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer c.Request.Body.Close()
		err = db.Model(user).Where("id =?", c.Param("id")).Updates(user).Error
		if err != nil {
			resp.ErrorCode = 2
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		c.JSON(http.StatusOK, user)
	})

	authGroup.GET("/users/:id", func(c *gin.Context) {
		resp := &APIResponse{}
		user := &User{}
		err := db.Model(user).Where("id = ?", c.Param("id")).First(user).Error
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = user
		c.JSON(http.StatusOK, resp)
	})

	authGroup.POST("/dummies", func(c *gin.Context) {
		resp := &APIResponse{}

		dummy := &Dummy{}
		err := json.NewDecoder(c.Request.Body).Decode(dummy)
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer c.Request.Body.Close()
		err = db.Model(dummy).Create(dummy).Error
		if err != nil {
			resp.ErrorCode = 2
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		c.JSON(http.StatusOK, dummy)
	})

	authGroup.GET("/dummies", func(c *gin.Context) {

		resp := &APIResponse{}
		dummies := make([]Dummy, 0)

		sizeParam := c.Query("size")
		pageParam := c.Query("page")
		size, _ := strconv.ParseInt(sizeParam, 10, 0)
		page, _ := strconv.ParseInt(pageParam, 10, 0)
		if size <= 0 {
			size = 2
		}
		if page <= 0 {
			page = 1
		}
		pagination := NewPagination(page, size)
		pagination.Items = &dummies
		query := db.Model(&Dummy{}).Order("id desc")
		pagination.Query = query

		urlQ := c.Request.URL.Query()
		urlQ.Del("page")
		urlQ.Del("size")

		err := pagination.Paginate().Find(pagination.Items).Error
		pagination.UrlQuery = urlQ.Encode()

		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = pagination
		c.JSON(http.StatusOK, resp)
	})

	authGroup.DELETE("/dummies", func(c *gin.Context) {
		resp := &APIResponse{}
		items := make([]int, 0)
		dummies := make([]Dummy, 0)
		err := json.NewDecoder(c.Request.Body).Decode(&items)
		fmt.Println(items)

		if err != nil {
			fmt.Println(err.Error())
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer c.Request.Body.Close()

		err = db.Model(&Dummy{}).Where("id in ?", items).Delete(&dummies).Error
		if err != nil {
			resp.ErrorCode = 2
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = dummies
		c.JSON(http.StatusOK, resp)
	})

	authGroup.DELETE("/dummies/:id", func(c *gin.Context) {
		resp := &APIResponse{}

		dummy := &Dummy{}
		err := db.Model(dummy).Delete(dummy, c.Param("id")).Error
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = dummy
		c.JSON(http.StatusOK, resp)
	})

	authGroup.PUT("/dummies/:id", func(c *gin.Context) {
		resp := &APIResponse{}

		dummy := &Dummy{}
		err := json.NewDecoder(c.Request.Body).Decode(dummy)
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		defer c.Request.Body.Close()
		err = db.Model(dummy).Where("id =?", c.Param("id")).Updates(dummy).Error
		if err != nil {
			resp.ErrorCode = 2
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		c.JSON(http.StatusOK, dummy)
	})

	authGroup.GET("/dummies/:id", func(c *gin.Context) {
		resp := &APIResponse{}
		dummy := &Dummy{}
		err := db.Model(dummy).Where("id = ?", c.Param("id")).First(dummy).Error
		if err != nil {
			resp.ErrorCode = 1
			resp.ErrorMessage = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data = dummy
		c.JSON(http.StatusOK, resp)
	})

	// engine.POST("/auth", func(c *gin.Context) {
	// 	// persiter jwt for the user, also check with jwt.id(unique and exsited)
	// 	up := &LoginPayload{}

	// 	resp := &APIResponse{}

	// 	err := json.NewDecoder(c.Request.Body).Decode(up)
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	defer c.Request.Body.Close()

	// 	user := &User{}
	// 	err = db.Model(user).Where("name = ? and intro = ?", up.Username, up.Password).First(&user).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 2
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}

	// 	claims := auth.NewJWTClaims(fmt.Sprintf("%d", user.ID), xid.New().String())
	// 	token, err := util.GenerateJWT(JWTKey, jwt.SigningMethodHS256, claims)
	// 	if err != nil {
	// 		resp.ErrorCode = 3
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = token

	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.POST("/users", func(c *gin.Context) {
	// 	resp := &APIResponse{}

	// 	user := &User{}
	// 	err := json.NewDecoder(c.Request.Body).Decode(user)
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	defer c.Request.Body.Close()
	// 	err = db.Model(user).Create(user).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 2
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, user)
	// })

	// engine.GET("/users", func(c *gin.Context) {

	// 	resp := &APIResponse{}
	// 	users := make([]User, 0)

	// 	sizeParam := c.Query("size")
	// 	pageParam := c.Query("page")
	// 	size, _ := strconv.ParseInt(sizeParam, 10, 0)
	// 	page, _ := strconv.ParseInt(pageParam, 10, 0)
	// 	if size <= 0 {
	// 		size = 2
	// 	}
	// 	if page <= 0 {
	// 		page = 1
	// 	}
	// 	pagination := NewPagination(page, size)
	// 	pagination.Items = &users
	// 	query := db.Model(&User{}).Order("id desc")
	// 	pagination.Query = query

	// 	urlQ := c.Request.URL.Query()
	// 	urlQ.Del("page")
	// 	urlQ.Del("size")

	// 	err := pagination.Paginate().Find(pagination.Items).Error
	// 	pagination.UrlQuery = urlQ.Encode()

	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = pagination
	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.DELETE("/users", func(c *gin.Context) {
	// 	resp := &APIResponse{}
	// 	items := make([]int, 0)
	// 	users := make([]User, 0)
	// 	err := json.NewDecoder(c.Request.Body).Decode(&items)
	// 	fmt.Println(items)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	defer c.Request.Body.Close()

	// 	err = db.Model(&User{}).Where("id in ?", items).Delete(&users).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 2
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = users
	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.DELETE("/users/:id", func(c *gin.Context) {
	// 	resp := &APIResponse{}

	// 	user := &User{}
	// 	err := db.Model(user).Delete(user, c.Param("id")).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = user
	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.PUT("/users/:id", func(c *gin.Context) {
	// 	resp := &APIResponse{}

	// 	user := &User{}
	// 	err := json.NewDecoder(c.Request.Body).Decode(user)
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	defer c.Request.Body.Close()
	// 	err = db.Model(user).Where("id =?", c.Param("id")).Updates(user).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 2
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, user)
	// })

	// engine.GET("/users/:id", func(c *gin.Context) {
	// 	resp := &APIResponse{}
	// 	user := &User{}
	// 	err := db.Model(user).Where("id = ?", c.Param("id")).First(user).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = user
	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.POST("/dummies", func(c *gin.Context) {
	// 	resp := &APIResponse{}

	// 	dummy := &Dummy{}
	// 	err := json.NewDecoder(c.Request.Body).Decode(dummy)
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	defer c.Request.Body.Close()
	// 	err = db.Model(dummy).Create(dummy).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 2
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, dummy)
	// })

	// engine.GET("/dummies", func(c *gin.Context) {

	// 	resp := &APIResponse{}
	// 	dummies := make([]Dummy, 0)

	// 	sizeParam := c.Query("size")
	// 	pageParam := c.Query("page")
	// 	size, _ := strconv.ParseInt(sizeParam, 10, 0)
	// 	page, _ := strconv.ParseInt(pageParam, 10, 0)
	// 	if size <= 0 {
	// 		size = 2
	// 	}
	// 	if page <= 0 {
	// 		page = 1
	// 	}
	// 	pagination := NewPagination(page, size)
	// 	pagination.Items = &dummies
	// 	query := db.Model(&Dummy{}).Order("id desc")
	// 	pagination.Query = query

	// 	urlQ := c.Request.URL.Query()
	// 	urlQ.Del("page")
	// 	urlQ.Del("size")

	// 	err := pagination.Paginate().Find(pagination.Items).Error
	// 	pagination.UrlQuery = urlQ.Encode()

	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = pagination
	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.DELETE("/dummies", func(c *gin.Context) {
	// 	resp := &APIResponse{}
	// 	items := make([]int, 0)
	// 	dummies := make([]Dummy, 0)
	// 	err := json.NewDecoder(c.Request.Body).Decode(&items)
	// 	fmt.Println(items)

	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	defer c.Request.Body.Close()

	// 	err = db.Model(&Dummy{}).Where("id in ?", items).Delete(&dummies).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 2
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = dummies
	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.DELETE("/dummies/:id", func(c *gin.Context) {
	// 	resp := &APIResponse{}

	// 	dummy := &Dummy{}
	// 	err := db.Model(dummy).Delete(dummy, c.Param("id")).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = dummy
	// 	c.JSON(http.StatusOK, resp)
	// })

	// engine.PUT("/dummies/:id", func(c *gin.Context) {
	// 	resp := &APIResponse{}

	// 	dummy := &Dummy{}
	// 	err := json.NewDecoder(c.Request.Body).Decode(dummy)
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	defer c.Request.Body.Close()
	// 	err = db.Model(dummy).Where("id =?", c.Param("id")).Updates(dummy).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 2
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, dummy)
	// })

	// engine.GET("/dummies/:id", func(c *gin.Context) {
	// 	resp := &APIResponse{}
	// 	dummy := &Dummy{}
	// 	err := db.Model(dummy).Where("id = ?", c.Param("id")).First(dummy).Error
	// 	if err != nil {
	// 		resp.ErrorCode = 1
	// 		resp.ErrorMessage = err.Error()
	// 		c.JSON(http.StatusOK, resp)
	// 		return
	// 	}
	// 	resp.Data = dummy
	// 	c.JSON(http.StatusOK, resp)
	// })

	log.Fatalln(engine.Run(":8980"))
}
