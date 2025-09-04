package domain

import (
	"context"
	translationDomain "easy-dictionary-server/domain/translation"
)

type Word struct {
	ID           int     `json:"id"`
	DictionaryId int     `json:"dictionary_id"`
	Original     string  `json:"original"`
	Phonetic     *string `json:"phonetic"`
	Type         *string `json:"type"`
}

type WordWithTranslations struct {
	ID           int                              `json:"id"`
	DictionaryId int                              `json:"dictionary_id"`
	Original     string                           `json:"original"`
	Phonetic     *string                          `json:"phonetic"`
	Type         *string                          `json:"type"`
	Translations *[]translationDomain.Translation `json:"translations"`
}

type WordWithTranslationsAndCategories struct {
	ID           int                                            `json:"id"`
	DictionaryId int                                            `json:"dictionary_id"`
	Original     string                                         `json:"original"`
	Phonetic     *string                                        `json:"phonetic"`
	Type         *string                                        `json:"type"`
	Translations *[]translationDomain.TranslationWithCategories `json:"translations"`
}

type WordRequest struct {
	DictionaryId int     `json:"dictionary_id" binding:"required"`
	Original     string  `json:"original" binding:"required"`
	Phonetic     *string `json:"phonetic"`
	Type         *string `json:"type"`
}

type WordWithTranslationRequest struct {
	DictionaryId int                                                `json:"dictionary_id" binding:"required"`
	Original     string                                             `json:"original" binding:"required"`
	Phonetic     *string                                            `json:"phonetic"`
	Type         *string                                            `json:"type"`
	Translations *[]translationDomain.TranslationWithoutWordRequest `json:"translations" binding:"required"`
}

type EditWordRequest struct {
	ID           int     `json:"id" binding:"required"`
	DictionaryId int     `json:"dictionary_id" binding:"required"`
	Original     string  `json:"original" binding:"required"`
	Phonetic     *string `json:"phonetic"`
	Type         *string `json:"type"`
}

type WordsWithPaginationResponse struct {
	Words    []WordWithTranslationsAndCategories `json:"words"`
	PageSize int                                 `json:"page_size"`
	LatestId int                                 `json:"latest_id"`
}

type WordUseCase interface {
	GetAllForDictionary(c context.Context, userId int, dictionaryId int, lastId int, pageSize int) (*[]WordWithTranslationsAndCategories, error)
	SearchWordsForDictionary(c context.Context, query string, userId int, dictionaryId int, lastId int, pageSize int) (*[]WordWithTranslationsAndCategories, error)
	Create(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string) error
	CreateWithTranslations(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string, translations *[]translationDomain.TranslationWithoutWordRequest) error
	Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType *string) error
	DeleteById(c context.Context, id int) (int64, error)
}

type WordRepository interface {
	GetAllForDictionary(c context.Context, userId int, dictionaryId int, lastId int, pageSize int) (*[]WordWithTranslationsAndCategories, error)
	SearchWordsForDictionary(c context.Context, query string, userId int, dictionaryId int, lastId int, pageSize int) (*[]WordWithTranslationsAndCategories, error)
	Create(c context.Context, dictionaryId int, word *Word) error
	CreateWithTranslations(c context.Context, dictionaryId int, word *WordWithTranslations) error
	Update(c context.Context, word *Word) error
	DeleteById(c context.Context, id int) (int64, error)
}
