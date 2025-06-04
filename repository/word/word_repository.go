package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbWord "easy-dictionary-server/db/word"
	domain "easy-dictionary-server/domain/word"
	wordMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type wordRepository struct {
	db *database.Database
}

func NewWordRepository(db *database.Database) domain.WordRepository {
	return &wordRepository{db: db}
}

func (wr *wordRepository) Create(c context.Context, dictionaryId int, word *domain.Word) error {
	zap.S().Debugf("Create word %s for user %d", word.Original, dictionaryId)
	err := dbWord.CreateWord(wr.db, dictionaryId, wordMapper.FromWordDomain(word))
	return err
}

func (wr *wordRepository) GetAllForDictionary(c context.Context, dictionaryId int, lastId int, pageSize int) (*[]domain.Word, error) {
	zap.S().Debugf("GetAllForDictionary %d", dictionaryId)
	wordEntities, err := dbWord.GetAllWordsForDictionary(wr.db, dictionaryId, lastId, pageSize)
	if err != nil {
		return nil, err
	}
	var words []domain.Word
	for _, wEntity := range *wordEntities {
		words = append(words, *wordMapper.ToWordDomain(&wEntity))
	}
	return &words, nil
}

func (wr *wordRepository) SearchWordsForDictionary(c context.Context, query string, dictionaryId int, lastId int, pageSize int) (*[]domain.Word, error) {
	zap.S().Debugf("SearchWordsForDictionary %d %s", dictionaryId, query)
	wordEntities, err := dbWord.SearchWordsForDictionary(wr.db, query, dictionaryId, lastId, pageSize)
	if err != nil {
		return nil, err
	}
	var words []domain.Word
	for _, wEntity := range *wordEntities {
		words = append(words, *wordMapper.ToWordDomain(&wEntity))
	}
	return &words, nil
}

func (wr *wordRepository) Update(c context.Context, word *domain.Word) error {
	zap.S().Debugf("Update word %s for dictionary %d", word.Original, word.DictionaryId)
	_, err := dbWord.UpdateWord(wr.db, wordMapper.FromWordDomain(word))
	return err
}

func (wr *wordRepository) DeleteById(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbWord.DeleteWordById(wr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
