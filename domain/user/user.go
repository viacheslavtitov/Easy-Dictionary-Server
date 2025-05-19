package domain

import (
	"context"
	"time"
)

type User struct {
	ID         int              `json:"id"`
	FirstName  string           `json:"first_name"`
	SecondName string           `json:"second_name"`
	CreatedAt  time.Time        `json:"created_at"`
	Providers  *[]UserProviders `json:"providers"`
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
	SecondName    string `json:"second_name" binding:"required"`
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
	RegisterUser(context context.Context, firstName string, secondName string,
		email string, provider string, password string, providerToken string) (*User, error)
	UpdateUser(context context.Context, id int, firstName string, secondName string) (*User, error)
	DeleteUser(context context.Context, id int) error
}

type UserRepository interface {
	Create(context context.Context, user *User) (*User, error)
	GetAllUsers(context context.Context) ([]*User, error)
	GetByEmail(context context.Context, email string) (*User, error)
	GetByID(context context.Context, id int) (*User, error)
	UpdateUser(context context.Context, user *User) (*User, error)
	DeleteUser(context context.Context, id int) error
}
