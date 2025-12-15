package mapper

import (
	dbTense "easy-dictionary-server/db/tense"
	domain "easy-dictionary-server/domain/dictionary"
	domainTense "easy-dictionary-server/domain/dictionary"
)

func ToTenseDomain(d *dbTense.TenseEntity) *domainTense.Tense {
	return &domainTense.Tense{
		ID:   d.ID,
		Name: d.Name,
	}
}

func FromTenseDomain(d *domainTense.Tense, dictionaryId int) *dbTense.TenseEntity {
	return &dbTense.TenseEntity{
		ID:           d.ID,
		DictionaryId: dictionaryId,
		Name:         d.Name,
	}
}

func FromTenseDomainArray(tenses *[]domainTense.Tense, dictionaryId int) *[]dbTense.TenseEntity {
	if tenses == nil || len(*tenses) == 0 {
		return &[]dbTense.TenseEntity{}
	}
	entities := make([]dbTense.TenseEntity, 0, len(*tenses))
	for _, entity := range *tenses {
		entities = append(entities, *FromTenseDomain(&entity, dictionaryId))
	}
	return &entities
}

func ToTenseDomainArray(tenseEntities *[]dbTense.TenseEntity) *[]domainTense.Tense {
	if tenseEntities == nil || len(*tenseEntities) == 0 {
		return &[]domain.Tense{}
	}
	tenses := make([]domain.Tense, 0, len(*tenseEntities))
	for _, entity := range *tenseEntities {
		tenses = append(tenses, *ToTenseDomain(&entity))
	}
	return &tenses
}
