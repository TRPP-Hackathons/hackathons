package database

import (
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"

	"hackathons/internal/infrastructure/database/postgres"
)

var (
	//go:embed migrations/hackathons/*.sql
	migrations embed.FS
)

func MigrateHackathonsDB(db *postgres.Postgres) error {
	if err := migrate(db, "migrations/hackathons"); err != nil {
		return fmt.Errorf("migrate: %v", err)
	}

	return nil
}

func migrate(db *postgres.Postgres, dir string) error {
	goose.SetBaseFS(migrations)

	if err := goose.Up(db.SqlDB(), dir); err != nil {
		return fmt.Errorf("goose up: %v", err)
	}

	return nil
}
