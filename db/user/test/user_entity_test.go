package db_test

import (
	"testing"

	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	db "easy-dictionary-server/db/user"
	testutils "easy-dictionary-server/internalenv/testutils"

	"fmt"

	"github.com/stretchr/testify/assert"
)

var userProviderEntity1 = db.UserProviderEntity{
	ProviderName: "google",
}

var userProviderEntity2 = db.UserProviderEntity{
	ProviderName:   "email",
	HashedPassword: "hashed_password",
	Email:          "example@email.com",
}

var testClientEntity1 = db.UserEntity{
	FirstName: "John",
	LastName:  "Doe",
	Role:      middleware.Client.VALUE,
	Providers: &[]db.UserProviderEntity{userProviderEntity1},
}

var testClientEntity2 = db.UserEntity{
	FirstName: "Jane",
	LastName:  "Doe",
	Role:      middleware.Client.VALUE,
	Providers: &[]db.UserProviderEntity{userProviderEntity2},
}

func TestMain(m *testing.M) {
	testutils.SetupTestDB(m)
}

func TestUsers_Integration(t *testing.T) {
	clearTables()

	createdId1, errCreateU1 := db.CreateUser(testutils.TestDB, &testClientEntity1)
	assert.NoError(t, errCreateU1)
	assert.Greater(t, createdId1, 0)

	createdId2, errCreateU2 := db.CreateUser(testutils.TestDB, &testClientEntity2)
	assert.NoError(t, errCreateU2)
	assert.Greater(t, createdId2, 0)

	users, errAllL := db.GetAllUsers(testutils.TestDB, database.OrderByASC)
	assert.NoError(t, errAllL)
	assert.Len(t, users, 2)
	assert.Equal(t, middleware.Client.VALUE, users[0].Role)

	userByEmail, errByEmail := db.GetUserByEmail(testutils.TestDB, "example@email.com")
	assert.NoError(t, errByEmail)
	assert.Equal(t, "example@email.com", (*userByEmail.Providers)[0].Email)

	testClientEntity1.FirstName = "Bill"
	errUpU1 := db.UpdateUser(testutils.TestDB, &testClientEntity1)
	assert.NoError(t, errUpU1)
	assert.Equal(t, "Bill", testClientEntity1.FirstName)

	errDelL := db.DeleteUserById(testutils.TestDB, userByEmail.ID)
	assert.NoError(t, errDelL)
}

func clearTables() {
	_, err := testutils.TestDB.SQLDB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
	if err != nil {
		fmt.Printf("Failed to clear users table %v", err)
	} else {
		fmt.Println("Successfull to clear users table")
	}
	// testutils.TestDB.SQLDB.Exec("TRUNCATE TABLE user_providers RESTART IDENTITY CASCADE")
}
