package usecase

import (
	"context"
	"time"

	translationDomain "easy-dictionary-server/domain/translation"
	domainWord "easy-dictionary-server/domain/word"
	wordTagDomain "easy-dictionary-server/domain/word/tag"
	wordTenseDomain "easy-dictionary-server/domain/word/tense"
	commonUseCase "easy-dictionary-server/usecase"
)

type wordUsecase struct {
	wordRepository domainWord.WordRepository
	contextTimeout int
}

func NewWordUsecase(wordRepository domainWord.WordRepository, timeout int) domainWord.WordUseCase {
	return &wordUsecase{
		wordRepository: wordRepository,
		contextTimeout: timeout,
	}
}

func (wu *wordUsecase) SearchWordsForDictionary(c context.Context, userId int, query string, dictionaryId int, lastId int, pageSize int, createdFrom *time.Time, createdTo *time.Time,
	wordTypes *[]string, categoryIds *[]int, tagIds *[]int) (*domainWord.WordsWithPaginationResponse, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.SearchWordsForDictionary(ctx, userId, query, dictionaryId, lastId, pageSize, createdFrom, createdTo, wordTypes, categoryIds, tagIds)
}

func (wu *wordUsecase) Create(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.Create(ctx, dictionaryId, &domainWord.Word{
		DictionaryId: dictionaryId,
		Original:     original,
		Phonetic:     phonetic,
		Type:         wordType})
}

func (wu *wordUsecase) CreateWithTranslations(c context.Context, dictionaryId int, original string, phonetic *string, wordType *string,
	translations *[]translationDomain.TranslationWithoutWordRequest, tags *[]wordTagDomain.WordTag, tenses *[]wordTenseDomain.WordTense) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	var convertedTranslations []translationDomain.Translation
	for _, t := range *translations {
		convertedTranslations = append(convertedTranslations, translationDomain.Translation{
			ID:          -1,
			WordId:      -1,
			CategoryId:  t.CategoryId,
			Translate:   t.Translate,
			Description: t.Description,
		})
	}

	return wu.wordRepository.CreateWithTranslations(ctx, dictionaryId, &domainWord.WordWithTranslations{
		DictionaryId: dictionaryId,
		Original:     original,
		Phonetic:     phonetic,
		Type:         wordType,
		WordTags:     tags,
		Translations: &convertedTranslations,
		WordTenses:   tenses})
}

func (wu *wordUsecase) Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType *string, tagIds []int) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.Update(ctx, id, dictionaryId, original, phonetic, wordType, tagIds)
}

func (wu *wordUsecase) DeleteById(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.DeleteById(ctx, id)
}
