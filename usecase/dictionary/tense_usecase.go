package usecase

import (
	"context"

	domainTense "easy-dictionary-server/domain/dictionary"
	commonUseCase "easy-dictionary-server/usecase"
)

type tenseUsecase struct {
	tenseRepository domainTense.TenseRepository
	contextTimeout  int
}

func NewTenseUsecase(tenseRepository domainTense.TenseRepository, timeout int) domainTense.TenseUseCase {
	return &tenseUsecase{
		tenseRepository: tenseRepository,
		contextTimeout:  timeout,
	}
}

func (du *tenseUsecase) GetAllForDictionary(c context.Context, dictionaryId int) (*[]domainTense.Tense, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.tenseRepository.GetAllForDictionary(dictionaryId)
}

func (du *tenseUsecase) Create(c context.Context, dictionaryId int, name string) error {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.tenseRepository.Create(dictionaryId, name)
}

func (du *tenseUsecase) Update(c context.Context, id int, name string) error {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.tenseRepository.Update(id, name)
}

func (du *tenseUsecase) DeleteById(c context.Context, id int) (int64, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.tenseRepository.DeleteById(id)
}
