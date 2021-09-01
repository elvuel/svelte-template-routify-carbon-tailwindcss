package util

import (
	"net/http"
)

const BearerPrefix = "Bearer "

// Get bearer token from http request header
func GetBearerToken(r *http.Request) string {
	auth := r.Header.Get("Authorization")
	n := len(BearerPrefix)
	if len(auth) > n {
		return auth[n:]
	}

	return ""
}
