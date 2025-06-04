package usecase

import (
	"context"
	"time"

	domain "easy-dictionary-server/domain"
	domainUser "easy-dictionary-server/domain/user"
	"easy-dictionary-server/internalenv/utils"
	commonUseCase "easy-dictionary-server/usecase"
)

type authUsecase struct {
	userRepository domainUser.UserRepository
	contextTimeout int
}

func NewAuthUsecase(userRepository domainUser.UserRepository, timeout int) domain.AuthUseCase {
	return &authUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *authUsecase) GetUserByEmail(c context.Context, email string) (*domainUser.User, *int, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *authUsecase) CreateAccessToken(user *domainUser.User, appName string, secret string, role string, duration time.Duration, userId int) (accessToken string, err error) {
	return utils.CreateAccessToken(user, appName, secret, role, duration, userId)
}
