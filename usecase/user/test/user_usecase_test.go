package usecase_test

import (
	userDomain "easy-dictionary-server/domain/user"
	userDomainMock "easy-dictionary-server/domain/user/mocks"
	testutils "easy-dictionary-server/internalenv/testutils"
	userUseCase "easy-dictionary-server/usecase/user"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser_UnitTest(t *testing.T) {
	mockURepository := new(userDomainMock.MockUserRepository)
	expectedUser := userDomainMock.GetMockUser(0, "example1@email.com")
	mockURepository.On("Create", expectedUser).Return(expectedUser, nil)
	mockURepository.On("Create", mock.MatchedBy(func(u *userDomain.User) bool {
		return u.FirstName == "Jane" && u.Role == "client"
	})).Return(expectedUser, nil)
	userUseCase := userUseCase.NewUserUsecase(mockURepository, 10)

	registeredUser, err := userUseCase.RegisterUser(testutils.GetTestGinContext(), expectedUser.FirstName, expectedUser.LastName,
		expectedUser.Role, expectedUser.FindEmailProvider().Email, expectedUser.FindEmailProvider().ProviderName, expectedUser.FindEmailProvider().HashedPassword, "token")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.FirstName, registeredUser.FirstName)
}

func TestUpdateUser_UnitTest(t *testing.T) {
	mockURepository := new(userDomainMock.MockUserRepository)
	expectedUser := userDomainMock.GetMockUser(0, "example1@email.com")
	mockURepository.On("UpdateUser", expectedUser).Return(expectedUser, nil)
	userUseCase := userUseCase.NewUserUsecase(mockURepository, 10)
	updatedUser, err := userUseCase.UpdateUser(testutils.GetTestGinContext(), expectedUser.ID, expectedUser.FirstName, expectedUser.LastName)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, updatedUser.ID)
}

func TestGetByID_UnitTest(t *testing.T) {
	mockURepository := new(userDomainMock.MockUserRepository)
	expectedUser := userDomainMock.GetMockUser(0, "example1@email.com")
	mockURepository.On("GetByID", expectedUser).Return(expectedUser, nil)
	userUseCase := userUseCase.NewUserUsecase(mockURepository, 10)
	user, err := userUseCase.GetByID(testutils.GetTestGinContext(), expectedUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
}

func TestGetAllUsers_UnitTest(t *testing.T) {
	mockURepository := new(userDomainMock.MockUserRepository)
	expectedUsers := []*userDomain.User{
		userDomainMock.GetMockUser(1, "example1@email.com"),
		userDomainMock.GetMockUser(2, "example2@email.com"),
	}
	mockURepository.On("GetAllUsers", mock.Anything).Return(expectedUsers, nil)
	userUseCase := userUseCase.NewUserUsecase(mockURepository, 10)
	users, err := userUseCase.GetAllUsers(testutils.GetTestGinContext())
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, expectedUsers[0].Role, users[0].Role)
}

func TestDeleteUser_UnitTest(t *testing.T) {
	mockURepository := new(userDomainMock.MockUserRepository)
	userUseCase := userUseCase.NewUserUsecase(mockURepository, 10)
	err := userUseCase.DeleteUser(testutils.GetTestGinContext(), 1)
	assert.NoError(t, err)
}
