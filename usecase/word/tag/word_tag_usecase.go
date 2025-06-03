package usecase

import (
	"context"

	domainWordTag "easy-dictionary-server/domain/word/tag"
	commonUseCase "easy-dictionary-server/usecase"
)

type wordTagUsecase struct {
	wordTagRepository domainWordTag.WordTagRepository
	contextTimeout    int
}

func NewWordTagUsecase(wordTagRepository domainWordTag.WordTagRepository, timeout int) domainWordTag.WordTagUseCase {
	return &wordTagUsecase{
		wordTagRepository: wordTagRepository,
		contextTimeout:    timeout,
	}
}

func (wu *wordTagUsecase) GetAllForDictionary(c context.Context, dictionaryId int) (*[]domainWordTag.WordTag, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTagRepository.GetAllForDictionary(ctx, dictionaryId)
}

func (wu *wordTagUsecase) Create(c context.Context, dictionaryId int, name string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTagRepository.Create(ctx, &domainWordTag.WordTag{
		DictionaryId: dictionaryId,
		Name:         name})
}

func (wu *wordTagUsecase) Update(c context.Context, id int, dictionaryId int, name string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTagRepository.Update(ctx, &domainWordTag.WordTag{
		ID:           id,
		DictionaryId: dictionaryId,
		Name:         name})
}

func (wu *wordTagUsecase) DeleteById(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTagRepository.DeleteById(ctx, id)
}
