package usecase

import (
	"context"

	domainTranslationCategory "easy-dictionary-server/domain/translation/category"
	commonUseCase "easy-dictionary-server/usecase"
)

type translationCategoryUsecase struct {
	translationCategoryRepository domainTranslationCategory.TranslationCategoryRepository
	contextTimeout                int
}

func NewTranslationCategoryUsecase(translationCategoryRepository domainTranslationCategory.TranslationCategoryRepository, timeout int) domainTranslationCategory.TranslationCategoryUseCase {
	return &translationCategoryUsecase{
		translationCategoryRepository: translationCategoryRepository,
		contextTimeout:                timeout,
	}
}

func (lu *translationCategoryUsecase) GetAllForUser(c context.Context, userId int) (*[]domainTranslationCategory.TranslationCategory, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationCategoryRepository.GetAllForUser(ctx, userId)
}

func (lu *translationCategoryUsecase) Create(c context.Context, userId int, dictionaryId int, name string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationCategoryRepository.Create(ctx, userId, &domainTranslationCategory.TranslationCategory{
		UserId:       userId,
		Name:         name,
		DictionaryId: dictionaryId})
}

func (lu *translationCategoryUsecase) Update(c context.Context, userId int, id int, dictionaryId int, name string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationCategoryRepository.Update(ctx, userId, &domainTranslationCategory.TranslationCategory{
		ID:           id,
		Name:         name,
		DictionaryId: dictionaryId,
		UserId:       userId})
}

func (lu *translationCategoryUsecase) DeleteById(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(lu.contextTimeout))
	defer cancel()
	return lu.translationCategoryRepository.DeleteById(ctx, id)
}
