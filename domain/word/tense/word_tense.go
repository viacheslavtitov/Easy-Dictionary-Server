package domain

import (
	"context"
)

type WordTense struct {
	ID       int     `json:"id"`
	TenseId  int     `json:"tense_id"`
	WordId   int     `json:"word_id"`
	Original string  `json:"original"`
	Phonetic *string `json:"phonetic"`
}

type WordTenseRequest struct {
	WordId   int     `json:"word_id"`
	TenseId  int     `json:"tense_id"`
	Original string  `json:"original"`
	Phonetic *string `json:"phonetic"`
}

type EditWordTenseRequest struct {
	ID       int     `json:"id"`
	WordId   int     `json:"word_id"`
	TenseId  int     `json:"tense_id"`
	Original string  `json:"original"`
	Phonetic *string `json:"phonetic"`
}

type WordTenseUseCase interface {
	GetAllWordTenses(c context.Context, wordId int) (*[]WordTense, error)
	Create(c context.Context, wordId int, tenseId int, original string, phonetic *string) (int, error)
	Update(c context.Context, id int, wordId int, tenseId int, original string, phonetic *string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type WordTenseRepository interface {
	GetAllWordTenses(c context.Context, wordId int) (*[]WordTense, error)
	Create(c context.Context, wordTense *WordTense) (int, error)
	Update(c context.Context, wordTense *WordTense) error
	DeleteById(c context.Context, id int) (int64, error)
}
