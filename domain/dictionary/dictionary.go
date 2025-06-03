package domain

import (
	"context"
)

type Dictionary struct {
	ID         int    `json:"id"`
	Dialect    string `json:"dialect"`
	LangFromId int    `json:"lang_from_id"`
	LangToId   int    `json:"lang_to_id"`
}

type DictionaryRequest struct {
	Dialect    string `json:"dialect" binding:"required"`
	LangFromId int    `json:"lang_from_id" binding:"required"`
	LangToId   int    `json:"lang_to_id" binding:"required"`
}

type EditDictionaryRequest struct {
	ID         int    `json:"id" binding:"required"`
	Dialect    string `json:"dialect" binding:"required"`
	LangFromId int    `json:"lang_from_id" binding:"required"`
	LangToId   int    `json:"lang_to_id" binding:"required"`
}

type DictionaryUseCase interface {
	GetAllForUser(c context.Context, userId int) (*[]Dictionary, error)
	Create(c context.Context, userId int, dialect string, langFromId int, langToId int) error
	Update(c context.Context, userId int, id int, dialect string, langFromId int, langToId int) error
	DeleteById(c context.Context, id int) (int64, error)
}

type DictionaryRepository interface {
	GetAllForUser(c context.Context, userId int) (*[]Dictionary, error)
	Create(c context.Context, userId int, dictionary Dictionary) error
	Update(c context.Context, userId int, dictionary Dictionary) error
	DeleteById(c context.Context, id int) (int64, error)
}
