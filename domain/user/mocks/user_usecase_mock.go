package mocks

import (
	"context"
	userDomain "easy-dictionary-server/domain/user"

	middleware "easy-dictionary-server/api/middleware"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func GetMockUser(id int, email string) *userDomain.User {
	return &userDomain.User{
		ID:        id,
		FirstName: "Jane",
		LastName:  "Doe",
		Role:      middleware.Client.VALUE,
		Providers: &[]userDomain.UserProviders{
			{
				ProviderName:   "email",
				HashedPassword: "hashed_password",
				Email:          email,
			},
		},
	}
}

func (m *MockUserUseCase) RegisterUser(context context.Context, firstName string, lastName string, role string,
	email string, provider string, password string, providerToken string) (*userDomain.User, error) {
	user := userDomain.User{
		ID:        1,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		Providers: &[]userDomain.UserProviders{
			{
				Email:          email,
				ProviderName:   provider,
				HashedPassword: password,
				ProviderToken:  providerToken,
			},
		},
	}
	return &user, nil
}

func (m *MockUserUseCase) GetByID(context context.Context, id int) (*userDomain.User, error) {
	return GetMockUser(id, ""), nil
}

func (m *MockUserUseCase) GetByEmail(context context.Context, email string) (*userDomain.User, error) {
	return GetMockUser(1, email), nil
}

func (m *MockUserUseCase) GetAllUsers(context context.Context) ([]*userDomain.User, error) {
	users := []*userDomain.User{
		GetMockUser(1, "example1@email.com"),
		GetMockUser(2, "example2@email.com"),
	}
	return users, nil
}

func (m *MockUserUseCase) UpdateUser(context context.Context, user *userDomain.User) (*userDomain.User, error) {
	return user, nil
}

func (m *MockUserUseCase) DeleteUser(context context.Context, id int) error {
	return nil
}
