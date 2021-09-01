package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	jwt.StandardClaims
	// Audience  string `json:"aud,omitempty"` // 受众
	// ExpiresAt int64  `json:"exp,omitempty"`  // 过期时间
	// Id        string `json:"jti,omitempty"` // 唯一标识
	// IssuedAt  int64  `json:"iat,omitempty"` // 签发时间
	// Issuer    string `json:"iss,omitempty"` // 签发人
	// NotBefore int64  `json:"nbf,omitempty"` // 生效时间
	// Subject   string `json:"sub,omitempty"` // 主题

	// iss (issuer): Issuer of the JWT
	// sub (subject): Subject of the JWT (the user)
	// aud (audience): Recipient for which the JWT is intended
	// exp (expiration time): Time after which the JWT expires
	// nbf (not before time): Time before which the JWT must not be accepted for processing
	// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
	// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)
}

func NewJWTClaims(aud, id string) *JWTClaims {
	return &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			Audience:  aud,
			Id:        id,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "platform.abntee",
			NotBefore: time.Now().Unix(),
			Subject:   "auth.platform.abntee",
		},
	}
}
