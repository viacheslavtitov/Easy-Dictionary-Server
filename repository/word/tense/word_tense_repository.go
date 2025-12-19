package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbWordTense "easy-dictionary-server/db/word/tense"
	domain "easy-dictionary-server/domain/word/tense"
	wordTenseMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type wordTenseRepository struct {
	db *database.Database
}

func NewWordTagRepository(db *database.Database) domain.WordTenseRepository {
	return &wordTenseRepository{db: db}
}

func (wr *wordTenseRepository) Create(c context.Context, wordTense *domain.WordTense) (int, error) {
	zap.S().Debugf("Create word tense %s for word %d", wordTense.Original, wordTense.WordId)
	id, err := dbWordTense.CreateWordTense(wr.db, c, wordTenseMapper.FromWordTenseDomain(wordTense))
	return id, err
}

func (wr *wordTenseRepository) GetAllWordTenses(c context.Context, wordId int) (*[]domain.WordTense, error) {
	zap.S().Debugf("GetAllWordTenses %d", wordId)
	wordEntities, err := dbWordTense.GetAllWordTensesForWord(wr.db, wordId)
	if err != nil {
		return nil, err
	}
	wordTenses := wordTenseMapper.ToWordTenseDomainArray(wordEntities)
	return wordTenses, nil
}

func (wr *wordTenseRepository) Update(c context.Context, wordTense *domain.WordTense) error {
	zap.S().Debugf("Update word tense %s for word %d", wordTense.Original, wordTense.WordId)
	_, err := dbWordTense.UpdateWordTense(wr.db, wordTenseMapper.FromWordTenseDomain(wordTense))
	return err
}

func (wr *wordTenseRepository) DeleteById(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbWordTense.DeleteWordTenseById(wr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
