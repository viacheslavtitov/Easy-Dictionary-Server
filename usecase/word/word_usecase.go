package usecase

import (
	"context"

	domainWord "easy-dictionary-server/domain/word"
	commonUseCase "easy-dictionary-server/usecase"
)

type wordUsecase struct {
	wordRepository domainWord.WordRepository
	contextTimeout int
}

func NewWordUsecase(wordRepository domainWord.WordRepository, timeout int) domainWord.WordUseCase {
	return &wordUsecase{
		wordRepository: wordRepository,
		contextTimeout: timeout,
	}
}

func (wu *wordUsecase) GetAllForDictionary(c context.Context, dictionaryId int) (*[]domainWord.Word, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.GetAllForDictionary(ctx, dictionaryId)
}

func (wu *wordUsecase) Create(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string, categoryId *int) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.Create(ctx, dictionaryId, &domainWord.Word{
		DictionaryId: dictionaryId,
		Original:     original,
		Phonetic:     phonetic,
		Type:         wordType,
		CategoryId:   categoryId})
}

func (wu *wordUsecase) Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType *string, categoryId *int) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.Update(ctx, &domainWord.Word{
		ID:           id,
		DictionaryId: dictionaryId,
		Original:     original,
		Phonetic:     phonetic,
		Type:         wordType,
		CategoryId:   categoryId})
}

func (wu *wordUsecase) DeleteById(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.DeleteById(ctx, id)
}
