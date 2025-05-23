package db_test

import (
	"os"
	"testing"

	db "easy-dictionary-server/db"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var testDB *sqlx.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sqlx.Connect("postgres", "host=localhost port=5433 user=testuser password=testpass dbname=testdb sslmode=disable")
	if err != nil {
		zap.S().Fatalf("failed to connect to test DB: %v", err)
		os.Exit(0)
		return
	}
	db.RunMigrations(testDB)
	code := m.Run()

	testDB.Close()
	os.Exit(code)
}
