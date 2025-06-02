package mapper

import (
	dbTranslationCategory "easy-dictionary-server/db/translation/category"
	domainTranslationCategory "easy-dictionary-server/domain/translation/category"
)

func ToTranslationCategoryDomain(tc *dbTranslationCategory.TranslationCategoryEntity) *domainTranslationCategory.TranslationCategory {
	return &domainTranslationCategory.TranslationCategory{
		ID:           tc.ID,
		Name:         tc.Name,
		DictionaryId: tc.DictionaryId,
		UserId:       tc.UserId,
	}
}

func FromTranslationCategoryDomain(tc *domainTranslationCategory.TranslationCategory, userId int) *dbTranslationCategory.TranslationCategoryEntity {
	return &dbTranslationCategory.TranslationCategoryEntity{
		ID:           tc.ID,
		Name:         tc.Name,
		DictionaryId: tc.DictionaryId,
		UserId:       tc.UserId,
	}
}
