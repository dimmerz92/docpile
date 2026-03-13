package core

import (
	"database/sql"
	"docpile/migrations"
	"errors"
	"fmt"
	"maps"
	"path"
	"slices"
	"strings"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

var supportedDatabaseDrivers = map[string]string{
	"sqlite": "sqlite",
}

var (
	ErrDatabaseDriver = fmt.Errorf(
		"unsupported database driver, use any of %v",
		slices.Collect(maps.Keys(supportedDatabaseDrivers)),
	)
	ErrDatabaseDSN  = fmt.Errorf("database dsn (connection string) cannot be empty")
	ErrDatabaseOpen = fmt.Errorf("failed to open database")
)

type DatabaseConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
	Test   bool
}

func (dc *DatabaseConfig) Validate() error {
	dc.Driver = strings.TrimSpace(dc.Driver)
	if _, ok := supportedDatabaseDrivers[dc.Driver]; !ok {
		return ErrDatabaseDriver
	}

	dc.DSN = strings.TrimSpace(dc.DSN)
	if dc.DSN == "" {
		return ErrDatabaseDSN
	}

	return nil
}

type Database struct {
	*sql.DB
	driver string
	test   bool
}

func NewDatabase(cfg DatabaseConfig) (*Database, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(supportedDatabaseDrivers[cfg.Driver], cfg.DSN)
	if err != nil {
		return nil, errors.Join(ErrDatabaseOpen, err)
	}

	return &Database{
		DB:     db,
		driver: supportedDatabaseDrivers[cfg.Driver],
		test:   cfg.Test,
	}, nil
}

func (db *Database) Migrate() error {
	goose.SetBaseFS(migrations.FS)

	err := goose.SetDialect(db.driver)
	if err != nil {
		return err
	}

	migrationDir := IIF(db.test, path.Join(db.driver, "test"), db.driver)

	return goose.Up(db.DB, migrationDir)
}
