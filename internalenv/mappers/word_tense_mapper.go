package mapper

import (
	dbWordTense "easy-dictionary-server/db/word/tense"
	domainWordTense "easy-dictionary-server/domain/word/tense"
)

func ToWordTenseDomain(w *dbWordTense.WordTenseEntity) *domainWordTense.WordTense {
	return &domainWordTense.WordTense{
		ID:       w.ID,
		WordId:   w.WordId,
		TenseId:  w.TenseId,
		Original: w.Original,
		Phonetic: w.Phonetic,
	}
}

func ToWordTenseDomainArray(w *[]dbWordTense.WordTenseEntity) *[]domainWordTense.WordTense {
	if w == nil {
		return &[]domainWordTense.WordTense{}
	}
	var tags []domainWordTense.WordTense
	for _, tag := range *w {
		tags = append(tags, *ToWordTenseDomain(&tag))
	}
	return &tags
}

func FromWordTenseDomainArray(w *[]domainWordTense.WordTense) *[]dbWordTense.WordTenseEntity {
	if w == nil {
		return &[]dbWordTense.WordTenseEntity{}
	}
	var tags []dbWordTense.WordTenseEntity
	for _, tag := range *w {
		tags = append(tags, *FromWordTenseDomain(&tag))
	}
	return &tags
}

func FromWordTenseDomain(w *domainWordTense.WordTense) *dbWordTense.WordTenseEntity {
	return &dbWordTense.WordTenseEntity{
		ID:       w.ID,
		WordId:   w.WordId,
		TenseId:  w.TenseId,
		Original: w.Original,
		Phonetic: w.Phonetic,
	}
}
