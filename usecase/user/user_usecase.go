package usecase

import (
	"context"

	domainUser "easy-dictionary-server/domain/user"
	commonUseCase "easy-dictionary-server/usecase"
)

type userUsecase struct {
	userRepository domainUser.UserRepository
	contextTimeout int
}

func NewUserUsecase(userRepository domainUser.UserRepository, timeout int) domainUser.UserUseCase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (usecase *userUsecase) RegisterUser(c context.Context, firstName string, lastName string, role string, email string,
	provider string, password string, providerToken string) (*domainUser.User, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
	defer cancel()
	user := &domainUser.User{
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		Providers: &[]domainUser.UserProviders{
			{
				Email:          email,
				ProviderName:   provider,
				HashedPassword: password,
				ProviderToken:  providerToken,
			},
		},
	}
	return usecase.userRepository.Create(ctx, user)
}

func (lu *userUsecase) UpdateUser(c context.Context, id int, firstName string, LastName string) (*domainUser.User, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	user := &domainUser.User{
		ID:        id,
		FirstName: firstName,
		LastName:  LastName,
	}
	return lu.userRepository.UpdateUser(ctx, user)
}

func (lu *userUsecase) DeleteUser(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.DeleteUser(ctx, id)
}

func (lu *userUsecase) GetByID(c context.Context, id int) (*domainUser.User, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetByID(ctx, id)
}

func (lu *userUsecase) GetAllUsers(c context.Context) ([]*domainUser.User, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetAllUsers(ctx)
}
