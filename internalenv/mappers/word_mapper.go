package mapper

import (
	translationEntity "easy-dictionary-server/db/translation"
	dbWord "easy-dictionary-server/db/word"
	domainTranslation "easy-dictionary-server/domain/translation"
	domainWord "easy-dictionary-server/domain/word"
)

func ToWordDomain(w *dbWord.WordEntity) *domainWord.Word {
	return &domainWord.Word{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
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

func FromWordWithTranslationDomain(w *domainWord.WordWithTranslations) *dbWord.WordEntity {
	var translations []translationEntity.TranslationEmptyEntity
	for _, wEntity := range translations {
		translations = append(translations, translationEntity.TranslationEmptyEntity{
			CategoryId:  wEntity.CategoryId,
			Translate:   wEntity.Translate,
			Description: wEntity.Description,
		})
	}
	return &dbWord.WordEntity{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
		Translations: &translations,
	}
}

func ToWordWithTranslationAndCategoryDomain(w *dbWord.WordFullEntity, userId int, dictionaryId int) *domainWord.WordWithTranslationsAndCategories {
	var translations []domainTranslation.TranslationWithCategories
	for _, tc := range *w.Translations {
		translations = append(translations, *ToTranslationWithCategoryDomain(&tc, userId, dictionaryId))
	}
	return &domainWord.WordWithTranslationsAndCategories{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
		Translations: &translations,
	}
}
