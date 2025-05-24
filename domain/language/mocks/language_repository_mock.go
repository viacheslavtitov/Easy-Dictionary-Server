package mocks

import (
	"context"
	languageDomain "easy-dictionary-server/domain/language"

	"github.com/stretchr/testify/mock"
)

type MockLanguageRepository struct {
	mock.Mock
}

func (m *MockLanguageRepository) Create(c context.Context, userId int, language languageDomain.Language) error {
	return nil
}

func (m *MockLanguageRepository) GetAllForUser(c context.Context, userId int) (*[]languageDomain.Language, error) {
	languages := []languageDomain.Language{
		*GetMockLanguage(1, "English", "en"),
		*GetMockLanguage(2, "Ukrainian", "uk"),
	}
	return &languages, nil
}

func (m *MockLanguageRepository) Update(c context.Context, userId int, language languageDomain.Language) error {
	return nil
}

func (m *MockLanguageRepository) DeleteById(c context.Context, id int) error {
	return nil
}
