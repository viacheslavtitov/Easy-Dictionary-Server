package utils

import (
	domainUser "easy-dictionary-server/domain/user"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *domainUser.User, secret string) (accessToken string, err error) {
	claims := &jwt.RegisteredClaims{
		Issuer: "user",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}
