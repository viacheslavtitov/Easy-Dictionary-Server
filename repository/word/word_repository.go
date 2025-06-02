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
	err := dbWord.CreateWord(wr.db, wordMapper.FromWordDomain(word))
	return err
}

func (wr *wordRepository) GetAllForDictionary(c context.Context, dictionaryId int) (*[]domain.Word, error) {
	zap.S().Debugf("GetAllForDictionary %d", dictionaryId)
	wordEntities, err := dbWord.GetAllWordsForDictionary(wr.db, dictionaryId)
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

func (wr *wordRepository) DeleteById(c context.Context, id int) error {
	zap.S().Debugf("DeleteById %d", id)
	err := dbWord.DeleteWordById(wr.db, id)
	return err
}
