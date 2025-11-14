package usecase

import (
	"context"

	translationDomain "easy-dictionary-server/domain/translation"
	domainWord "easy-dictionary-server/domain/word"
	wordTagDomain "easy-dictionary-server/domain/word/tag"
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

func (wu *wordUsecase) GetAllForDictionary(c context.Context, userId int, dictionaryId int, lastId int, pageSize int) (*domainWord.WordsWithPaginationResponse, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.GetAllForDictionary(ctx, userId, dictionaryId, lastId, pageSize)
}

func (wu *wordUsecase) SearchWordsForDictionary(c context.Context, query string, userId int, dictionaryId int, lastId int, pageSize int) (*domainWord.WordsWithPaginationResponse, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.SearchWordsForDictionary(ctx, query, userId, dictionaryId, lastId, pageSize)
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
	translations *[]translationDomain.TranslationWithoutWordRequest, tags *[]wordTagDomain.WordTag) error {
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
		Translations: &convertedTranslations})
}

func (wu *wordUsecase) Update(c context.Context, id int, dictionaryId int, original string, phonetic *string, wordType *string) error {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.Update(ctx, &domainWord.Word{
		ID:           id,
		DictionaryId: dictionaryId,
		Original:     original,
		Phonetic:     phonetic,
		Type:         wordType})
}

func (wu *wordUsecase) DeleteById(c context.Context, id int) (int64, error) {
	ctx, cancel := context.WithTimeout(c, commonUseCase.ReadWriteTimeOut(wu.contextTimeout))
	defer cancel()
	return wu.wordRepository.DeleteById(ctx, id)
}
