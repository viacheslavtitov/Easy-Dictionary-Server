package domain

import (
	"context"
	domainUser "easy-dictionary-server/domain/user"
	"time"
)

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Provider string `json:"provider" binding:"required"`
	UID      string `json:"uid" binding:"required"`
}

type AuthResponse struct {
	AccessToken string
}

type AuthUseCase interface {
	GetUserByEmail(context context.Context, email string) (*domainUser.User, error)
	CreateAccessToken(user *domainUser.User, appName string, secret string, duration time.Duration) (accessToken string, err error)
}
