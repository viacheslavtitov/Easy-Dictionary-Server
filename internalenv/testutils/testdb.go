package testutils

import (
	"os"
	"testing"

	middleware "easy-dictionary-server/api/middleware"
	database "easy-dictionary-server/db"
	dbLanguage "easy-dictionary-server/db/language"
	dbUser "easy-dictionary-server/db/user"
	utils "easy-dictionary-server/internalenv/utils"

	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var TestDB = &database.Database{}

func SetupTestDB(m *testing.M) {
	fmt.Println("SetupTestDB")
	if TestDB != nil && TestDB.SQLDB != nil {
		fmt.Println("Test DB is already initialized")
		return
	}
	fmt.Println("Start running local database integration testing...")
	var err error
	TestDB.SQLDB, err = sqlx.Connect("postgres", "host=localhost port=5436 user=testadmin password=qwerty123 dbname=local_test_database sslmode=disable")
	if err != nil {
		fmt.Printf("failed to connect to test DB: %v", err)
		os.Exit(0)
	} else {
		fmt.Println("Connected to test local database")
	}
	database.RunMigrations(TestDB.SQLDB, utils.GetMigrationsDir())
	code := m.Run()
	fmt.Printf("Test run finished with result code %d", code)

	TestDB.SQLDB.Close()
	os.Exit(code)
}

func PrepareTestUsersData(t *testing.T) {
	testClientEntity1 := dbUser.UserEntity{
		FirstName: "John",
		LastName:  "Doe",
		Role:      middleware.Client.VALUE,
		Providers: &[]dbUser.UserProviderEntity{
			{
				ProviderName: "google",
			},
		},
	}

	testClientEntity2 := dbUser.UserEntity{
		FirstName: "Jane",
		LastName:  "Doe",
		Role:      middleware.Client.VALUE,
		Providers: &[]dbUser.UserProviderEntity{
			{
				ProviderName:   "email",
				HashedPassword: "hashed_password",
				Email:          "example@email.com",
			},
		},
	}
	_, errCreateU1 := dbUser.CreateUser(TestDB, &testClientEntity1)
	assert.NoError(t, errCreateU1)
	_, errCreateU2 := dbUser.CreateUser(TestDB, &testClientEntity2)
	assert.NoError(t, errCreateU2)
}

func PrepareTestLanguageData(t *testing.T) {
	testLanguageEntity1 := dbLanguage.LanguageEntity{
		UserId: 1,
		Name:   "English",
		Code:   "en-US",
	}

	testLanguageEntity2 := dbLanguage.LanguageEntity{
		UserId: 1,
		Name:   "Ukrainian",
		Code:   "uk-UA",
	}
	errCreateL1 := dbLanguage.CreateLanguage(TestDB, 1, &testLanguageEntity1)
	assert.NoError(t, errCreateL1)

	errCreateL2 := dbLanguage.CreateLanguage(TestDB, 1, &testLanguageEntity2)
	assert.NoError(t, errCreateL2)
}
