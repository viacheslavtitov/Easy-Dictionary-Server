package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbWordTag "easy-dictionary-server/db/word/tag"
	domain "easy-dictionary-server/domain/word/tag"
	wordTagMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type wordTagRepository struct {
	db *database.Database
}

func NewWordTagRepository(db *database.Database) domain.WordTagRepository {
	return &wordTagRepository{db: db}
}

func (wr *wordTagRepository) Create(c context.Context, wordTag *domain.WordTag) error {
	zap.S().Debugf("Create word tag %s for dictionary %d", wordTag.Name, wordTag.DictionaryId)
	err := dbWordTag.CreateWordTag(wr.db, wordTagMapper.FromWordTagDomain(wordTag))
	return err
}

func (wr *wordTagRepository) GetAllForDictionary(c context.Context, dictionaryId int) (*[]domain.WordTag, error) {
	zap.S().Debugf("GetAllForDictionary %d", dictionaryId)
	wordEntities, err := dbWordTag.GetAllWordTagsForDictionary(wr.db, dictionaryId)
	if err != nil {
		return nil, err
	}
	var words []domain.WordTag
	for _, wEntity := range *wordEntities {
		words = append(words, *wordTagMapper.ToWordTagDomain(&wEntity))
	}
	return &words, nil
}

func (wr *wordTagRepository) Update(c context.Context, wordTag *domain.WordTag) error {
	zap.S().Debugf("Update word tag %s for dictionary %d", wordTag.Name, wordTag.DictionaryId)
	_, err := dbWordTag.UpdateWordTag(wr.db, wordTagMapper.FromWordTagDomain(wordTag))
	return err
}

func (wr *wordTagRepository) DeleteById(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbWordTag.DeleteWordTagById(wr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
