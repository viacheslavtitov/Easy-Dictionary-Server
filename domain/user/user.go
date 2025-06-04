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
	Role      string
}

type UserProviders struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	HashedPassword string
	ProviderName   string `json:"name"`
	ProviderToken  string
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
	GetByUUID(context context.Context, uuid string) (*User, error)
	GetAllUsers(context context.Context) ([]*User, error)
}

type UserRepository interface {
	Create(context context.Context, user *User) (*User, error)
	GetAllUsers(context context.Context) ([]*User, error)
	GetByEmail(context context.Context, email string) (*User, *int, error)
	GetByID(context context.Context, id int) (*User, error)
	GetByUUID(context context.Context, uuid string) (*User, error)
	UpdateUser(context context.Context, user *User, userId int) (*User, error)
	DeleteUser(context context.Context, id int) (int64, error)
}
