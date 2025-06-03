package repository

import (
	"context"

	database "easy-dictionary-server/db"
	dbTranslationCategory "easy-dictionary-server/db/translation/category"
	domain "easy-dictionary-server/domain/translation/category"
	translationCategoryMapper "easy-dictionary-server/internalenv/mappers"

	"go.uber.org/zap"
)

type translationCategoryRepository struct {
	db *database.Database
}

func NewTranslationCategoryRepository(db *database.Database) domain.TranslationCategoryRepository {
	return &translationCategoryRepository{db: db}
}

func (lr *translationCategoryRepository) Create(c context.Context, userId int, ts *domain.TranslationCategory) error {
	zap.S().Debugf("Create translation category %s for user %d", ts.Name, userId)
	err := dbTranslationCategory.CreateTranslationCategory(lr.db, translationCategoryMapper.FromTranslationCategoryDomain(ts, userId))
	return err
}

func (lr *translationCategoryRepository) GetAllForUser(c context.Context, userId int) (*[]domain.TranslationCategory, error) {
	zap.S().Debugf("GetAllForUser %d", userId)
	tsEntity, err := dbTranslationCategory.GetAllTranslationCategoriesForUser(lr.db, userId)
	if err != nil {
		return nil, err
	}
	var categories []domain.TranslationCategory
	for _, tCategory := range *tsEntity {
		categories = append(categories, *translationCategoryMapper.ToTranslationCategoryDomain(&tCategory))
	}
	return &categories, nil
}

func (lr *translationCategoryRepository) Update(c context.Context, userId int, ts *domain.TranslationCategory) error {
	zap.S().Debugf("Update translation category %s for user %d", ts.Name, userId)
	_, err := dbTranslationCategory.UpdateTranslationCategory(lr.db, translationCategoryMapper.FromTranslationCategoryDomain(ts, userId))
	return err
}

func (lr *translationCategoryRepository) DeleteById(c context.Context, id int) (int64, error) {
	zap.S().Debugf("DeleteById %d", id)
	rowsDeleted, errQuery := dbTranslationCategory.DeleteTranslationCategoryById(lr.db, id)
	if errQuery != nil {
		return 0, errQuery
	}
	deletedRows, err := rowsDeleted.RowsAffected()
	return deletedRows, err
}
