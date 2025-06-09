package db

import (
	"database/sql"
	database "easy-dictionary-server/db"
	dictionaryDb "easy-dictionary-server/db/dictionary"
	"time"
)

type QuizEntity struct {
	ID           int        `db:"id"`
	DictionaryId int        `db:"dictionary_id"`
	Name         string     `db:"name"`
	Time         *time.Time `db:"time"`
}

type QuizWordEntity struct {
	ID     int `db:"id"`
	QuizId int `db:"quiz_id"`
	WordId int `db:"word_id"`
}

type QuizResultEntity struct {
	ID     int        `db:"id"`
	QuizId int        `db:"quiz_id"`
	WordId int        `db:"word_id"`
	Time   *time.Time `db:"time"`
}

type QuizWordResultEntity struct {
	ID           int     `db:"id"`
	QuizId       int     `db:"quiz_id"`
	WordId       int     `db:"word_id"`
	QuizResultId int     `db:"quiz_result_id"`
	Answer       *string `db:"answer"`
}

type userQuizWithDictionaryRow struct {
	ID                  int        `db:"id"`
	Name                string     `db:"name"`
	Time                *time.Time `db:"time"`
	DictionaryId        int        `db:"dictionary_id"`
	Dialect             string     `db:"dialect"`
	LangFromId          int        `db:"lang_from_id"`
	LangToId            int        `db:"lang_to_id"`
	QuizWordCount       int64      `db:"quiz_word_count"`
	QuizResultCount     int64      `db:"quiz_result_count"`
	QuizWordResultCount int64      `db:"quiz_word_result_count"`
}

type userQuizDetailWithDictionaryRow struct {
	ID                       int        `db:"id"`
	Name                     string     `db:"name"`
	Time                     *time.Time `db:"time"`
	DictionaryId             int        `db:"dictionary_id"`
	Dialect                  string     `db:"dialect"`
	LangFromId               int        `db:"lang_from_id"`
	LangToId                 int        `db:"lang_to_id"`
	QuizWordId               int        `db:"quiz_word_id"`
	QuizWordEntityId         int        `db:"quiz_word_entity_id"`
	QuizResultId             int        `db:"quiz_result_id"`
	QuizResultWordId         int        `db:"quiz_result_word_id"`
	QuizResultTime           *time.Time `db:"quiz_result_time"`
	QuizWordResultId         int        `db:"quiz_word_result_id"`
	QuizWordEntityResultId   int        `db:"quiz_word_entity_result_id"`
	QuizResultEntityResultId int        `db:"quiz_result_entity_result_id"`
	Answer                   string     `db:"answer"`
}

type QuizDetailEntity struct {
	QuizItem            QuizEntity
	DictionaryItem      dictionaryDb.DictionaryEntity
	QuizWordCount       int64
	QuizResultCount     int64
	QuizWordResultCount int64
}

type QuizItemDetailEntity struct {
	QuizItem         QuizEntity
	DictionaryItem   dictionaryDb.DictionaryEntity
	QuizWords        *[]QuizWordEntity
	QuizResult       *QuizResultEntity
	QuizWordsResults *[]QuizWordResultEntity
}

func GetAllResultsByAllQuiz(db *database.Database, userId int) (*[]QuizDetailEntity, error) {
	var rows []userQuizWithDictionaryRow
	err := db.SQLDB.Select(&rows, getAllResultsByAllQuiz(), userId)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &[]QuizDetailEntity{}, nil
	}
	quizMap := make(map[int]*QuizDetailEntity)
	for _, row := range rows {
		quiz, exists := quizMap[row.ID]
		if !exists {
			quiz = &QuizDetailEntity{
				QuizItem: QuizEntity{
					ID:           row.ID,
					DictionaryId: row.DictionaryId,
					Name:         row.Name,
					Time:         row.Time,
				},
				DictionaryItem: dictionaryDb.DictionaryEntity{
					ID:         row.DictionaryId,
					UserId:     userId,
					Dialect:    row.Dialect,
					LangFromId: row.LangFromId,
					LangToId:   row.LangToId,
				},
				QuizWordCount:       row.QuizWordCount,
				QuizResultCount:     row.QuizResultCount,
				QuizWordResultCount: row.QuizWordResultCount,
			}
			quizMap[row.ID] = quiz
		}
	}
	quizes := make([]QuizDetailEntity, 0, len(quizMap))
	for _, u := range quizMap {
		quizes = append(quizes, *u)
	}
	return &quizes, err
}

func GetAllResultsByQuizId(db *database.Database, userId int, quizId int) (*QuizItemDetailEntity, error) {
	var rows []userQuizDetailWithDictionaryRow
	err := db.SQLDB.Select(&rows, getAllResultsByQuizId(), quizId)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return &QuizItemDetailEntity{}, nil
	}
	quizMap := make(map[int]*QuizItemDetailEntity)
	for _, row := range rows {
		quiz, exists := quizMap[row.ID]
		if !exists {
			quiz = &QuizItemDetailEntity{
				QuizItem: QuizEntity{
					ID:           row.ID,
					DictionaryId: row.DictionaryId,
					Name:         row.Name,
					Time:         row.Time,
				},
				DictionaryItem: dictionaryDb.DictionaryEntity{
					ID:         row.DictionaryId,
					UserId:     userId,
					Dialect:    row.Dialect,
					LangFromId: row.LangFromId,
					LangToId:   row.LangToId,
				},
			}
			quizMap[row.ID] = quiz
		}
		quizWordExists := hasQuizWord(quiz, row.QuizWordId)
		if !quizWordExists {
			if quiz.QuizWords == nil {
				quiz.QuizWords = &[]QuizWordEntity{}
			}
			quizWord := &QuizWordEntity{
				ID:     row.QuizWordId,
				QuizId: row.ID,
				WordId: row.QuizWordEntityId,
			}
			*quiz.QuizWords = append(*quiz.QuizWords, *quizWord)
		}
		quizWordResultExists := hasQuizWordResult(quiz, row.QuizWordResultId)
		if !quizWordResultExists {
			if quiz.QuizWordsResults == nil {
				quiz.QuizWordsResults = &[]QuizWordResultEntity{}
			}
			quizWordResult := &QuizWordResultEntity{
				ID:           row.QuizWordId,
				QuizId:       row.ID,
				WordId:       row.QuizWordEntityId,
				QuizResultId: row.QuizResultId,
				Answer:       &row.Answer,
			}
			*quiz.QuizWordsResults = append(*quiz.QuizWordsResults, *quizWordResult)
		}
		if quiz.QuizResult == nil && row.QuizResultEntityResultId > 0 {
			quiz.QuizResult = &QuizResultEntity{
				ID:     row.QuizResultEntityResultId,
				QuizId: quizId,
				WordId: row.QuizResultWordId,
			}
		}
	}
	if len(quizMap) > 0 {
		return quizMap[0], err
	} else {
		return &QuizItemDetailEntity{}, err
	}
}

func hasQuizWord(quizItem *QuizItemDetailEntity, idToFind int) bool {
	if quizItem.QuizWords == nil {
		return false
	}

	for _, q := range *quizItem.QuizWords {
		if q.ID == idToFind {
			return true
		}
	}
	return false
}

func hasQuizWordResult(quizItem *QuizItemDetailEntity, idToFind int) bool {
	if quizItem.QuizWordsResults == nil {
		return false
	}

	for _, q := range *quizItem.QuizWordsResults {
		if q.ID == idToFind {
			return true
		}
	}
	return false
}

func CreateQuiz(db *database.Database, entity *QuizEntity) (int, error) {
	var quizId int
	err := db.SQLDB.QueryRow(createQuizQuery(), entity.DictionaryId, entity.Name, entity.Time).Scan(&quizId)
	return quizId, err
}

func UpdateQuiz(db *database.Database, entity *QuizEntity) (int, error) {
	var quizId int
	err := db.SQLDB.QueryRow(updateQuizQuery(), entity.ID, entity.Name, entity.Time).Scan(&quizId)
	return quizId, err
}

func CreateQuizWord(db *database.Database, entity *QuizWordEntity) (int, error) {
	var quizWordId int
	err := db.SQLDB.QueryRow(createQuizWordQuery(), entity.QuizId, entity.WordId).Scan(&quizWordId)
	return quizWordId, err
}

func CreateQuizWordResult(db *database.Database, entity *QuizWordResultEntity) (int, error) {
	var quizWordResultId int
	err := db.SQLDB.QueryRow(createQuizWordResultQuery(), entity.QuizResultId, entity.WordId, entity.Answer).Scan(&quizWordResultId)
	return quizWordResultId, err
}

func DeleteQuizById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteQuizByIdQuery(), id)
	return rowsDeleted, err
}

func DeleteQuizWordById(db *database.Database, id int) (sql.Result, error) {
	rowsDeleted, err := db.SQLDB.Exec(deleteQuizWordByIdQuery(), id)
	return rowsDeleted, err
}
