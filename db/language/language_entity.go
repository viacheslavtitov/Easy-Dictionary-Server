package db

import (
	database "easy-dictionary-server/db"
)

type LanguageEntity struct {
	ID     int    `db:"id"`
	UserId int    `db:"user_id"`
	Name   string `db:"name"`
	Code   string `db:"code"`
}

func GetAllLanguagesForUser(db *database.Database, userId int) (*[]LanguageEntity, error) {
	var languages []LanguageEntity
	err := db.SQLDB.Select(&languages, getAllLanguagesForUserQuery(), userId)
	if err != nil {
		return nil, err
	}
	return &languages, err
}

func CreateLanguage(db *database.Database, userId int, entity *LanguageEntity) error {
	_, err := db.SQLDB.Exec(createUserLanguageQuery(), entity.Name, entity.Code, userId)
	return err
}

func UpdateLanguage(db *database.Database, entity *LanguageEntity) (*LanguageEntity, error) {
	var language LanguageEntity
	err := db.SQLDB.Get(&language, updateUserLanguageQuery(), entity.Name, entity.Code, entity.ID)
	if err != nil {
		return nil, err
	}
	return &language, nil
}

func DeleteLanguageById(db *database.Database, id int) error {
	_, err := db.SQLDB.Exec(deleteUserLanguageByIdQuery(), id)
	return err
}
