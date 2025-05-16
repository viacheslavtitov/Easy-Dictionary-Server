package domain

import (
	"context"
	domainUser "easy-dictionary-server/domain/user"
)

type AuthRequest struct {
	Email      string `json:"email" binding:"required,email"`
	ProviderID string `json:"provider_id" binding:"required"`
	UID        string `json:"uid" binding:"required"`
}

type AuthResponse struct {
	AccessToken string
}

type AuthUseCase interface {
	GetUserByEmail(context context.Context, email string) (*domainUser.User, error)
	CreateAccessToken(user *domainUser.User, secret string) (accessToken string, err error)
}
