package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
)

type DictionaryEntity struct {
	ID         int    `db:"id"`
	UserId     int    `db:"user_id"`
	Dialect    string `db:"dialect"`
	LangFromId int    `db:"lang_from_id"`
	LangToId   int    `db:"lang_to_id"`
}

func GetAllDictionariesForUser(db *database.Database, userId int) (*[]DictionaryEntity, error) {
	var dictionaries []DictionaryEntity
	err := db.SQLDB.Select(&dictionaries, getAllDictionariesForUserQuery(), userId)
	if err != nil {
		return nil, err
	}
	return &dictionaries, err
}

func CreateDictionary(db *database.Database, userId int, entity *DictionaryEntity) error {
	_, err := db.SQLDB.Exec(createUserDictionaryQuery(), entity.Dialect, entity.LangFromId, entity.LangToId, userId)
	return err
}

func UpdateDictionary(db *database.Database, entity *DictionaryEntity) (*DictionaryEntity, error) {
	var dictionary DictionaryEntity
	err := db.SQLDB.Get(&dictionary, updateUserDictionaryQuery(), entity.Dialect, entity.LangFromId, entity.LangToId, entity.ID)
	if err != nil {
		return nil, err
	}
	return &dictionary, nil
}

func DeleteDictionaryById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteUserDictionaryByIdQuery(), id)
	return rowsDeleted, err
}
