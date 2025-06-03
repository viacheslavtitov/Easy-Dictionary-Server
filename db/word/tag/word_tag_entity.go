package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
)

type WordTagEntity struct {
	ID           int    `db:"id"`
	DictionaryId int    `db:"dictionary_id"`
	Name         string `db:"name"`
}

func GetAllWordTagsForDictionary(db *database.Database, dictionaryId int) (*[]WordTagEntity, error) {
	var words []WordTagEntity
	err := db.SQLDB.Select(&words, getAllWordTagsByDictionaryQuery(), dictionaryId)
	if err != nil {
		return nil, err
	}
	return &words, err
}

func CreateWordTag(db *database.Database, entity *WordTagEntity) error {
	_, err := db.SQLDB.Exec(createWordTagQuery(), entity.Name, entity.DictionaryId)
	return err
}

func UpdateWordTag(db *database.Database, entity *WordTagEntity) (*WordTagEntity, error) {
	var word WordTagEntity
	err := db.SQLDB.Get(&word, updateWordTagQuery(), entity.Name, entity.ID)
	if err != nil {
		return nil, err
	}
	return &word, nil
}

func DeleteWordTagById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteWordTagByIdQuery(), id)
	return rowsDeleted, err
}
