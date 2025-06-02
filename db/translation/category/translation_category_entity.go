package db

import (
	database "easy-dictionary-server/db"
)

type TranslationCategoryEntity struct {
	ID           int    `db:"id"`
	DictionaryId int    `db:"dictionary_id"`
	UserId       int    `db:"user_id"`
	Name         string `db:"name"`
}

func GetAllTranslationCategoriesForUser(db *database.Database, userId int) (*[]TranslationCategoryEntity, error) {
	var tc []TranslationCategoryEntity
	err := db.SQLDB.Select(&tc, getAllTranslationCategoriesForUserQuery(), userId)
	if err != nil {
		return nil, err
	}
	return &tc, err
}

func CreateTranslationCategory(db *database.Database, entity *TranslationCategoryEntity) error {
	_, err := db.SQLDB.Exec(createUserTranslationCategoryQuery(), entity.Name, entity.UserId, entity.DictionaryId)
	return err
}

func UpdateTranslationCategory(db *database.Database, entity *TranslationCategoryEntity) (*TranslationCategoryEntity, error) {
	var tc TranslationCategoryEntity
	err := db.SQLDB.Get(&tc, updateUserTranslationCategoryQuery(), entity.Name, entity.ID)
	if err != nil {
		return nil, err
	}
	return &tc, nil
}

func DeleteTranslationCategoryById(db *database.Database, id int) error {
	_, err := db.SQLDB.Exec(deleteUserTranslationCategoryByIdQuery(), id)
	return err
}
