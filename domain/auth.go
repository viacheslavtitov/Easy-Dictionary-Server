package domain

import (
	"context"
)

type AuthRequest struct {
	Email      string
	ProviderId string
	UID        string
}

type AuthResponse struct {
	AccessToken string
}

type AuthUseCase interface {
	GetUserByEmail(context context.Context, email string) (*User, error)
	CreateAccessToken(user *User, secret string) (accessToken string, err error)
}
