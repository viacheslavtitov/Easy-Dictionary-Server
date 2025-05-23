package db_test

import (
	"testing"

	db "easy-dictionary-server/db/dictionary"

	"github.com/stretchr/testify/assert"
)

var testDictionaryEntity1 = db.DictionaryEntity{
	Dialect:    "en-US",
	LangFromId: 1,
	LangToId:   2,
}

var testDictionaryEntity2 = db.DictionaryEntity{
	Dialect:    "en-GB",
	LangFromId: 3,
	LangToId:   4,
}

func TestDictionaries_Integration(t *testing.T) {
	clearTables()

	errCreateD1 := db.CreateDictionary(testDB, 1, &testDictionaryEntity1)
	assert.NoError(t, errCreateD1)

	errCreateD2 := db.CreateDictionary(testDB, 1, &testDictionaryEntity2)
	assert.NoError(t, errCreateD2)

	dictionaries, errAllD := db.GetAllDictionariesForUser(testDB, 1)
	assert.NoError(t, errAllD)
	assert.Len(t, *dictionaries, 2)
	assert.Equal(t, "en-US", (*dictionaries)[0].Dialect)

	testDictionaryEntity1.Dialect = "en-CA"
	updatedDictionary, errUpD := db.UpdateDictionary(testDB, &testDictionaryEntity1)
	assert.NoError(t, errUpD)
	assert.Equal(t, "en-CA", updatedDictionary.Dialect)

	errDelD := db.DeleteDictionaryById(testDB, updatedDictionary.ID)
	assert.NoError(t, errDelD)
}

func clearTables() {
	testDB.Exec("TRUNCATE TABLE dictionary RESTART IDENTITY CASCADE")
}
