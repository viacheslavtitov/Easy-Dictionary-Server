package db

import (
	"go.uber.org/zap"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file:// source
	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		zap.S().Fatalf("could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migrations",
		"postgres", driver,
	)
	if err != nil {
		zap.S().Fatalf("could not init migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		zap.S().Fatalf("migration failed: %v", err)
	}
	zap.S().Info("Migrations applied successfully")
}
