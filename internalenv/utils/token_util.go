package utils

import (
	middleware "easy-dictionary-server/api/middleware"
	domainUser "easy-dictionary-server/domain/user"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func CreateAccessToken(user *domainUser.User, appName string, secret string, role string, duration time.Duration, userId int) (string, error) {
	claims := middleware.Claims{
		Role:   role,
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    appName,
			Subject:   fmt.Sprintf("%s", user.UUID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	zap.S().Debugf("Create access token with user role %s, user id %d, user uuid %s", claims.Role, claims.UserID, claims.Subject)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
