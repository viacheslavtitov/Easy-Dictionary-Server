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
	GetAllUsers(context context.Context) ([]*User, error)
	GetByEmail(context context.Context, email string) (*User, error)
	GetByID(context context.Context, id int) (*User, error)
	UpdateUser(context context.Context, user *User) (*User, error)
	DeleteUser(context context.Context, id int) error
}
