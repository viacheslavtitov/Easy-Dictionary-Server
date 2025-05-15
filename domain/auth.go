package domain

import (
	"context"
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
	GetUserByEmail(context context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string) (accessToken string, err error)
}
