package mapper

import (
	translationEntity "easy-dictionary-server/db/translation"
	dbWord "easy-dictionary-server/db/word"
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
