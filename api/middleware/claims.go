package middleware

import (
	"github.com/golang-jwt/jwt/v5"
)

type Role struct {
	VALUE string
}

var (
	Admin  = Role{VALUE: `admin`}
	Client = Role{VALUE: `client`}
)

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}
