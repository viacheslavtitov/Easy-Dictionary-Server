package domain

import (
	"context"
)

type WordTag struct {
	ID           int    `json:"id"`
	DictionaryId int    `json:"dictionary_id"`
	WordId       int    `json:"word_id"`
	Name         string `json:"name"`
}

type WordTagRequest struct {
	DictionaryId int    `json:"dictionary_id"`
	WordId       int    `json:"word_id"`
	Name         string `json:"name"`
}

type EditWordTagRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type WordTagUseCase interface {
	GetAllForDictionary(c context.Context, dictionaryId int) (*[]WordTag, error)
	GetAllForWord(c context.Context, wordId int) (*[]WordTag, error)
	Create(c context.Context, dictionaryId int, wordId int, name string) error
	Update(c context.Context, id int, name string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type WordTagRepository interface {
	GetAllForDictionary(c context.Context, dictionaryId int) (*[]WordTag, error)
	GetAllForWord(c context.Context, wordId int) (*[]WordTag, error)
	Create(c context.Context, word *WordTag) error
	Update(c context.Context, word *WordTag) error
	DeleteById(c context.Context, id int) (int64, error)
}
