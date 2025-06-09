package domain

import (
	"context"
	dictionaryDomain "easy-dictionary-server/domain/dictionary"
	"time"
)

type Quiz struct {
	ID           int        `json:"id"`
	DictionaryId int        `json:"dictionary_id"`
	Name         string     `json:"name"`
	Time         *time.Time `json:"time"`
}

type QuizWord struct {
	ID     int `json:"id"`
	WordId int `json:"word_id"`
}

type QuizResult struct {
	ID     int        `json:"id"`
	WordId int        `json:"word_id"`
	Time   *time.Time `json:"time"`
}

type QuizWordResult struct {
	ID           int     `json:"id"`
	WordId       int     `json:"word_id"`
	QuizResultId int     `json:"quiz_result_id"`
	Answer       *string `json:"answer"`
}

type QuizDetail struct {
	QuizItem            Quiz
	DictionaryItem      dictionaryDomain.Dictionary
	QuizWordCount       int64
	QuizResultCount     int64
	QuizWordResultCount int64
}

type QuizItemDetail struct {
	QuizItem         Quiz
	DictionaryItem   dictionaryDomain.Dictionary
	QuizWords        *[]QuizWord
	QuizResult       *QuizResult
	QuizWordsResults *[]QuizWordResult
}

type QuizRequest struct {
	DictionaryId int        `json:"dictionary_id" binding:"required"`
	Name         string     `json:"name" binding:"required"`
	Time         *time.Time `json:"time"`
	WordIds      *[]int     `json:"word_ids"`
}

type EditQuizRequest struct {
	ID   int        `json:"id" binding:"required"`
	Name string     `json:"name" binding:"required"`
	Time *time.Time `json:"time"`
}

type QuizWordRequest struct {
	QuizId int `json:"quiz_id" binding:"required"`
	WordId int `json:"word_id" binding:"required"`
}

type QuizWordRezultRequest struct {
	WordId int     `json:"word_id" binding:"required"`
	Answer *string `json:"answer"`
}

type QuizResultRequest struct {
	QuizId      string                   `json:"quiz_id" binding:"required"`
	WordResults *[]QuizWordRezultRequest `json:"word_results"`
}

type QuizUseCase interface {
	CreateQuiz(context context.Context, dictionaryId int, name string, time *time.Time, worIds *[]int) (int, error)
	AddWordToQuiz(context context.Context, quizId int, wordId int) (int, error)
	UpdateQuiz(context context.Context, id int, name string, time *time.Time) (int, error)
	DeleteQuizById(context context.Context, id int) (int64, error)
	DeleteWordFromQuizById(context context.Context, quizWordId int) (int64, error)
	GetAllResultsByAllQuiz(context context.Context, userId int) ([]*QuizDetail, error)
	GetAllResultsByQuizId(context context.Context, userId int, quizId int) (*QuizItemDetail, error)
}

type QuizRepository interface {
	CreateQuiz(context context.Context, quiz Quiz) (int, error)
	AddWordToQuiz(context context.Context, quizId int, wordId int) (int, error)
	UpdateQuiz(context context.Context, quiz Quiz) (int, error)
	DeleteQuizById(context context.Context, quizId int) (int64, error)
	DeleteWordFromQuizById(context context.Context, quizWord QuizWord) (int64, error)
	GetAllResultsByAllQuiz(context context.Context, userId int) ([]*QuizDetail, error)
	GetAllResultsByQuizId(context context.Context, userId int, quizId int) (*QuizItemDetail, error)
}
