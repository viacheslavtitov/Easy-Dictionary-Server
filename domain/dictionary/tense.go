package domain

import (
	"context"
)

type Tense struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TenseRequest struct {
	Name         string `json:"name" binding:"required"`
	DictionaryId int    `json:"dictionary_id" binding:"required"`
}

type EditTenseRequest struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type TenseUseCase interface {
	GetAllForDictionary(c context.Context, dictionaryId int) (*[]Tense, error)
	Create(c context.Context, dictionaryId int, name string) error
	Update(c context.Context, id int, name string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type TenseRepository interface {
	GetAllForDictionary(dictionaryId int) (*[]Tense, error)
	Create(dictionaryId int, name string) error
	Update(id int, name string) error
	DeleteById(id int) (int64, error)
}
