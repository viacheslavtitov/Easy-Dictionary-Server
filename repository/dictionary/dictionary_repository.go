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

func (dr *dictionaryRepository) Create(c context.Context, userId int, dictionary domain.Dictionary, tenses []string) error {
	zap.S().Debugf("Create dictionary for user %d", userId)
	_, err := dbDictionary.CreateDictionary(dr.db, c, userId, dictionaryMapper.FromDictionaryDomain(&dictionary, userId), tenses)
	return err
}

func (dr *dictionaryRepository) GetDetailForUser(c context.Context, id int) (*domain.DetailDictionary, error) {
	zap.S().Debugf("GetDetailForUser %d", id)
	dictionaryEntity, err := dbDictionary.GetDetailForUser(dr.db, id)
	if err != nil {
		return nil, err
	}
	return dictionaryMapper.ToDetailDictionaryDomain(dictionaryEntity), nil
}

func (dr *dictionaryRepository) GetAllForUser(userId int) (*[]domain.Dictionary, error) {
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

func (dr *dictionaryRepository) GetAllDetailShortForUser(userId int) (*[]domain.DetailShortDictionary, error) {
	zap.S().Debugf("GetAllDetailShortForUser %d", userId)
	dictionariesEntity, err := dbDictionary.GetAllDetailShortForUser(dr.db, userId)
	if err != nil {
		return nil, err
	}
	var dictionaries []domain.DetailShortDictionary
	for _, dictionary := range *dictionariesEntity {
		dictionaries = append(dictionaries, *dictionaryMapper.ToDetailShortDictionaryDomain(&dictionary))
	}
	return &dictionaries, nil
}

func (dr *dictionaryRepository) Update(userId int, id int, dialect *string, tenses []string) error {
	zap.S().Debugf("Update dictionary for user %d", userId)
	err := dbDictionary.UpdateDictionary(dr.db, id, dialect, tenses)
	return err
}

func (dr *dictionaryRepository) DeleteById(id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbDictionary.DeleteDictionaryById(dr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
