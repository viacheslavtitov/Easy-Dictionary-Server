package db

import (
	"context"
	"database/sql"
	database "easy-dictionary-server/db"
	"fmt"
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

func GetAllWordTagsForWord(db *database.Database, wordId int) (*[]WordTagEntity, error) {
	var words []WordTagEntity
	err := db.SQLDB.Select(&words, getAllWordTagsByWordQuery(), wordId)
	if err != nil {
		return nil, err
	}
	return &words, err
}

func CreateWordTag(db *database.Database, ctx context.Context, entity *WordTagEntity) (int, error) {
	var tagId int
	var err error
	err = db.SQLDB.QueryRowContext(ctx, createWordTagQuery(), entity.Name, entity.DictionaryId).Scan(&tagId)
	if err != nil {
		return 0, fmt.Errorf("insert tag: %w", err)
	}
	return tagId, err
}

func AddWordTagToWord(db *database.Database, tagId int, wordId int) error {
	_, err := db.SQLDB.Exec(AddWordTagToWordQuery(), tagId, wordId)
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
