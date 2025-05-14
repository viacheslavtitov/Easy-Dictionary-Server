package db

import (
	internalenv "easy-dictionary-server/internalenv"
	utils "easy-dictionary-server/internalenv/utils"

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
	defer db.Close()
	database := Database{SQLDB: db}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		zap.S().Error(err)
		zap.S().Fatal("Couldn't connect to database")
	} else {
		zap.S().Info("Successfully Connected")
	}

	//create database if not exists
	sqlCreateDB := PrepareRunSetupBaseDBQuery(env)
	zap.S().Debug(sqlCreateDB)
	result, err := db.Exec(sqlCreateDB)
	if err != nil {
		zap.S().Error(err)
		zap.S().Fatal("Couldn't create database")
	}
	lastInsertId, _ := result.LastInsertId()
	rowsAffected, _ := result.RowsAffected()
	zap.S().Debugf("%d %d", lastInsertId, rowsAffected)

	//create tables and relations if not exists
	initDataQuery, err := utils.ReadFile("../Easy-Dictionary.sql")
	if err != nil {
		zap.S().Fatal("Failed to read init sql file")
	} else {
		_, err := db.Exec(initDataQuery)
		if err != nil {
			zap.S().Error(err)
			zap.S().Fatal("Something wrong of run sql query")
		}
	}
	return &database
}

func PrepareConnectionQuery(env *internalenv.Env) string {
	return "user=" + env.DBUser + " sslmode=disable password=" + env.DBPassword + " host=" + env.DBHost + " dbname=" + env.DBName + " port=" + strconv.Itoa(env.DBPort)
}

func PrepareRunSetupBaseDBQuery(env *internalenv.Env) string {
	return "SELECT 'CREATE DATABASE " + env.DBName + " OWNER = " + env.DBUser + " ENCODING = UTF8' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '" + env.DBName + "')"
}
