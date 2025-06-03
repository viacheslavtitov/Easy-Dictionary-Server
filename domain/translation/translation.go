package domain

import (
	"context"
)

type Translation struct {
	ID          int     `json:"id"`
	WordId      int     `json:"word_id"`
	CategoryId  *int    `json:"category_id"`
	Translate   string  `json:"translate"`
	Description *string `json:"description"`
}

type TranslationRequest struct {
	WordId      int     `json:"word_id" binding:"required"`
	CategoryId  *int    `json:"category_id"`
	Translate   string  `json:"translate" binding:"required"`
	Description *string `json:"description"`
}

type EditTranslationRequest struct {
	ID          int     `json:"id" binding:"required"`
	WordId      int     `json:"word_id" binding:"required"`
	CategoryId  *int    `json:"category_id"`
	Translate   string  `json:"translate" binding:"required"`
	Description *string `json:"description"`
}

type TranslationUseCase interface {
	GetAllForWord(c context.Context, wordId int) (*[]Translation, error)
	Create(c context.Context, wordId int, categoryId *int, translate string, description *string) error
	Update(c context.Context, id int, categoryId *int, translate string, description *string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type TranslationRepository interface {
	GetAllForWord(c context.Context, wordId int) (*[]Translation, error)
	Create(c context.Context, wordId int, translation *Translation) error
	Update(c context.Context, translation *Translation) error
	DeleteById(c context.Context, id int) (int64, error)
}
