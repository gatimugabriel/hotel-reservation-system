package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func RunMigrations(db *sql.DB) error {
	// Create an instance of the postgres driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres migration driver: %w", err)
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"hestia", driver)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	// Run migrations up to the latest version
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not run up migrations: %w", err)
	}

	log.Printf("Migrations ran successfully")
	return nil
}