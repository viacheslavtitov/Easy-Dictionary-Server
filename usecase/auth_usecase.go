package usecase

import (
	"context"

	"github.com/viacheslavtitov/easy-dictionary-server/domain"
	"github.com/viacheslavtitov/easy-dictionary-server/internal/utils"
)

type authUsecase struct {
	userRepository domain.UserRepository
	contextTimeout int
}

func NewAuthUsecase(userRepository domain.UserRepository, timeout int) domain.AuthUseCase {
	return &authUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *authUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *authUsecase) CreateAccessToken(user *domain.User, secret string) (accessToken string, err error) {
	return utils.CreateAccessToken(user, secret)
}
