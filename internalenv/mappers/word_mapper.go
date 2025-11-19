package mapper

import (
	translationEntity "easy-dictionary-server/db/translation"
	dbWord "easy-dictionary-server/db/word"
	wordTagEntity "easy-dictionary-server/db/word/tag"
	domainTranslation "easy-dictionary-server/domain/translation"
	domainWord "easy-dictionary-server/domain/word"
	domainWordTag "easy-dictionary-server/domain/word/tag"
)

func ToWordDomain(w *dbWord.WordEntity) *domainWord.Word {
	return &domainWord.Word{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
		CreatedAt:    w.CreatedAt,
	}
}

func FromWordDomain(w *domainWord.Word) *dbWord.WordEntity {
	return &dbWord.WordEntity{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
	}
}

func FromWordDomainToUpdate(id int, dictionaryId int, original string, phonetic *string, wordType *string, tagIds []int) *dbWord.UpdateWordEntity {
	return &dbWord.UpdateWordEntity{
		ID:           id,
		DictionaryId: dictionaryId,
		Original:     original,
		Phonetic:     phonetic,
		Type:         wordType,
		WordTagsIds:  tagIds,
	}
}

func FromWordWithTranslationDomain(w *domainWord.WordWithTranslations) *dbWord.WordEntity {
	var translations []translationEntity.TranslationEmptyEntity
	for _, wEntity := range *w.Translations {
		translations = append(translations, translationEntity.TranslationEmptyEntity{
			CategoryId:  wEntity.CategoryId,
			Translate:   wEntity.Translate,
			Description: wEntity.Description,
		})
	}
	var wordTags []wordTagEntity.WordTagEntity
	for _, wtEntity := range *w.WordTags {
		wordTags = append(wordTags, wordTagEntity.WordTagEntity{
			ID:   wtEntity.ID,
			Name: wtEntity.Name,
		})
	}
	return &dbWord.WordEntity{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
		Translations: &translations,
		WordTags:     &wordTags,
	}
}

func ToWordWithTranslationAndCategoryDomain(w *dbWord.WordFullEntity, userId int, dictionaryId int) *domainWord.WordWithTranslationsAndCategories {
	var translations []domainTranslation.TranslationWithCategories
	for _, tc := range *w.Translations {
		translations = append(translations, *ToTranslationWithCategoryDomain(&tc, userId, dictionaryId))
	}
	var wordTags []domainWordTag.WordTag
	for _, wt := range *w.WordTags {
		wordTags = append(wordTags, *ToWordTagDomain(&wt))
	}
	return &domainWord.WordWithTranslationsAndCategories{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
		Translations: &translations,
		CreatedAt:    w.CreatedAt,
		WordTags:     &wordTags,
	}
}
