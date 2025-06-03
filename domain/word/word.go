package domain

import (
	"context"
)

type Word struct {
	ID           int     `json:"id"`
	DictionaryId int     `json:"dictionary_id"`
	Original     string  `json:"original"`
	Phonetic     *string `json:"phonetic"`
	Type         string  `json:"type"`
	CategoryId   *string `json:"category_id"`
}

type WordRequest struct {
	DictionaryId int    `json:"dictionary_id" binding:"required"`
	Original     string `json:"original" binding:"required"`
	Phonetic     string `json:"phonetic"`
	Type         string `json:"type" binding:"required"`
	CategoryId   string `json:"category_id"`
}

type EditWordRequest struct {
	ID           int    `json:"id" binding:"required"`
	DictionaryId int    `json:"dictionary_id" binding:"required"`
	Original     string `json:"original" binding:"required"`
	Phonetic     string `json:"phonetic"`
	Type         string `json:"type" binding:"required"`
	CategoryId   string `json:"category_id"`
}

type WordUseCase interface {
	GetAllForDictionary(c context.Context, dictionaryId int) (*[]Word, error)
	Create(c context.Context, dictionaryId int, original string, phonetic *string, wordType string, categoryId *string) error
	Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType string, categoryId *string) error
	DeleteById(c context.Context, id int) error
}

type WordRepository interface {
	GetAllForDictionary(c context.Context, dictionaryId int) (*[]Word, error)
	Create(c context.Context, dictionaryId int, word *Word) error
	Update(c context.Context, word *Word) error
	DeleteById(c context.Context, id int) error
}
