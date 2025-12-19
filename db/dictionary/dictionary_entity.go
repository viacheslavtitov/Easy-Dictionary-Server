package db

import (
	"context"
	"database/sql"
	database "easy-dictionary-server/db"
	dbLanguage "easy-dictionary-server/db/language"
	dbTense "easy-dictionary-server/db/tense"
	dbTranslationCategory "easy-dictionary-server/db/translation/category"
	dbWordTag "easy-dictionary-server/db/word/tag"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
	"go.uber.org/zap"
)

type DictionaryEntity struct {
	ID         int     `db:"id"`
	UserId     int     `db:"user_id"`
	Dialect    *string `db:"dialect"`
	LangFromId int     `db:"lang_from_id"`
	LangToId   int     `db:"lang_to_id"`
}

type dictionaryDetailShortRow struct {
	DictionaryID int     `db:"dictionary_id"`
	Dialect      *string `db:"dialect"`
	LangFromID   int     `db:"lang_from_id"`
	LangFromName string  `db:"lang_from_name"`
	LangFromCode *string `db:"lang_from_code"`
	LangToID     int     `db:"lang_to_id"`
	LangToName   string  `db:"lang_to_name"`
	LangToCode   *string `db:"lang_to_code"`
	WordCount    int     `db:"word_count"`
	WordTagCount int     `db:"word_tag_count"`
	QuizCount    int     `db:"quiz_count"`
}

type DetailShortDictionaryEntity struct {
	ID           int                        `json:"id"`
	Dialect      *string                    `json:"dialect"`
	LangFrom     *dbLanguage.LanguageEntity `json:"lang_from"`
	LangTo       *dbLanguage.LanguageEntity `json:"lang_to"`
	WordTagCount int                        `json:"word_tag_count"`
	WordCount    int                        `json:"word_count"`
	QuizCount    int                        `json:"quiz_count"`
}

type detailRow struct {
	ID         int            `db:"dictionary_id"`
	Dialect    *string        `db:"dictionary_dialect"`
	LangFrom   []byte         `db:"lang_from"`  // JSON
	LangTo     []byte         `db:"lang_to"`    // JSON
	WordTags   []byte         `db:"word_tags"`  // JSON array
	Categories []byte         `db:"categories"` // JSON array
	Tenses     []byte         `db:"tenses"`     // JSON array
	WordTypes  pq.StringArray `db:"word_types"` // text[]
}

type DetailDictionaryEntity struct {
	ID         int                                                     `json:"id"`
	Dialect    *string                                                 `json:"dialect"`
	LangFrom   *dbLanguage.LanguageEntity                              `json:"lang_from"`
	LangTo     *dbLanguage.LanguageEntity                              `json:"lang_to"`
	WordTags   *[]dbWordTag.WordTagEntity                              `json:"tags"`
	Categories *[]dbTranslationCategory.TranslationCategoryShortEntity `json:"categories"`
	Tenses     *[]dbTense.TenseEntity                                  `json:"tenses"`
	WordTypes  *[]string                                               `json:"types"`
}

func GetAllDictionariesForUser(db *database.Database, userId int) (*[]DictionaryEntity, error) {
	var dictionaries []DictionaryEntity
	err := db.SQLDB.Select(&dictionaries, getAllDictionariesForUserQuery(), userId)
	if err != nil {
		return nil, err
	}
	return &dictionaries, err
}

func CreateDictionary(db *database.Database, ctx context.Context, userId int, entity *DictionaryEntity, tenses []string) (int, error) {
	zap.S().Debugf("CreateDictionary for user %d with tenses %d", userId, len(tenses))
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
	var dictionaryId int
	err = db.SQLDB.QueryRowContext(ctx, createUserDictionaryQuery(), entity.Dialect, entity.LangFromId, entity.LangToId, userId).Scan(&dictionaryId)
	if err != nil {
		return 0, err
	}
	zap.S().Debugf("Dictionary inserted by id %d", dictionaryId)
	stmt, err := tx.PrepareContext(ctx, dbTense.CreateTenseQuery())
	if err != nil {
		return 0, fmt.Errorf("prepare translation insert: %w", err)
	}
	defer stmt.Close()
	for _, t := range tenses {
		if _, err = stmt.ExecContext(ctx, dictionaryId, t); err != nil {
			zap.S().Debugln("Failed to insert transaction")
			return 0, fmt.Errorf("insert translation: %w", err)
		}
	}
	return dictionaryId, err
}

func UpdateDictionary(db *database.Database, id int, dialect *string, tenses []string) error {
	tx, err := db.SQLDB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// update dictionary
	res, err := tx.Exec(updateUserDictionaryQuery(), dialect, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("dictionary not found or not belongs to dictionary")
	}

	if len(tenses) > 0 {
		// add new tenses for dictionary
		_, err = tx.Exec(dbTense.BulkInsertTensesForDictionaryQuery(), id, pq.Array(tenses))
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}

func DeleteDictionaryById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteUserDictionaryByIdQuery(), id)
	return rowsDeleted, err
}

func GetAllDetailShortForUser(db *database.Database, userId int) (*[]DetailShortDictionaryEntity, error) {
	var dictionaryRows []dictionaryDetailShortRow
	err := db.SQLDB.Select(&dictionaryRows, getAllDictionariesWithShortInfoForUserQuery(), userId)
	if err != nil {
		return nil, err
	}
	if len(dictionaryRows) == 0 {
		return &[]DetailShortDictionaryEntity{}, nil
	}
	dictionaryMap := make(map[int]*DetailShortDictionaryEntity)
	for _, row := range dictionaryRows {
		dictionary, exists := dictionaryMap[row.DictionaryID]
		if !exists {
			dictionary = &DetailShortDictionaryEntity{
				ID:           row.DictionaryID,
				Dialect:      row.Dialect,
				WordTagCount: row.WordTagCount,
				WordCount:    row.WordCount,
				QuizCount:    row.QuizCount,
				LangFrom: &dbLanguage.LanguageEntity{
					ID:   row.LangFromID,
					Name: row.LangFromName,
					Code: row.LangFromCode,
				},
				LangTo: &dbLanguage.LanguageEntity{
					ID:   row.LangToID,
					Name: row.LangToName,
					Code: row.LangToCode,
				},
			}
			dictionaryMap[row.DictionaryID] = dictionary
		}
	}
	dictionaries := make([]DetailShortDictionaryEntity, 0, len(dictionaryMap))
	for _, u := range dictionaryMap {
		dictionaries = append(dictionaries, *u)
	}
	return &dictionaries, err
}

func GetDetailForUser(db *database.Database, dictionaryId int) (*DetailDictionaryEntity, error) {
	var row detailRow
	err := db.SQLDB.Get(&row, getDetailDictionaryForUserQuery(), dictionaryId)
	if err != nil {
		return nil, err
	}
	var detail DetailDictionaryEntity
	detail.ID = row.ID
	detail.Dialect = row.Dialect

	// parse LangFrom / LangTo
	var lf dbLanguage.LanguageEntity
	json.Unmarshal(row.LangFrom, &lf)
	detail.LangFrom = &lf

	var lt dbLanguage.LanguageEntity
	json.Unmarshal(row.LangTo, &lt)
	detail.LangTo = &lt

	// tags
	var tags []dbWordTag.WordTagEntity
	json.Unmarshal(row.WordTags, &tags)
	detail.WordTags = &tags

	// categories
	var cats []dbTranslationCategory.TranslationCategoryShortEntity
	json.Unmarshal(row.Categories, &cats)
	detail.Categories = &cats

	// tenses
	var tenses []dbTense.TenseEntity
	json.Unmarshal(row.Tenses, &tenses)
	detail.Tenses = &tenses

	// word_types
	wt := []string(row.WordTypes)
	detail.WordTypes = &wt

	return &detail, nil
}
