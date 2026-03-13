package core

import "database/sql"

type DatabaseConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
}

func (dc *DatabaseConfig) Validate() error

type Database struct {
	*sql.DB
}

func NewDatabase(cfg DatabaseConfig) (*Database, error)

func (db *Database) Migrate() error

/***********************************/
/* DRIVER SPECIFIC IMPLEMENTATIONS */
/***********************************/

func newSqlite(cfg DatabaseConfig) (*sql.DB, error)
