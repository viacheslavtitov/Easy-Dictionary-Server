package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file:// source
	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB, migrationPath string) {
	fmt.Println("Start migration...")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		// zap.S().Fatalf("could not create migration driver: %v", err)
		fmt.Printf("could not create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres", driver,
	)
	if err != nil {
		// zap.S().Fatalf("could not init migrate: %v", err)
		fmt.Printf("could not init migrate: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		// zap.S().Fatalf("migration failed: %v", err)
		fmt.Printf("migration failed: %v", err)
	}
	// zap.S().Info("Migrations applied successfully")
	fmt.Println("Migrations applied successfully")
}
