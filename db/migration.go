package db

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // file:// source
	"github.com/jmoiron/sqlx"

	"os"
	"path/filepath"
	"strings"
)

func RunMigrations(db *sqlx.DB, migrationPath string) {
	if runtime.GOOS != "windows" {
		runMigrationsForUnix(db, migrationPath)
	} else {
		runMigrationsForWindows(db, migrationPath)
	}
}

func runMigrationsForWindows(db *sqlx.DB, migrationPath string) {
	fmt.Printf("Start migration for Windows... %s\n", migrationPath)

	if wd, err := os.Getwd(); err == nil {
		fmt.Println("WD:", wd)
	}
	if exe, err := os.Executable(); err == nil {
		fmt.Println("EXE:", exe)
		fmt.Println("EXE DIR:", filepath.Dir(exe))
	}
	printDir(".")
	printDir("migrations")

	srcURL := migrationPath
	if !strings.HasPrefix(strings.ToLower(migrationPath), "file://") {
		abs, err := filepath.Abs(migrationPath)
		if err != nil {
			fmt.Printf("resolve migrations path failed: %v\n", err)
			return
		}
		srcURL = "file://" + filepath.ToSlash(abs)
	}
	pathPart := strings.TrimPrefix(srcURL, "file://")
	if st, err := os.Stat(pathPart); err != nil || !st.IsDir() {
		fmt.Printf("migrations path not found or not a dir: %s (err=%v)\n", pathPart, err)
		return
	}
	fmt.Println("Migrations source:", srcURL)

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		fmt.Printf("could not create migration driver: %v\n", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(srcURL, "postgres", driver)
	if err != nil {
		fmt.Printf("could not init migrate: %v\n", err)
		return
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			fmt.Printf("close source error: %v\n", srcErr)
		}
		if dbErr != nil {
			fmt.Printf("close db error: %v\n", dbErr)
		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("migration failed: %v\n", err)
		return
	}
	fmt.Println("Migrations applied successfully")
}

func runMigrationsForUnix(db *sqlx.DB, migrationPath string) {
	fmt.Printf("Start migration for Unix system... %s", migrationPath)
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

func printDir(rel string) {
	entries, err := os.ReadDir(rel)
	if err != nil {
		fmt.Printf("DIR %s: error: %v\n", rel, err)
		return
	}
	fmt.Printf("DIR %s:\n", rel)
	for _, e := range entries {
		info, _ := e.Info()
		size := int64(0)
		if info != nil {
			size = info.Size()
		}
		fmt.Printf("  %s  %8d  %s\n",
			map[bool]string{true: "<DIR>", false: "     "}[e.IsDir()],
			size,
			e.Name(),
		)
	}
}
