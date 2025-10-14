package db

import (
	"context"
	internalenv "easy-dictionary-server/internalenv"
	"fmt"
	"time"

	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Database struct {
	SQLDB *sqlx.DB
}

func (db *Database) GetConnection() (*sqlx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	conn, err := db.SQLDB.Connx(ctx)
	if err != nil {
		return nil, fmt.Errorf("get conn: %w", err)
	}
	defer conn.Close()
	return conn, nil
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
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)
	lis := pq.NewListener(sqlConnectQuery, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		switch ev {
		case pq.ListenerEventConnected:
			zap.S().Info("DB connected")
		case pq.ListenerEventReconnected:
			zap.S().Warn("DB reconnected")
		case pq.ListenerEventDisconnected:
			zap.S().Error("DB disconnected: ", err)
		case pq.ListenerEventConnectionAttemptFailed:
			zap.S().Error("DB connect attempt failed: ", err)
		}
	})

	if err := lis.Listen("health_channel"); err != nil {
		zap.S().Error("DB listen failed: ", err)
	}
	return &database
}

func PrepareConnectionQuery(env *internalenv.Env) string {
	return "user=" + env.DBUser + " sslmode=disable password=" + env.DBPassword + " host=" + env.DBHost + " dbname=" + env.DBName + " port=" + strconv.Itoa(env.DBPort)
}
