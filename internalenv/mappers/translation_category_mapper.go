package mapper

import (
	dbTranslationCategory "easy-dictionary-server/db/translation/category"
	domainTranslationCategory "easy-dictionary-server/domain/translation/category"
)

func ToTranslationCategoryDomain(tc *dbTranslationCategory.TranslationCategoryEntity) *domainTranslationCategory.TranslationCategory {
	if tc == nil {
		return nil
	}
	return &domainTranslationCategory.TranslationCategory{
		ID:           tc.ID,
		Name:         tc.Name,
		DictionaryId: tc.DictionaryId,
		UserId:       tc.UserId,
	}
}

func ToShortTranslationCategoryDomain(tc *dbTranslationCategory.TranslationCategoryShortEntity) *domainTranslationCategory.ShortTranslationCategory {
	if tc == nil {
		return nil
	}
	return &domainTranslationCategory.ShortTranslationCategory{
		ID:   tc.ID,
		Name: tc.Name,
	}
}

func ToTranslationCategoryDomainArray(tc *[]dbTranslationCategory.TranslationCategoryShortEntity) *[]domainTranslationCategory.ShortTranslationCategory {
	if tc == nil {
		return &[]domainTranslationCategory.ShortTranslationCategory{}
	}
	var categories []domainTranslationCategory.ShortTranslationCategory
	for _, cat := range *tc {
		categories = append(categories, *ToShortTranslationCategoryDomain(&cat))
	}
	return &categories
}

func FromTranslationCategoryDomain(tc *domainTranslationCategory.TranslationCategory, userId int) *dbTranslationCategory.TranslationCategoryEntity {
	if tc == nil {
		return nil
	}
	return &dbTranslationCategory.TranslationCategoryEntity{
		ID:           tc.ID,
		Name:         tc.Name,
		DictionaryId: tc.DictionaryId,
		UserId:       tc.UserId,
	}
}

func ToTranslationCategoryShortDomain(tc *dbTranslationCategory.TranslationCategoryShortEntity, userId int, dictionaryId int) *domainTranslationCategory.TranslationCategory {
	if tc == nil {
		return nil
	}
	return &domainTranslationCategory.TranslationCategory{
		ID:           tc.ID,
		Name:         tc.Name,
		DictionaryId: dictionaryId,
		UserId:       userId,
	}
}

func ToTranslationCategoryResponseDomain(tc *domainTranslationCategory.TranslationCategory) *domainTranslationCategory.TranslationCategoryResponse {
	if tc == nil {
		return nil
	}
	return &domainTranslationCategory.TranslationCategoryResponse{
		ID:           tc.ID,
		Name:         tc.Name,
		DictionaryId: tc.DictionaryId,
	}
}
