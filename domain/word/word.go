package domain

import (
	"context"
)

type Word struct {
	ID           int     `json:"id"`
	DictionaryId int     `json:"dictionary_id"`
	Original     string  `json:"original"`
	Phonetic     *string `json:"phonetic"`
	Type         *string `json:"type"`
	CategoryId   *int    `json:"category_id"`
}

type WordRequest struct {
	DictionaryId int     `json:"dictionary_id" binding:"required"`
	Original     string  `json:"original" binding:"required"`
	Phonetic     *string `json:"phonetic"`
	Type         *string `json:"type"`
	CategoryId   *int    `json:"category_id"`
}

type EditWordRequest struct {
	ID           int     `json:"id" binding:"required"`
	DictionaryId int     `json:"dictionary_id" binding:"required"`
	Original     string  `json:"original" binding:"required"`
	Phonetic     *string `json:"phonetic"`
	Type         *string `json:"type"`
	CategoryId   *int    `json:"category_id"`
}

type WordsWithPaginationResponse struct {
	Words    []Word `json:"words"`
	PageSize int    `json:"page_size"`
	LatestId int    `json:"latest_id"`
}

type WordUseCase interface {
	GetAllForDictionary(c context.Context, dictionaryId int, lastId int, pageSize int) (*[]Word, error)
	Create(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string, categoryId *int) error
	Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType *string, categoryId *int) error
	DeleteById(c context.Context, id int) (int64, error)
}

type WordRepository interface {
	GetAllForDictionary(c context.Context, dictionaryId int, lastId int, pageSize int) (*[]Word, error)
	Create(c context.Context, dictionaryId int, word *Word) error
	Update(c context.Context, word *Word) error
	DeleteById(c context.Context, id int) (int64, error)
}
