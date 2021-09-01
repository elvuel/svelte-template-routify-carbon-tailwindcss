package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// Generate jwt
func GenerateJWT(key interface{}, method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	return jwt.NewWithClaims(method, claims).SignedString(key)
}

// Parse jwt
func ParseJWT(token string, key interface{}, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	// return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return key, nil
	// })
}
