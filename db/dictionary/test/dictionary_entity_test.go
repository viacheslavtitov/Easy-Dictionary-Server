package db_test

import (
	"fmt"
	"testing"

	db "easy-dictionary-server/db/dictionary"

	testutils "easy-dictionary-server/internalenv/testutils"

	"github.com/stretchr/testify/assert"
)

var testDictionaryEntity1 = db.DictionaryEntity{
	UserId:     1,
	Dialect:    "en-US",
	LangFromId: 1,
	LangToId:   2,
}

var testDictionaryEntity2 = db.DictionaryEntity{
	UserId:     1,
	Dialect:    "en-GB",
	LangFromId: 2,
	LangToId:   1,
}

func TestMain(m *testing.M) {
	testutils.SetupTestDB(m)
}

func TestDictionaries_Integration(t *testing.T) {
	clearTables()
	testutils.PrepareTestUsersData(t)
	testutils.PrepareTestLanguageData(t)

	errCreateD1 := db.CreateDictionary(testutils.TestDB, 1, &testDictionaryEntity1)
	assert.NoError(t, errCreateD1)
	testDictionaryEntity1.ID = 1

	errCreateD2 := db.CreateDictionary(testutils.TestDB, 1, &testDictionaryEntity2)
	assert.NoError(t, errCreateD2)
	testDictionaryEntity2.ID = 2

	dictionaries, errAllD := db.GetAllDictionariesForUser(testutils.TestDB, 1)
	assert.NoError(t, errAllD)
	assert.Len(t, *dictionaries, 2)
	assert.Equal(t, "en-US", (*dictionaries)[0].Dialect)

	testDictionaryEntity1.Dialect = "en-CA"
	updatedDictionary, errUpD := db.UpdateDictionary(testutils.TestDB, &testDictionaryEntity1)
	assert.NoError(t, errUpD)
	assert.Equal(t, "en-CA", updatedDictionary.Dialect)

	errDelD := db.DeleteDictionaryById(testutils.TestDB, updatedDictionary.ID)
	assert.NoError(t, errDelD)
}

func clearTables() {
	_, err := testutils.TestDB.SQLDB.Exec("TRUNCATE TABLE dictionary RESTART IDENTITY CASCADE")
	if err != nil {
		fmt.Printf("Failed to clear dictionary table %v", err)
	} else {
		fmt.Println("Successfull to clear dictionary table")
	}
}
