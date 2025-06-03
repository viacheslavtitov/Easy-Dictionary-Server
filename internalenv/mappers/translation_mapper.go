package mapper

import (
	dbTranslation "easy-dictionary-server/db/translation"
	domainTranslation "easy-dictionary-server/domain/translation"
)

func ToTranslationDomain(tc *dbTranslation.TranslationEntity) *domainTranslation.Translation {
	return &domainTranslation.Translation{
		ID:          tc.ID,
		WordId:      tc.WordId,
		CategoryId:  tc.CategoryId,
		Translate:   tc.Translate,
		Description: tc.Description,
	}
}

func FromTranslationDomain(tc *domainTranslation.Translation) *dbTranslation.TranslationEntity {
	return &dbTranslation.TranslationEntity{
		ID:          tc.ID,
		WordId:      tc.WordId,
		CategoryId:  tc.CategoryId,
		Translate:   tc.Translate,
		Description: tc.Description,
	}
}
