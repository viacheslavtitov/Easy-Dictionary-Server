package db

import (
	"fmt"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file:// source
	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB, migrationPath string) {
	fmt.Printf("Start migration... %s", migrationPath)
	cmd := exec.Command("pwd")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("PWD:", string(out))

	cmdLs := exec.Command("sh", "-c", "ls -la")
	outLs, err := cmdLs.Output()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("PWD:", string(outLs))

	cmdLsD := exec.Command("sh", "-c", "ls -la /easydictionary/migrations")
	outLsD, err := cmdLsD.Output()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("PWD:", string(outLsD))

	cmdLsApp := exec.Command("sh", "-c", "ls -la migrations")
	outLsApp, err := cmdLsApp.Output()
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("PWD:", string(outLsApp))

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
