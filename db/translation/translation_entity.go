package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
	translationCategoryDB "easy-dictionary-server/db/translation/category"
	"time"
)

type TranslationEntity struct {
	ID          int       `db:"id"`
	WordId      int       `db:"word_id"`
	CategoryId  *int      `db:"category_id"`
	Translate   string    `db:"translate"`
	Description *string   `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type TranslationEmptyEntity struct {
	CategoryId  *int      `db:"category_id"`
	Translate   string    `db:"translate"`
	Description *string   `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

type TranslationWithCategoryEntity struct {
	ID          int                                                   `db:"id"`
	WordId      int                                                   `db:"word_id"`
	Category    *translationCategoryDB.TranslationCategoryShortEntity `db:"category"`
	Translate   string                                                `db:"translate"`
	Description *string                                               `db:"description"`
	CreatedAt   time.Time                                             `db:"created_at"`
}

func GetAllTranslationsForWord(db *database.Database, wordId int) (*[]TranslationEntity, error) {
	var tc []TranslationEntity
	err := db.SQLDB.Select(&tc, getAllTranslationsForWordQuery(), wordId)
	if err != nil {
		return nil, err
	}
	return &tc, err
}

func CreateTranslation(db *database.Database, entity *TranslationEntity) (*int, error) {
	var creatdId int
	err := db.SQLDB.QueryRow(CreateTranslationQuery(), entity.WordId, entity.CategoryId, entity.Translate, entity.Description).Scan(&creatdId)
	if err != nil {
		return nil, err
	}
	return &creatdId, nil
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
