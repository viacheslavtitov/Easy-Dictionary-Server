package usecase

import (
	"context"

	domainDictionary "easy-dictionary-server/domain/dictionary"
	commonUseCase "easy-dictionary-server/usecase"
)

type dictionaryUsecase struct {
	dictionaryRepository domainDictionary.DictionaryRepository
	contextTimeout       int
}

func NewDictionaryUsecase(dictionaryRepository domainDictionary.DictionaryRepository, timeout int) domainDictionary.DictionaryUseCase {
	return &dictionaryUsecase{
		dictionaryRepository: dictionaryRepository,
		contextTimeout:       timeout,
	}
}

func (du *dictionaryUsecase) GetDetailForUser(c context.Context, id int) (*domainDictionary.DetailDictionary, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.dictionaryRepository.GetDetailForUser(c, id)
}

func (du *dictionaryUsecase) GetAllForUser(c context.Context, userId int) (*[]domainDictionary.Dictionary, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.dictionaryRepository.GetAllForUser(userId)
}

func (du *dictionaryUsecase) GetAllDetailShortForUser(c context.Context, userId int) (*[]domainDictionary.DetailShortDictionary, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.dictionaryRepository.GetAllDetailShortForUser(userId)
}

func (du *dictionaryUsecase) Create(c context.Context, userId int, dialect *string, langFromId int, langToId int) error {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.dictionaryRepository.Create(userId, domainDictionary.Dictionary{
		Dialect:    dialect,
		LangFromId: langFromId,
		LangToId:   langToId})
}

func (du *dictionaryUsecase) Update(c context.Context, userId int, id int, dialect *string) error {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.dictionaryRepository.Update(userId, id, dialect)
}

func (du *dictionaryUsecase) DeleteById(c context.Context, id int) (int64, error) {
	_, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(du.contextTimeout))
	defer cancel()
	return du.dictionaryRepository.DeleteById(id)
}
