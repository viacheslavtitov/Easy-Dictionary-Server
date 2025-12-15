package mapper

import (
	dbDictionary "easy-dictionary-server/db/dictionary"
	domainDictionary "easy-dictionary-server/domain/dictionary"
)

func ToDictionaryDomain(d *dbDictionary.DictionaryEntity) *domainDictionary.Dictionary {
	return &domainDictionary.Dictionary{
		ID:         d.ID,
		Dialect:    d.Dialect,
		LangFromId: d.LangFromId,
		LangToId:   d.LangToId,
	}
}

func ToDetailShortDictionaryDomain(d *dbDictionary.DetailShortDictionaryEntity) *domainDictionary.DetailShortDictionary {
	return &domainDictionary.DetailShortDictionary{
		ID:            d.ID,
		Dialect:       d.Dialect,
		LangFrom:      ToLanguageDomain(d.LangFrom),
		LangTo:        ToLanguageDomain(d.LangTo),
		WordTagsCount: d.WordTagCount,
		WordsCount:    d.WordCount,
		QuizCount:     d.QuizCount,
	}
}

func FromDictionaryDomain(d *domainDictionary.Dictionary, userId int) *dbDictionary.DictionaryEntity {
	return &dbDictionary.DictionaryEntity{
		ID:         d.ID,
		Dialect:    d.Dialect,
		LangFromId: d.LangFromId,
		LangToId:   d.LangToId,
		UserId:     userId,
	}
}

func ToDetailDictionaryDomain(d *dbDictionary.DetailDictionaryEntity) *domainDictionary.DetailDictionary {
	return &domainDictionary.DetailDictionary{
		ID:         d.ID,
		Dialect:    d.Dialect,
		LangFrom:   ToLanguageDomain(d.LangFrom),
		LangTo:     ToLanguageDomain(d.LangTo),
		WordTags:   ToWordTagDomainArray(d.WordTags),
		Categories: ToTranslationCategoryDomainArray(d.Categories),
		WordTypes:  d.WordTypes,
		Tenses:     ToTenseDomainArray(d.Tenses),
	}
}
