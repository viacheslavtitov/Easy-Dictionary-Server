package mocks

import (
	"context"
	userDomain "easy-dictionary-server/domain/user"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(context context.Context, user *userDomain.User) (*userDomain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*userDomain.User), nil
}

func (m *MockUserRepository) GetByID(context context.Context, id int) (*userDomain.User, error) {
	return GetMockUser(id, ""), nil
}

func (m *MockUserRepository) GetByEmail(context context.Context, email string) (*userDomain.User, error) {
	return GetMockUser(1, email), nil
}

func (m *MockUserRepository) GetAllUsers(context context.Context) ([]*userDomain.User, error) {
	users := []*userDomain.User{
		GetMockUser(1, "example1@email.com"),
		GetMockUser(2, "example2@email.com"),
	}
	return users, nil
}

func (m *MockUserRepository) UpdateUser(context context.Context, user *userDomain.User) (*userDomain.User, error) {
	return user, nil
}

func (m *MockUserRepository) DeleteUser(context context.Context, id int) error {
	return nil
}
