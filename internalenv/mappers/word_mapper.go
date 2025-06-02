package mapper

import (
	dbWord "easy-dictionary-server/db/word"
	domainWord "easy-dictionary-server/domain/word"
)

func ToWordDomain(w *dbWord.WordEntity) *domainWord.Word {
	return &domainWord.Word{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		CategoryId:   w.CategoryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
	}
}

func FromWordDomain(w *domainWord.Word) *dbWord.WordEntity {
	return &dbWord.WordEntity{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		CategoryId:   w.CategoryId,
		Original:     w.Original,
		Phonetic:     w.Phonetic,
		Type:         w.Type,
	}
}
