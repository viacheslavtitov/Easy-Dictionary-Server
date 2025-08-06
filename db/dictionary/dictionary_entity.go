package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
	dbLanguage "easy-dictionary-server/db/language"
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

func UpdateDictionary(db *database.Database, id int, dialect *string) (*DictionaryEntity, error) {
	var dictionary DictionaryEntity
	err := db.SQLDB.Get(&dictionary, updateUserDictionaryQuery(), dialect, id)
	if err != nil {
		return nil, err
	}
	return &dictionary, nil
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
