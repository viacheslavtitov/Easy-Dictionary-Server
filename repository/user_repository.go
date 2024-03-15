package repository

import (
	"context"
	"sort"

	"github.com/viacheslavtitov/easy-dictionary-server/domain"
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
	users := []domain.User{
		{ID: 1, Email: "example1@gmail.com", ProviderId: "Google", UID: "example1@gmail.com"},
		{ID: 2, Email: "example2@gmail.com", ProviderId: "Google", UID: "example2@gmail.com"},
	}

	return users, nil
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	zap.S().Debugf("GetByEmail %s", email)
	users, error := ur.Fetch(c)
	zap.S().Debugf("found users %d", len(users))
	if error != nil {
		zap.S().Debugf("fetch error %s", error.Error())
	} else {
		idx := sort.Search(len(users), func(i int) bool {
			return users[i].Email == email
		})
		zap.S().Debugf("found index %d", idx)
		if idx <= len(users) && idx >= 0 {
			return &users[idx-1], nil
		}
	}
	return nil, ErrUserNotFound
}

func (ur *userRepository) GetByID(c context.Context, id int) (*domain.User, error) {
	users, error := ur.Fetch(c)
	if error != nil {
		idx := sort.Search(len(users), func(i int) bool {
			return users[i].ID == id
		})
		if idx < len(users) && idx >= 0 {
			return &users[idx], nil
		} else {
			return nil, ErrUserNotFound
		}
	} else {
		return nil, error
	}
}
