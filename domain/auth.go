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
	AccessToken     string
	RefreshToken    string
	RefreshTokenExp time.Time
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"RefreshToken" binding:"required"`
}

type AuthUseCase interface {
	GetUserByEmail(context context.Context, email string) (*domainUser.User, *int, error)
	CreateAccessToken(user *domainUser.User, appName string, secret string, role string, duration time.Duration, userId int) (accessToken string, err error)
	CreateRefreshToken(context context.Context, userUUID string, duration time.Duration) (refreshToken string, expiresAt time.Time, createdAt *time.Time, err error)
}
