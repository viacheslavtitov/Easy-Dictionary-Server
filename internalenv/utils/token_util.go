package utils

import (
	domainUser "easy-dictionary-server/domain/user"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *domainUser.User, appName string, secret string, duration time.Duration) (string, error) {
	claims := &jwt.RegisteredClaims{
		Issuer:    appName,
		Subject:   fmt.Sprintf("%d", user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
