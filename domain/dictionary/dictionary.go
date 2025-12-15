package domain

import (
	"context"
	languageDomain "easy-dictionary-server/domain/language"
	translationCategoryDomain "easy-dictionary-server/domain/translation/category"
	wordTagDomain "easy-dictionary-server/domain/word/tag"
)

type Dictionary struct {
	ID         int     `json:"id"`
	Dialect    *string `json:"dialect"`
	LangFromId int     `json:"lang_from_id"`
	LangToId   int     `json:"lang_to_id"`
}

type DetailDictionary struct {
	ID         int                                                   `json:"id"`
	Dialect    *string                                               `json:"dialect"`
	LangFrom   *languageDomain.Language                              `json:"lang_from"`
	LangTo     *languageDomain.Language                              `json:"lang_to"`
	WordTags   *[]wordTagDomain.WordTag                              `json:"tags"`
	Categories *[]translationCategoryDomain.ShortTranslationCategory `json:"categories"`
	WordTypes  *[]string                                             `json:"word_types"`
	Tenses     *[]Tense                                              `json:"tenses"`
}

type DictionaryRequest struct {
	Dialect    *string  `json:"dialect"`
	Tenses     []string `json:"tenses"`
	LangFromId int      `json:"lang_from_id" binding:"required"`
	LangToId   int      `json:"lang_to_id" binding:"required"`
}

type EditDictionaryRequest struct {
	ID      int      `json:"id" binding:"required"`
	Dialect *string  `json:"dialect"`
	Tenses  []string `json:"tenses"`
}

type DetailShortDictionary struct {
	ID            int                      `json:"id"`
	Dialect       *string                  `json:"dialect"`
	LangFrom      *languageDomain.Language `json:"lang_from"`
	LangTo        *languageDomain.Language `json:"lang_to"`
	WordTagsCount int                      `json:"word_tags_count"`
	WordsCount    int                      `json:"words_count"`
	QuizCount     int                      `json:"quiz_count"`
}

type DictionaryUseCase interface {
	GetDetailForUser(c context.Context, id int) (*DetailDictionary, error)
	GetAllForUser(c context.Context, userId int) (*[]Dictionary, error)
	GetAllDetailShortForUser(c context.Context, userId int) (*[]DetailShortDictionary, error)
	Create(c context.Context, userId int, dialect *string, langFromId int, langToId int, tenses []string) error
	Update(c context.Context, userId int, id int, dialect *string, tenses []string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type DictionaryRepository interface {
	GetDetailForUser(c context.Context, id int) (*DetailDictionary, error)
	GetAllForUser(userId int) (*[]Dictionary, error)
	GetAllDetailShortForUser(userId int) (*[]DetailShortDictionary, error)
	Create(c context.Context, userId int, dictionary Dictionary, tenses []string) error
	Update(userId int, id int, dialect *string, tenses []string) error
	DeleteById(id int) (int64, error)
}
