package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
)

type TenseEntity struct {
	ID           int    `db:"id"`
	DictionaryId int    `db:"dictionary_id"`
	Name         string `db:"name"`
}

func GetAllTenseForDictionary(db *database.Database, dictionaryId int) (*[]TenseEntity, error) {
	var tenses []TenseEntity
	err := db.SQLDB.Select(&tenses, getAllTensesForDictionaryById(), dictionaryId)
	if err != nil {
		return nil, err
	}
	return &tenses, err
}

func CreateTense(db *database.Database, dictionaryId int, name string) error {
	_, err := db.SQLDB.Exec(CreateTenseQuery(), dictionaryId, name)
	return err
}

func UpdateTense(db *database.Database, id int, name string) (*TenseEntity, error) {
	var tense TenseEntity
	err := db.SQLDB.Get(&tense, updateTenseQuery(), name, id)
	if err != nil {
		return nil, err
	}
	return &tense, nil
}

func DeleteTenseById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteTenseByIdQuery(), id)
	return rowsDeleted, err
}
