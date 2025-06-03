package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
)

type TranslationEntity struct {
	ID          int     `db:"id"`
	WordId      int     `db:"word_id"`
	CategoryId  *int    `db:"category_id"`
	Translate   string  `db:"translate"`
	Description *string `db:"description"`
}

func GetAllTranslationsForWord(db *database.Database, wordId int) (*[]TranslationEntity, error) {
	var tc []TranslationEntity
	err := db.SQLDB.Select(&tc, getAllTranslationsForWordQuery(), wordId)
	if err != nil {
		return nil, err
	}
	return &tc, err
}

func CreateTranslation(db *database.Database, entity *TranslationEntity) error {
	_, err := db.SQLDB.Exec(createTranslationQuery(), entity.WordId, entity.CategoryId, entity.Translate, entity.Description)
	return err
}

func UpdateTranslation(db *database.Database, entity *TranslationEntity) (*TranslationEntity, error) {
	var tc TranslationEntity
	err := db.SQLDB.Get(&tc, updateTranslationQuery(), entity.CategoryId, entity.Translate, entity.Description, entity.ID)
	if err != nil {
		return nil, err
	}
	return &tc, nil
}

func DeleteTranslationById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteTranslationByIdQuery(), id)
	return rowsDeleted, err
}
