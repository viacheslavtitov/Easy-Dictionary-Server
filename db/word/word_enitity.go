package db

import (
	"context"
	"database/sql"
	database "easy-dictionary-server/db"
	translationEntity "easy-dictionary-server/db/translation"
	translationCategoryDB "easy-dictionary-server/db/translation/category"
	wordTagEntity "easy-dictionary-server/db/word/tag"
	pointers "easy-dictionary-server/internalenv/utils"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type WordEntity struct {
	ID           int       `db:"id"`
	DictionaryId int       `db:"dictionary_id"`
	Original     string    `db:"original"`
	Phonetic     *string   `db:"phonetic"`
	Type         *string   `db:"type"`
	CreatedAt    time.Time `db:"created_at"`
	Translations *[]translationEntity.TranslationEmptyEntity
	WordTags     *[]wordTagEntity.WordTagEntity
}

type WordFullEntity struct {
	ID           int       `db:"id"`
	DictionaryId int       `db:"dictionary_id"`
	Original     string    `db:"original"`
	Phonetic     *string   `db:"phonetic"`
	Type         *string   `db:"type"`
	CreatedAt    time.Time `db:"created_at"`
	Translations *[]translationEntity.TranslationWithCategoryEntity
	WordTags     *[]wordTagEntity.WordTagEntity
}

type wordFullEntityRow struct {
	ID                     int       `db:"word_id"`
	DictionaryId           int       `db:"word_dictionary_id"`
	Original               string    `db:"word_original"`
	Phonetic               *string   `db:"word_phonetic"`
	Type                   *string   `db:"word_type"`
	TranslationId          *int      `db:"translation_id"`
	TranslationDescription *string   `db:"translation_description"`
	TranslationTranslate   *string   `db:"translation_text"`
	CategoryId             *int      `db:"category_id"`
	CategoryName           *string   `db:"category_name"`
	WordCreatedAt          time.Time `db:"word_created_at"`
	TranslationCreatedAt   time.Time `db:"translation_created_at"`
	WordTagId              *int      `db:"tag_id"`
	WordTagName            *string   `db:"tag_name"`
}

func GetAllWordsForDictionary(db *database.Database, dictionaryId int, lastId int, pageSize int) (*[]WordFullEntity, error) {
	var words []wordFullEntityRow
	err := db.SQLDB.Select(&words, getAllWordsByDictionaryQuery(), dictionaryId, lastId, pageSize+1)
	if err != nil {
		return nil, err
	}
	return mapWordsFullToEntity(err, words)
}

func SearchWordsForDictionary(db *database.Database, query string, dictionaryId int, lastId int, pageSize int) (*[]WordFullEntity, error) {
	var words []wordFullEntityRow
	err := db.SQLDB.Select(&words, getSearchWordsByDictionaryQuery(), dictionaryId, query, lastId, pageSize+1)
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
		return &[]WordFullEntity{}, nil
	}
	wordsByID := make(map[int]*WordFullEntity, len(rows))
	order := make([]int, 0, len(rows))

	for _, r := range rows {
		w, ok := wordsByID[r.ID]
		if !ok {
			translations := make([]translationEntity.TranslationWithCategoryEntity, 0, 4)
			wordTags := make([]wordTagEntity.WordTagEntity, 0, 4)
			w = &WordFullEntity{
				ID:           r.ID,
				DictionaryId: r.DictionaryId,
				Original:     r.Original,
				Phonetic:     r.Phonetic,
				Type:         r.Type,
				Translations: &translations,
				CreatedAt:    r.WordCreatedAt,
				WordTags:     &wordTags,
			}
			wordsByID[r.ID] = w
			order = append(order, r.ID)
		}

		if r.WordTagId != nil {
			wt := wordTagEntity.WordTagEntity{
				ID:           *r.WordTagId,
				DictionaryId: r.DictionaryId,
				Name:         *r.WordTagName,
			}
			*w.WordTags = append(*w.WordTags, wt)
		}

		if r.TranslationId == nil {
			continue
		}

		t := translationEntity.TranslationWithCategoryEntity{
			ID:          *r.TranslationId,
			WordId:      r.ID,
			Description: r.TranslationDescription,
			CreatedAt:   r.TranslationCreatedAt,
		}
		t.Translate = pointers.Deref(r.TranslationTranslate)

		if r.CategoryId != nil || r.CategoryName != nil {
			cat := &translationCategoryDB.TranslationCategoryShortEntity{
				ID:   pointers.DerefInt(r.CategoryId),
				Name: pointers.Deref(r.CategoryName),
			}
			t.Category = cat
		}

		*w.Translations = append(*w.Translations, t)
	}

	result := make([]WordFullEntity, 0, len(wordsByID))
	for _, id := range order {
		result = append(result, *wordsByID[id])
	}
	return &result, nil
}

func CreateWord(db *database.Database, dictionaryId int, entity *WordEntity) error {
	_, err := db.SQLDB.Exec(createWordQuery(), entity.Original, entity.Phonetic, entity.Type, dictionaryId)
	return err
}

func CreateWordWithTranslations(db *database.Database, ctx context.Context, dictionaryId int, entity *WordEntity) (int, error) {
	zap.S().Debugf("CreateWordWithTranslations for dictionary %d with translations %d", dictionaryId, len(*entity.Translations))
	tx, err := db.SQLDB.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		zap.S().Debugln("Failed to start transaction")
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
	zap.S().Debugf("Word inserted by id %d", wordId)
	stmt, err := tx.PrepareContext(ctx, translationEntity.CreateTranslationQuery())
	if err != nil {
		return 0, fmt.Errorf("prepare translation insert: %w", err)
	}
	defer stmt.Close()
	for _, t := range *entity.Translations {
		if _, err = stmt.ExecContext(ctx, wordId, t.CategoryId, t.Translate, t.Description); err != nil {
			zap.S().Debugln("Failed to insert transaction")
			return 0, fmt.Errorf("insert translation: %w", err)
		}
	}
	if entity.WordTags != nil && len(*entity.WordTags) > 0 {
		stmtTag, err := tx.PrepareContext(ctx, wordTagEntity.AddWordTagToWordQuery())
		if err != nil {
			return 0, fmt.Errorf("prepare word tag insert: %w", err)
		}
		defer stmtTag.Close()
		for _, tag := range *entity.WordTags {
			if _, err = stmtTag.ExecContext(ctx, tag.ID, wordId); err != nil {
				zap.S().Debugln("Failed to insert tag")
				return 0, fmt.Errorf("insert tag: %w", err)
			}
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
