package domain

import (
	"context"
)

type User struct {
	ID         int
	Email      string
	ProviderId string
	UID        string
}

type UserRepository interface {
	Create(context context.Context, user *User) error
	Fetch(context context.Context) ([]User, error)
	GetByEmail(context context.Context, email string) (*User, error)
	GetByID(context context.Context, id int) (*User, error)
}
