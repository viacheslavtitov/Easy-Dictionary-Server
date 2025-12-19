package db

import (
	"context"
	"database/sql"
	database "easy-dictionary-server/db"
	"fmt"
)

type WordTenseEntity struct {
	ID       int     `db:"id"`
	TenseId  int     `db:"tense_id"`
	WordId   int     `db:"word_id"`
	Original string  `db:"original"`
	Phonetic *string `db:"phonetic"`
}

func GetAllWordTensesForWord(db *database.Database, wordId int) (*[]WordTenseEntity, error) {
	var wordTenses []WordTenseEntity
	err := db.SQLDB.Select(&wordTenses, getAllWordTensesByWordQuery(), wordId)
	if err != nil {
		return nil, err
	}
	return &wordTenses, err
}

func CreateWordTense(db *database.Database, ctx context.Context, entity *WordTenseEntity) (int, error) {
	var wordTenseId int
	var err error
	err = db.SQLDB.QueryRowContext(ctx, createWordTenseAndReturnIdQuery(), entity.TenseId, entity.WordId, entity.Original, entity.Phonetic).Scan(&wordTenseId)
	if err != nil {
		return 0, fmt.Errorf("insert word tense: %w", err)
	}
	return wordTenseId, err
}

func UpdateWordTense(db *database.Database, entity *WordTenseEntity) (*WordTenseEntity, error) {
	var wordTense WordTenseEntity
	err := db.SQLDB.Get(&wordTense, updateWordTenseQuery(), entity.Original, entity.Phonetic, entity.ID)
	if err != nil {
		return nil, err
	}
	return &wordTense, nil
}

func DeleteWordTenseById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteWordTenseByIdQuery(), id)
	return rowsDeleted, err
}
