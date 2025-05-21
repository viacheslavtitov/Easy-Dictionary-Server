package domain

import (
	"context"
)

type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type LanguageRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type EditLanguageRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type LanguageUseCase interface {
	GetAllForUser(c context.Context, userId int) (*[]Language, error)
	Create(c context.Context, userId int, name string, code string) error
	Update(c context.Context, userId int, id int, name string, code string) error
	DeleteById(c context.Context, id int) error
}

type LanguageRepository interface {
	GetAllForUser(c context.Context, userId int) (*[]Language, error)
	Create(c context.Context, userId int, language Language) error
	Update(c context.Context, userId int, language Language) error
	DeleteById(c context.Context, id int) error
}
