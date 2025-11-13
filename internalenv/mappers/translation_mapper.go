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
		CreatedAt:   tc.CreatedAt,
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

func ToTranslationWithCategoryDomain(tc *dbTranslation.TranslationWithCategoryEntity, userId int, dictionaryId int) *domainTranslation.TranslationWithCategories {
	return &domainTranslation.TranslationWithCategories{
		ID:          tc.ID,
		WordId:      tc.WordId,
		Category:    ToTranslationCategoryShortDomain(tc.Category, userId, dictionaryId),
		Translate:   tc.Translate,
		Description: tc.Description,
		CreatedAt:   tc.CreatedAt,
	}
}
