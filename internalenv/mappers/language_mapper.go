package mapper

import (
	dbLanguage "easy-dictionary-server/db/language"
	domainLanguage "easy-dictionary-server/domain/language"
)

func ToLanguageDomain(l *dbLanguage.LanguageEntity) *domainLanguage.Language {
	return &domainLanguage.Language{
		ID:   l.ID,
		Name: l.Name,
		Code: l.Code,
	}
}

func FromLanguageDomain(l *domainLanguage.Language, userId int) *dbLanguage.LanguageEntity {
	return &dbLanguage.LanguageEntity{
		ID:     l.ID,
		Name:   l.Name,
		Code:   l.Code,
		UserId: userId,
	}
}
