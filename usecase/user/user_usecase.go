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
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(usecase.contextTimeout))
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
	return usecase.userRepository.Create(user)
}

func (lu *userUsecase) UpdateUser(c context.Context, id int, uuid string, firstName string, LastName string) (*domainUser.User, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	user := &domainUser.User{
		UUID:      uuid,
		FirstName: firstName,
		LastName:  LastName,
	}
	return lu.userRepository.UpdateUser(user, id)
}

func (lu *userUsecase) DeleteUser(c context.Context, id int) (int64, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.DeleteUser(id)
}

func (lu *userUsecase) GetByID(c context.Context, id int) (*domainUser.User, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetByID(id)
}

func (lu *userUsecase) GetByUUID(c context.Context, uuid string) (*domainUser.User, *int, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetByUUID(uuid)
}

func (lu *userUsecase) GetAllUsers(c context.Context) ([]*domainUser.User, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetAllUsers()
}

func (lu *userUsecase) GetRefreshToken(c context.Context, refreshToken string) (*domainUser.RefreshToken, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetRefreshToken(refreshToken)
}

func (lu *userUsecase) GetRefreshTokenByUserUUID(c context.Context, userUUID string) (*domainUser.RefreshToken, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetRefreshTokenByUserUUID(userUUID)
}

func (lu *userUsecase) DeleteRefreshToken(c context.Context, tokenUUID string) (int64, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.DeleteRefreshToken(tokenUUID)
}
