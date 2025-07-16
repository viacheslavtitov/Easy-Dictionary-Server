package domain

import (
	"context"
	languageDomain "easy-dictionary-server/domain/language"
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

type DetailShortDictionary struct {
	ID            int                      `json:"id"`
	Dialect       string                   `json:"dialect"`
	LangFrom      *languageDomain.Language `json:"lang_from"`
	LangTo        *languageDomain.Language `json:"lang_to"`
	WordTagsCount int                      `json:"word_tags_count"`
	WordsCount    int                      `json:"words_count"`
	QuizCount     int                      `json:"quiz_count"`
}

type DictionaryUseCase interface {
	GetAllForUser(c context.Context, userId int) (*[]Dictionary, error)
	GetAllDetailShortForUser(c context.Context, userId int) (*[]DetailShortDictionary, error)
	Create(c context.Context, userId int, dialect string, langFromId int, langToId int) error
	Update(c context.Context, userId int, id int, dialect string, langFromId int, langToId int) error
	DeleteById(c context.Context, id int) (int64, error)
}

type DictionaryRepository interface {
	GetAllForUser(userId int) (*[]Dictionary, error)
	GetAllDetailShortForUser(userId int) (*[]DetailShortDictionary, error)
	Create(userId int, dictionary Dictionary) error
	Update(userId int, dictionary Dictionary) error
	DeleteById(id int) (int64, error)
}
