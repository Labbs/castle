package domain

import (
	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
