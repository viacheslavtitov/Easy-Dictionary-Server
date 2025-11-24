package mapper

import (
	dbWordTag "easy-dictionary-server/db/word/tag"
	domainWordTag "easy-dictionary-server/domain/word/tag"
)

func ToWordTagDomain(w *dbWordTag.WordTagEntity) *domainWordTag.WordTag {
	return &domainWordTag.WordTag{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Name:         w.Name,
	}
}

func ToWordTagDomainArray(w *[]dbWordTag.WordTagEntity) *[]domainWordTag.WordTag {
	if w == nil {
		return &[]domainWordTag.WordTag{}
	}
	var tags []domainWordTag.WordTag
	for _, tag := range *w {
		tags = append(tags, *ToWordTagDomain(&tag))
	}
	return &tags
}

func FromWordTagDomain(w *domainWordTag.WordTag) *dbWordTag.WordTagEntity {
	return &dbWordTag.WordTagEntity{
		ID:           w.ID,
		DictionaryId: w.DictionaryId,
		Name:         w.Name,
	}
}
