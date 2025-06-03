package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbDictionary "easy-dictionary-server/db/dictionary"
	domain "easy-dictionary-server/domain/dictionary"
	dictionaryMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type dictionaryRepository struct {
	db *database.Database
}

func NewDictionaryRepository(db *database.Database) domain.DictionaryRepository {
	return &dictionaryRepository{db: db}
}

func (dr *dictionaryRepository) Create(c context.Context, userId int, dictionary domain.Dictionary) error {
	zap.S().Debugf("Create dictionary for user %d", userId)
	err := dbDictionary.CreateDictionary(dr.db, userId, dictionaryMapper.FromDictionaryDomain(&dictionary, userId))
	return err
}

func (dr *dictionaryRepository) GetAllForUser(c context.Context, userId int) (*[]domain.Dictionary, error) {
	zap.S().Debugf("GetAllForUser %d", userId)
	dictionariesEntity, err := dbDictionary.GetAllDictionariesForUser(dr.db, userId)
	if err != nil {
		return nil, err
	}
	var dictionaries []domain.Dictionary
	for _, dictionary := range *dictionariesEntity {
		dictionaries = append(dictionaries, *dictionaryMapper.ToDictionaryDomain(&dictionary))
	}
	return &dictionaries, nil
}

func (dr *dictionaryRepository) Update(c context.Context, userId int, dictionary domain.Dictionary) error {
	zap.S().Debugf("Update dictionary for user %d", userId)
	_, err := dbDictionary.UpdateDictionary(dr.db, dictionaryMapper.FromDictionaryDomain(&dictionary, userId))
	return err
}

func (dr *dictionaryRepository) DeleteById(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbDictionary.DeleteDictionaryById(dr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
