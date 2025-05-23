package db_test

import (
	"testing"

	db "easy-dictionary-server/db/language"

	"github.com/stretchr/testify/assert"
)

var testLanguageEntity1 = db.LanguageEntity{
	UserId: 1,
	Name:   "English",
	Code:   "en-US",
}

var testLanguageEntity2 = db.LanguageEntity{
	UserId: 1,
	Name:   "Ukrainian",
	Code:   "uk-UA",
}

func TestLanguages_Integration(t *testing.T) {
	clearTables()

	errCreateL1 := db.CreateLanguage(testDB, 1, &testLanguageEntity1)
	assert.NoError(t, errCreateL1)

	errCreateL2 := db.CreateLanguage(testDB, 1, &testLanguageEntity2)
	assert.NoError(t, errCreateL2)

	languages, errAllL := db.GetAllLanguagesForUser(testDB, 1)
	assert.NoError(t, errAllL)
	assert.Len(t, *languages, 2)
	assert.Equal(t, "English", (*languages)[0].Name)

	testLanguageEntity1.Name = "Deutsch"
	updatedLanguage, errUpL := db.UpdateLanguage(testDB, &testLanguageEntity1)
	assert.NoError(t, errUpL)
	assert.Equal(t, "Deutsch", updatedLanguage.Name)

	errDelL := db.DeleteLanguageById(testDB, updatedLanguage.ID)
	assert.NoError(t, errDelL)
}

func clearTables() {
	testDB.Exec("TRUNCATE TABLE language RESTART IDENTITY CASCADE")
}
