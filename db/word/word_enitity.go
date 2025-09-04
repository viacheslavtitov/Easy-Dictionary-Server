package db

import (
	"context"
	"database/sql"
	database "easy-dictionary-server/db"
	translationEntity "easy-dictionary-server/db/translation"
	translationCategoryDB "easy-dictionary-server/db/translation/category"
	pointers "easy-dictionary-server/internalenv/utils"
	"errors"
	"fmt"
)

type WordEntity struct {
	ID           int     `db:"id"`
	DictionaryId int     `db:"dictionary_id"`
	Original     string  `db:"original"`
	Phonetic     *string `db:"phonetic"`
	Type         *string `db:"type"`
	Translations *[]translationEntity.TranslationEmptyEntity
}

type WordFullEntity struct {
	ID           int     `db:"id"`
	DictionaryId int     `db:"dictionary_id"`
	Original     string  `db:"original"`
	Phonetic     *string `db:"phonetic"`
	Type         *string `db:"type"`
	Translations *[]translationEntity.TranslationWithCategoryEntity
}

type wordFullEntityRow struct {
	ID                     int     `db:"word_id"`
	DictionaryId           int     `db:"word_dictionary_id"`
	Original               string  `db:"word_original"`
	Phonetic               *string `db:"word_phonetic"`
	Type                   *string `db:"word_type"`
	TranslationId          *int    `db:"translation_id"`
	TranslationDescription *string `db:"translation_description"`
	TranslationTranslate   *string `db:"translation_text"`
	CategoryId             *int    `db:"category_id"`
	CategoryName           *string `db:"category_name"`
}

func GetAllWordsForDictionary(db *database.Database, dictionaryId int, lastId int, pageSize int) (*[]WordFullEntity, error) {
	var words []wordFullEntityRow
	err := db.SQLDB.Select(&words, getAllWordsByDictionaryQuery(), dictionaryId, lastId, pageSize)
	if err != nil {
		return nil, err
	}
	return mapWordsFullToEntity(err, words)
}

func SearchWordsForDictionary(db *database.Database, query string, dictionaryId int, lastId int, pageSize int) (*[]WordFullEntity, error) {
	var words []wordFullEntityRow
	err := db.SQLDB.Select(&words, getSearchWordsByDictionaryQuery(), dictionaryId, query, lastId, pageSize)
	if err != nil {
		return nil, err
	}
	return mapWordsFullToEntity(err, words)
}

func mapWordsFullToEntity(err error, rows []wordFullEntityRow) (*[]WordFullEntity, error) {
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, errors.New("Words were not found")
	}
	words := []WordFullEntity{}
	for _, wordEntity := range rows {
		word := WordFullEntity{
			ID:           wordEntity.ID,
			DictionaryId: wordEntity.DictionaryId,
			Original:     wordEntity.Original,
			Phonetic:     wordEntity.Phonetic,
			Type:         wordEntity.Type,
			Translations: &[]translationEntity.TranslationWithCategoryEntity{},
		}
		for _, wordRow := range rows {
			if wordRow.TranslationId != nil {
				translate := translationEntity.TranslationWithCategoryEntity{
					ID:          *wordRow.TranslationId,
					WordId:      word.ID,
					Translate:   pointers.Deref(wordRow.TranslationTranslate),
					Description: wordRow.TranslationDescription,
				}
				if wordRow.CategoryId != nil {
					translate.Category = &translationCategoryDB.TranslationCategoryShortEntity{
						ID:   *wordRow.CategoryId,
						Name: pointers.Deref(wordRow.CategoryName),
					}
				}
				*word.Translations = append(*word.Translations, translate)
			}
		}
		words = append(words, word)
	}
	return &words, err
}

func CreateWord(db *database.Database, dictionaryId int, entity *WordEntity) error {
	_, err := db.SQLDB.Exec(createWordQuery(), entity.Original, entity.Phonetic, entity.Type, dictionaryId)
	return err
}

func CreateWordWithTranslations(db *database.Database, ctx context.Context, dictionaryId int, entity *WordEntity) (int, error) {
	tx, err := db.SQLDB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var wordId int
	err = db.SQLDB.QueryRowContext(ctx, createWordAndReturnIdQuery(), entity.Original, entity.Phonetic, entity.Type, dictionaryId).Scan(&wordId)
	if err != nil {
		return 0, fmt.Errorf("insert word: %w", err)
	}
	stmt, err := tx.PrepareContext(ctx, translationEntity.CreateTranslationQuery())
	if err != nil {
		return 0, fmt.Errorf("prepare translation insert: %w", err)
	}
	defer stmt.Close()
	for _, t := range *entity.Translations {
		if _, err = stmt.ExecContext(ctx, wordId, t.CategoryId, t.Description, t.Translate); err != nil {
			return 0, fmt.Errorf("insert translation: %w", err)
		}
	}
	return wordId, nil
}

func UpdateWord(db *database.Database, entity *WordEntity) (*WordEntity, error) {
	var word WordEntity
	err := db.SQLDB.Get(&word, updateWordQuery(), entity.Original, entity.Phonetic, entity.Type, entity.ID)
	if err != nil {
		return nil, err
	}
	return &word, nil
}

func DeleteWordById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteWordByIdQuery(), id)
	return rowsDeleted, err
}
