package domain

import (
	"context"
	domainUser "easy-dictionary-server/domain/user"
	"time"
)

type AuthRequest struct {
	Email         string `json:"email" binding:"email"`
	Provider      string `json:"provider" binding:"required"`
	ProviderToken string `json:"provider_token"`
	Password      string `json:"password"`
}

type AuthResponse struct {
	AccessToken string
}

type AuthUseCase interface {
	GetUserByEmail(context context.Context, email string) (*domainUser.User, error)
	CreateAccessToken(user *domainUser.User, appName string, secret string, role string, duration time.Duration) (accessToken string, err error)
}
