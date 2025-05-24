package usecase_test

import (
	languageDomain "easy-dictionary-server/domain/language"
	languageDomainMock "easy-dictionary-server/domain/language/mocks"
	testutils "easy-dictionary-server/internalenv/testutils"
	languageUseCase "easy-dictionary-server/usecase/language"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func TestCreate_UnitTest(t *testing.T) {
	mockLRepository := new(languageDomainMock.MockLanguageRepository)
	languageUseCase := languageUseCase.NewLanguageUsecase(mockLRepository, 10)
	err := languageUseCase.Create(testutils.GetTestGinContext(), 1, "English", "en")
	assert.NoError(t, err)
}

func TestUpdate_UnitTest(t *testing.T) {
	mockLRepository := new(languageDomainMock.MockLanguageRepository)
	languageUseCase := languageUseCase.NewLanguageUsecase(mockLRepository, 10)
	err := languageUseCase.Update(testutils.GetTestGinContext(), 1, 1, "English", "en")
	assert.NoError(t, err)
}

func TestGetAllForUser_UnitTest(t *testing.T) {
	mockLRepository := new(languageDomainMock.MockLanguageRepository)
	expectedLanguages := []languageDomain.Language{
		*languageDomainMock.GetMockLanguage(1, "English", "en"),
		*languageDomainMock.GetMockLanguage(1, "Ukrainian", "uk"),
	}
	mockLRepository.On("GetAllUsers", mock.Anything).Return(&expectedLanguages, nil)
	languageUseCase := languageUseCase.NewLanguageUsecase(mockLRepository, 10)
	languages, err := languageUseCase.GetAllForUser(testutils.GetTestGinContext(), 1)
	assert.NoError(t, err)
	assert.Len(t, &languages, 2)
	assert.Equal(t, expectedLanguages[0].Code, (*languages)[0].Code)
}

func TestDeleteById_UnitTest(t *testing.T) {
	mockLRepository := new(languageDomainMock.MockLanguageRepository)
	languageUseCase := languageUseCase.NewLanguageUsecase(mockLRepository, 10)
	err := languageUseCase.DeleteById(testutils.GetTestGinContext(), 1)
	assert.NoError(t, err)
}
