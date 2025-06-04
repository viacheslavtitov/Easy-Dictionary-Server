package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
)

type WordEntity struct {
	ID           int     `db:"id"`
	DictionaryId int     `db:"dictionary_id"`
	Original     string  `db:"original"`
	Phonetic     *string `db:"phonetic"`
	Type         *string `db:"type"`
	CategoryId   *int    `db:"category_id"`
}

func GetAllWordsForDictionary(db *database.Database, dictionaryId int, lastId int, pageSize int) (*[]WordEntity, error) {
	var words []WordEntity
	err := db.SQLDB.Select(&words, getAllWordsByDictionaryQuery(), dictionaryId, lastId, pageSize)
	if err != nil {
		return nil, err
	}
	return &words, err
}

func SearchWordsForDictionary(db *database.Database, query string, dictionaryId int, lastId int, pageSize int) (*[]WordEntity, error) {
	var words []WordEntity
	err := db.SQLDB.Select(&words, getSearchWordsByDictionaryQuery(), dictionaryId, query, lastId, pageSize)
	if err != nil {
		return nil, err
	}
	return &words, err
}

func CreateWord(db *database.Database, dictionaryId int, entity *WordEntity) error {
	_, err := db.SQLDB.Exec(createWordQuery(), entity.Original, entity.Phonetic, entity.Type, entity.CategoryId, dictionaryId)
	return err
}

func UpdateWord(db *database.Database, entity *WordEntity) (*WordEntity, error) {
	var word WordEntity
	err := db.SQLDB.Get(&word, updateWordQuery(), entity.Original, entity.Phonetic, entity.Type, entity.CategoryId, entity.ID)
	if err != nil {
		return nil, err
	}
	return &word, nil
}

func DeleteWordById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteWordByIdQuery(), id)
	return rowsDeleted, err
}
