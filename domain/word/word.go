package domain

import (
	"context"
	translationDomain "easy-dictionary-server/domain/translation"
	wordTagDomain "easy-dictionary-server/domain/word/tag"
	wordTenseDomain "easy-dictionary-server/domain/word/tense"
	"time"
)

type Word struct {
	ID           int       `json:"id"`
	DictionaryId int       `json:"dictionary_id"`
	Original     string    `json:"original"`
	Phonetic     *string   `json:"phonetic"`
	Type         *string   `json:"type"`
	CreatedAt    time.Time `json:"created_at"`
}

type WordWithTranslations struct {
	ID           int                              `json:"id"`
	DictionaryId int                              `json:"dictionary_id"`
	Original     string                           `json:"original"`
	Phonetic     *string                          `json:"phonetic"`
	Type         *string                          `json:"type"`
	CreatedAt    time.Time                        `json:"created_at"`
	Translations *[]translationDomain.Translation `json:"translations"`
	WordTags     *[]wordTagDomain.WordTag         `json:"tags"`
	WordTenses   *[]wordTenseDomain.WordTense     `json:"tenses"`
}

type WordWithTranslationsAndCategories struct {
	ID           int                                            `json:"id"`
	DictionaryId int                                            `json:"dictionary_id"`
	Original     string                                         `json:"original"`
	Phonetic     *string                                        `json:"phonetic"`
	Type         *string                                        `json:"type"`
	CreatedAt    time.Time                                      `json:"created_at"`
	Translations *[]translationDomain.TranslationWithCategories `json:"translations"`
	WordTags     *[]wordTagDomain.WordTag                       `json:"tags"`
	WordTenses   *[]wordTenseDomain.WordTense                   `json:"tenses"`
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
	WordTags     *[]wordTagDomain.WordTag                           `json:"tags"`
	WordTenses   *[]wordTenseDomain.WordTense                       `json:"tenses"`
}

type EditWordRequest struct {
	ID           int     `json:"id" binding:"required"`
	DictionaryId int     `json:"dictionary_id" binding:"required"`
	Original     string  `json:"original" binding:"required"`
	Phonetic     *string `json:"phonetic"`
	Type         *string `json:"type"`
	TagIds       []int   `json:"tags"`
}

type WordsWithPaginationResponse struct {
	Words      []WordWithTranslationsAndCategories `json:"words"`
	NextLastID int                                 `json:"next_last_id"`
	HasMore    bool                                `json:"has_more"`
}

type WordUseCase interface {
	SearchWordsForDictionary(c context.Context, userId int, query string, dictionaryId int, lastId int, pageSize int, createdFrom *time.Time, createdTo *time.Time,
		wordTypes *[]string, categoryIds *[]int, tagIds *[]int) (*WordsWithPaginationResponse, error)
	Create(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string) error
	CreateWithTranslations(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string,
		translations *[]translationDomain.TranslationWithoutWordRequest, tags *[]wordTagDomain.WordTag, tenses *[]wordTenseDomain.WordTense) error
	Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType *string, tagIds []int) error
	DeleteById(c context.Context, id int) (int64, error)
}

type WordRepository interface {
	SearchWordsForDictionary(c context.Context, userId int, query string, dictionaryId int, lastId int, pageSize int, createdFrom *time.Time, createdTo *time.Time,
		wordTypes *[]string, categoryIds *[]int, tagIds *[]int) (*WordsWithPaginationResponse, error)
	Create(c context.Context, dictionaryId int, word *Word) error
	CreateWithTranslations(c context.Context, dictionaryId int, word *WordWithTranslations) error
	Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType *string, tagIds []int) error
	DeleteById(c context.Context, id int) (int64, error)
}
