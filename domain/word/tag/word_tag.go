package domain

import (
	"context"
)

type WordTag struct {
	ID           int    `json:"id"`
	DictionaryId int    `json:"dictionary_id"`
	Name         string `json:"name"`
}

type WordTagRequest struct {
	DictionaryId int    `json:"dictionary_id"`
	Name         string `json:"name"`
}

type EditWordTagRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DictionaryId int    `json:"dictionary_id"`
}

type WordTagUseCase interface {
	GetAllForDictionary(c context.Context, dictionaryId int) (*[]WordTag, error)
	Create(c context.Context, dictionaryId int, name string) error
	Update(c context.Context, id int, dictionaryId int, name string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type WordTagRepository interface {
	GetAllForDictionary(c context.Context, dictionaryId int) (*[]WordTag, error)
	Create(c context.Context, word *WordTag) error
	Update(c context.Context, word *WordTag) error
	DeleteById(c context.Context, id int) (int64, error)
}
