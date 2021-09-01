package middleware

import (
	"mmm/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type AuthFail func(c *gin.Context, err error) interface{}
type AuthSuccess func(c *gin.Context, token *jwt.Token)

var (
	ErrJWTNotFound = errors.New("JWT not found")
	ErrJWTInvalid  = errors.New("JWT invalid")
)

// gin framework JWT middleware
func JWT(key interface{}, claims jwt.Claims, fail AuthFail, success AuthSuccess) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := util.GetBearerToken(c.Request)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fail(c, ErrJWTNotFound))
			// c.AbortWithStatusJSON(http.StatusOK, fail(c, ErrJWTNotFound))
			return
		}

		token, err := util.ParseJWT(tokenString, key, claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fail(c, err))
			// c.AbortWithStatusJSON(http.StatusOK, fail(c, err))
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fail(c, ErrJWTInvalid))
			// c.AbortWithStatusJSON(http.StatusOK, fail(c, ErrJWTInvalid))
			return
		}

		success(c, token)

		c.Next()
	}
}
