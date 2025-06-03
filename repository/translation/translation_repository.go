package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbtranslation "easy-dictionary-server/db/translation"
	domain "easy-dictionary-server/domain/translation"
	translationMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type translationRepository struct {
	db *database.Database
}

func NewTranslationRepository(db *database.Database) domain.TranslationRepository {
	return &translationRepository{db: db}
}

func (lr *translationRepository) Create(c context.Context, wordId int, translation *domain.Translation) error {
	zap.S().Debugf("Create translation %s for word %d", translation.Translate, wordId)
	err := dbtranslation.CreateTranslation(lr.db, translationMapper.FromTranslationDomain(translation))
	return err
}

func (lr *translationRepository) GetAllForWord(c context.Context, wordId int) (*[]domain.Translation, error) {
	zap.S().Debugf("GetAllForWord %d", wordId)
	tsEntity, err := dbtranslation.GetAllTranslationsForWord(lr.db, wordId)
	if err != nil {
		return nil, err
	}
	var categories []domain.Translation
	for _, tCategory := range *tsEntity {
		categories = append(categories, *translationMapper.ToTranslationDomain(&tCategory))
	}
	return &categories, nil
}

func (lr *translationRepository) Update(c context.Context, translation *domain.Translation) error {
	zap.S().Debugf("Update translation %s for word %d", translation.Translate, translation.WordId)
	_, err := dbtranslation.UpdateTranslation(lr.db, translationMapper.FromTranslationDomain(translation))
	return err
}

func (lr *translationRepository) DeleteById(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbtranslation.DeleteTranslationById(lr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
