package mocks

import (
	"context"
	languageDomain "easy-dictionary-server/domain/language"

	"github.com/stretchr/testify/mock"
)

type MockLanguageUseCase struct {
	mock.Mock
}

func GetMockLanguage(id int, name string, code string) *languageDomain.Language {
	return &languageDomain.Language{
		ID:   id,
		Name: name,
		Code: code,
	}
}

func (m *MockLanguageUseCase) Create(c context.Context, userId int, name string, code string) error {
	return nil
}

func (m *MockLanguageUseCase) GetAllForUser(c context.Context, userId int) (*[]languageDomain.Language, error) {
	languages := []languageDomain.Language{
		*GetMockLanguage(1, "English", "en"),
		*GetMockLanguage(2, "Ukrainian", "uk"),
	}
	return &languages, nil
}

func (m *MockLanguageUseCase) Update(c context.Context, userId int, language languageDomain.Language) error {
	return nil
}

func (m *MockLanguageUseCase) DeleteById(c context.Context, id int) error {
	return nil
}
