package domain

import (
	"context"
)

type TranslationCategory struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DictionaryId int    `json:"dictionary_id"`
	UserId       int    `json:"user_id"`
}

type TranslationCategoryRequest struct {
	Name         string `json:"name" binding:"required"`
	DictionaryId int    `json:"dictionary_id" binding:"required"`
}

type EditTranslationCategoryRequest struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	DictionaryId int    `json:"dictionary_id" binding:"required"`
}

type TranslationCategoryUseCase interface {
	GetAllForUser(c context.Context, userId int) (*[]TranslationCategory, error)
	Create(c context.Context, userId int, dictionaryId int, name string) error
	Update(c context.Context, userId int, id int, dictionaryId int, name string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type TranslationCategoryRepository interface {
	GetAllForUser(c context.Context, userId int) (*[]TranslationCategory, error)
	Create(c context.Context, userId int, translationCategory *TranslationCategory) error
	Update(c context.Context, userId int, translationCategory *TranslationCategory) error
	DeleteById(c context.Context, id int) (int64, error)
}
