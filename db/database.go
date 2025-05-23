package db

import (
	internalenv "easy-dictionary-server/internalenv"

	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Database struct {
	SQLDB *sqlx.DB
}

func Setup(env *internalenv.Env) *Database {
	sqlConnectQuery := PrepareConnectionQuery(env)
	zap.S().Debug(sqlConnectQuery)
	db, err := sqlx.Open("postgres", sqlConnectQuery)
	if err != nil {
		zap.S().Error(err)
		zap.S().Fatal("Couldn't connect to database")
	}
	// defer db.Close()
	database := Database{SQLDB: db}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		zap.S().Error(err)
		zap.S().Fatal("Couldn't connect to database")
	} else {
		zap.S().Info("Successfully Connected")
	}
	return &database
}

func PrepareConnectionQuery(env *internalenv.Env) string {
	return "user=" + env.DBUser + " sslmode=disable password=" + env.DBPassword + " host=" + env.DBHost + " dbname=" + env.DBName + " port=" + strconv.Itoa(env.DBPort)
}
