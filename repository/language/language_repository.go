package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbLanguage "easy-dictionary-server/db/language"
	domain "easy-dictionary-server/domain/language"
	languageMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type languageRepository struct {
	db *database.Database
}

func NewLanguageRepository(db *database.Database) domain.LanguageRepository {
	return &languageRepository{db: db}
}

func (lr *languageRepository) Create(c context.Context, userId int, language domain.Language) error {
	zap.S().Debugf("Create language %s %s for user %d", language.Name, language.Code, userId)
	err := dbLanguage.CreateLanguage(lr.db, userId, languageMapper.FromLanguageDomain(&language, userId))
	return err
}

func (lr *languageRepository) GetAllForUser(c context.Context, userId int) (*[]domain.Language, error) {
	zap.S().Debugf("GetAllForUser %d", userId)
	languagesEntity, err := dbLanguage.GetAllLanguagesForUser(lr.db, userId)
	if err != nil {
		return nil, err
	}
	var languages []domain.Language
	for _, language := range *languagesEntity {
		languages = append(languages, *languageMapper.ToLanguageDomain(&language))
	}
	return &languages, nil
}

func (lr *languageRepository) Update(c context.Context, userId int, language domain.Language) error {
	zap.S().Debugf("Update language %s %s for user %d", language.Name, language.Code, userId)
	_, err := dbLanguage.UpdateLanguage(lr.db, languageMapper.FromLanguageDomain(&language, userId))
	return err
}

func (lr *languageRepository) DeleteById(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbLanguage.DeleteLanguageById(lr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
