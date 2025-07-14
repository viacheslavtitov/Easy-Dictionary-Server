package domain

import (
	"context"
	"time"
)

type User struct {
	UUID      string           `json:"uuid"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	CreatedAt time.Time        `json:"created_at"`
	Providers *[]UserProviders `json:"providers"`
	Role      string           `json:"-"`
}

type UserProviders struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	ProviderName   string    `json:"name"`
	ProviderToken  string    `json:"provider_token"`
	CreatedAt      time.Time `json:"created_at"`
}

type RegisterUserRequest struct {
	Email         string `json:"email" binding:"email"`
	Provider      string `json:"provider" binding:"required"`
	Password      string `json:"password"`
	ProviderToken string `json:"provider_token"`
	FirstName     string `json:"first_name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
}

type EditUserRequest struct {
	UUID      string `json:"uuid" binding:"required"`
	Email     string `json:"email" binding:"email,required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type RefreshToken struct {
	ID        string    `json:"-"`
	UserUUID  string    `json:"-"`
	Token     string    `json:"-"`
	ExpiresAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`
}

func (user *User) FindEmailProvider() (provider *UserProviders) {
	if user.Providers == nil {
		return nil
	}
	for _, p := range *user.Providers {
		if p.ProviderName == "email" {
			return &p
		}
	}
	return nil
}

type UserUseCase interface {
	RegisterUser(context context.Context, firstName string, lastName string, role string,
		email string, provider string, password string, providerToken string) (*User, error)
	UpdateUser(context context.Context, id int, uuid string, firstName string, lastName string) (*User, error)
	DeleteUser(context context.Context, id int) (int64, error)
	GetByID(context context.Context, id int) (*User, error)
	GetByUUID(context context.Context, uuid string) (*User, *int, error)
	GetAllUsers(context context.Context) ([]*User, error)
	GetRefreshToken(context context.Context, refreshToken string) (*RefreshToken, error)
	GetRefreshTokenByUserUUID(context context.Context, refreshToken string) (*RefreshToken, error)
	DeleteRefreshToken(context context.Context, tokenUUID string) (int64, error)
}

type UserRepository interface {
	Create(user *User) (*User, error)
	GetAllUsers() ([]*User, error)
	GetByEmail(email string) (*User, *int, error)
	GetByID(id int) (*User, error)
	GetByUUID(uuid string) (*User, *int, error)
	UpdateUser(user *User, userId int) (*User, error)
	DeleteUser(id int) (int64, error)
	AddRefreshToken(userUUID string, refreshToken string, expiresAt time.Time) (*time.Time, error)
	GetRefreshToken(refreshToken string) (*RefreshToken, error)
	GetRefreshTokenByUserUUID(refreshToken string) (*RefreshToken, error)
	DeleteRefreshToken(tokenUUID string) (int64, error)
}
