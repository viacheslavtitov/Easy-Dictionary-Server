package repository

import (
	database "easy-dictionary-server/db"
	dbTense "easy-dictionary-server/db/tense"
	domain "easy-dictionary-server/domain/dictionary"
	tenseMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type tenseRepository struct {
	db *database.Database
}

func NewTenseRepository(db *database.Database) domain.TenseRepository {
	return &tenseRepository{db: db}
}

func (dr *tenseRepository) Create(dictionaryId int, name string) error {
	zap.S().Debugf("Create tense for dictionary %d", dictionaryId)
	err := dbTense.CreateTense(dr.db, dictionaryId, name)
	return err
}

func (dr *tenseRepository) GetAllForDictionary(id int) (*[]domain.Tense, error) {
	zap.S().Debugf("GetAllForDictionary %d", id)
	tenseEntities, err := dbTense.GetAllTenseForDictionary(dr.db, id)
	if err != nil {
		return nil, err
	}
	return tenseMapper.ToTenseDomainArray(tenseEntities), nil
}

func (dr *tenseRepository) Update(id int, name string) error {
	zap.S().Debugf("Update tense by id %d", id)
	_, err := dbTense.UpdateTense(dr.db, id, name)
	return err
}

func (dr *tenseRepository) DeleteById(id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbTense.DeleteTenseById(dr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
