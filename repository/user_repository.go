package repository

import (
	"context"

	"easy-dictionary-server/domain"

	"go.uber.org/zap"
)

const (
	ErrUserNotFound = encodingError("User is not found")
)

type userRepository struct {
	// database   mongo.Database
	// collection string
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	return nil
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	//TODO: mock for now
	users := []domain.User{
		{ID: 1, Email: "example1@gmail.com", ProviderId: "Google", UID: "example1@gmail.com"},
		{ID: 2, Email: "example2@gmail.com", ProviderId: "Google", UID: "example2@gmail.com"},
	}

	return users, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	zap.S().Debugf("GetByEmail %s", email)
	//TODO: will be fetch from database
	users, error := ur.Fetch(c)
	zap.S().Debugf("found users %d", len(users))
	if error != nil {
		zap.S().Debugf("fetch error %s", error.Error())
	} else {
		var foundUser *domain.User
		for i := range users {
			if users[i].Email == email {
				foundUser = &users[i]
				break
			}

		}
		if foundUser != nil {
			return foundUser, nil
		}
	}
	return nil, ErrUserNotFound
}

func (ur *userRepository) GetByID(c context.Context, id int) (*domain.User, error) {
	zap.S().Debugf("GetByID %d", id)
	users, error := ur.Fetch(c)
	//TODO: will be fetch from database
	if error != nil {
		zap.S().Debugf("fetch error %s", error.Error())
	} else {
		var foundUser *domain.User
		for i := range users {
			if users[i].ID == id {
				foundUser = &users[i]
				break
			}

		}
		if foundUser != nil {
			return foundUser, nil
		}
	}
	return nil, ErrUserNotFound
}
