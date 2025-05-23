package db_test

import (
	"testing"

	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	db "easy-dictionary-server/db/user"

	"github.com/stretchr/testify/assert"
)

var testClientEntity1 = db.UserEntity{
	FirstName: "John",
	LastName:  "Doe",
}

var testClientEntity2 = db.UserEntity{
	FirstName: "Jane",
	LastName:  "Doe",
}

var userProviderEntity1 = db.UserProviderEntity{
	ProviderName: "google",
}

var userProviderEntity2 = db.UserProviderEntity{
	ProviderName:   "email",
	HashedPassword: "hashed_password",
	Email:          "example@email.com",
}

func TestUsers_Integration(t *testing.T) {
	clearTables()

	createdId1, errCreateU1 := db.CreateUser(testDB, 1, &testClientEntity1)
	assert.NoError(t, errCreateU1)
	assert.Greater(t, createdId1, 0)

	createdId2, errCreateU2 := db.CreateUser(testDB, 1, &testClientEntity2)
	assert.NoError(t, errCreateU2)
	assert.Greater(t, createdId2, 0)

	users, errAllL := db.GetAllUsers(testDB, database.OrderByASC)
	assert.NoError(t, errAllL)
	assert.Len(t, users, 2)
	assert.Equal(t, "John", users[0].FirstName)
	assert.Equal(t, middleware.Client.VALUE, users[0].Role)

	userByEmail, errByEmail := db.GetUserByEmail(testDB, "example@email.com")
	assert.NoError(t, errByEmail)
	assert.Equal(t, "example@email.com", (*userByEmail.Providers)[0].Email)

	testClientEntity1.FirstName = "Bill"
	errUpU1 := db.UpdateUser(testDB, &testClientEntity1)
	assert.NoError(t, errUpU1)
	assert.Equal(t, "Bill", testClientEntity1.FirstName)

	errDelL := db.DeleteUserById(testDB, userByEmail.ID)
	assert.NoError(t, errDelL)
}

func clearTables() {
	testDB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	testDB.Exec("TRUNCATE TABLE user_providers RESTART IDENTITY CASCADE")
}
