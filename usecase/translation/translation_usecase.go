package usecase

import (
	"context"

	domainTranslation "easy-dictionary-server/domain/translation"
	commonUseCase "easy-dictionary-server/usecase"
)

type translationUsecase struct {
	translationRepository domainTranslation.TranslationRepository
	contextTimeout        int
}

func NewTranslationUsecase(translationRepository domainTranslation.TranslationRepository, timeout int) domainTranslation.TranslationUseCase {
	return &translationUsecase{
		translationRepository: translationRepository,
		contextTimeout:        timeout,
	}
}

func (lu *translationUsecase) GetAllForWord(c context.Context, wordId int) (*[]domainTranslation.Translation, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationRepository.GetAllForWord(ctx, wordId)
}

func (lu *translationUsecase) Create(c context.Context, wordId int, categoryId *int, translate string, description *string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationRepository.Create(ctx, wordId, &domainTranslation.Translation{
		WordId:      wordId,
		CategoryId:  categoryId,
		Translate:   translate,
		Description: description})
}

func (lu *translationUsecase) Update(c context.Context, id int, categoryId *int, translate string, description *string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationRepository.Update(ctx, &domainTranslation.Translation{
		ID:          id,
		CategoryId:  categoryId,
		Translate:   translate,
		Description: description})
}

func (lu *translationUsecase) DeleteById(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationRepository.DeleteById(ctx, id)
}
