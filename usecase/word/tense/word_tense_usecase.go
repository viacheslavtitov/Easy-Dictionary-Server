package usecase

import (
	"context"

	domainWordTense "easy-dictionary-server/domain/word/tense"
	commonUseCase "easy-dictionary-server/usecase"
)

type wordTenseUsecase struct {
	wordTenseRepository domainWordTense.WordTenseRepository
	contextTimeout      int
}

func NewWordTenseUsecase(wordTenseRepository domainWordTense.WordTenseRepository, timeout int) domainWordTense.WordTenseUseCase {
	return &wordTenseUsecase{
		wordTenseRepository: wordTenseRepository,
		contextTimeout:      timeout,
	}
}

func (wu *wordTenseUsecase) GetAllWordTenses(c context.Context, wordId int) (*[]domainWordTense.WordTense, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTenseRepository.GetAllWordTenses(ctx, wordId)
}

func (wu *wordTenseUsecase) Create(c context.Context, wordId int, tenseId int, original string, phonetic *string) (int, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTenseRepository.Create(ctx, &domainWordTense.WordTense{
		WordId:   wordId,
		TenseId:  tenseId,
		Original: original,
		Phonetic: phonetic})
}

func (wu *wordTenseUsecase) Update(c context.Context, id int, wordId int, tenseId int, original string, phonetic *string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTenseRepository.Update(ctx, &domainWordTense.WordTense{
		WordId:   wordId,
		TenseId:  tenseId,
		Original: original,
		Phonetic: phonetic})
}

func (wu *wordTenseUsecase) DeleteById(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordTenseRepository.DeleteById(ctx, id)
}
