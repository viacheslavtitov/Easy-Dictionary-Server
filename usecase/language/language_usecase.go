package usecase

import (
	"context"

	domainLanguage "easy-dictionary-server/domain/language"
	commonUseCase "easy-dictionary-server/usecase"
)

type languageUsecase struct {
	languageRepository domainLanguage.LanguageRepository
	contextTimeout     int
}

func NewLanguageUsecase(languageRepository domainLanguage.LanguageRepository, timeout int) domainLanguage.LanguageUseCase {
	return &languageUsecase{
		languageRepository: languageRepository,
		contextTimeout:     timeout,
	}
}

func (lu *languageUsecase) GetAllForUser(c context.Context, userId int) (*[]domainLanguage.Language, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.languageRepository.GetAllForUser(ctx, userId)
}

func (lu *languageUsecase) Create(c context.Context, userId int, name string, code string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.languageRepository.Create(ctx, userId, domainLanguage.Language{
		Name: name,
		Code: code})
}

func (lu *languageUsecase) Update(c context.Context, userId int, id int, name string, code string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.languageRepository.Update(ctx, userId, domainLanguage.Language{
		ID:   id,
		Name: name,
		Code: code})
}

func (lu *languageUsecase) DeleteById(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.languageRepository.DeleteById(ctx, id)
}
