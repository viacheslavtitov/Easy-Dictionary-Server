package db_test

import (
	"testing"

	db "easy-dictionary-server/db/language"

	testutils "easy-dictionary-server/internalenv/testutils"
	"fmt"

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

func TestMain(m *testing.M) {
	testutils.SetupTestDB(m)
}

func TestLanguages_Integration(t *testing.T) {
	clearTables()
	testutils.PrepareTestUsersData(t)

	errCreateL1 := db.CreateLanguage(testutils.TestDB, 1, &testLanguageEntity1)
	assert.NoError(t, errCreateL1)
	testLanguageEntity1.ID = 1

	errCreateL2 := db.CreateLanguage(testutils.TestDB, 1, &testLanguageEntity2)
	assert.NoError(t, errCreateL2)
	testLanguageEntity2.ID = 2

	languages, errAllL := db.GetAllLanguagesForUser(testutils.TestDB, 1)
	assert.NoError(t, errAllL)
	assert.Len(t, *languages, 2)
	assert.Equal(t, "English", (*languages)[0].Name)

	testLanguageEntity1.Name = "Deutsch"
	updatedLanguage, errUpL := db.UpdateLanguage(testutils.TestDB, &testLanguageEntity1)
	assert.NoError(t, errUpL)
	assert.Equal(t, "Deutsch", updatedLanguage.Name)

	errDelL := db.DeleteLanguageById(testutils.TestDB, updatedLanguage.ID)
	assert.NoError(t, errDelL)
}

func clearTables() {
	_, err := testutils.TestDB.SQLDB.Exec("TRUNCATE TABLE language RESTART IDENTITY CASCADE")
	if err != nil {
		fmt.Printf("Failed to clear language table %v", err)
	} else {
		fmt.Println("Successfull to clear language table")
	}
}
