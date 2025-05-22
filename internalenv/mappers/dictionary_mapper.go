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

func FromDictionaryDomain(d *domainDictionary.Dictionary, userId int) *dbDictionary.DictionaryEntity {
	return &dbDictionary.DictionaryEntity{
		ID:         d.ID,
		Dialect:    d.Dialect,
		LangFromId: d.LangFromId,
		LangToId:   d.LangToId,
		UserId:     userId,
	}
}
