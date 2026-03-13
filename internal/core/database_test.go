package core_test

import (
	"docpile/internal/core"
	"testing"
)

func TestDatabaseConfig(t *testing.T) {
	tests := []struct {
		name string
		cfg  core.DatabaseConfig
		err  error
	}{
		{name: "valid sqlite", cfg: core.DatabaseConfig{Driver: "sqlite", DSN: ":memory:"}},
		{name: "missing dsn", cfg: core.DatabaseConfig{Driver: "sqlite"}},
		{name: "invalid dsn", cfg: core.DatabaseConfig{Driver: "sqlite3"}},
		{name: "missing driver", cfg: core.DatabaseConfig{DSN: ":memory:"}},
	}

	for _, test := range tests {
		err := test.cfg.Validate()
		if err != test.err {
			t.Errorf("expected %v, got %v", test.err, err)
		}
	}
}

const (
	sqliteQuery = "SELECT 1 FROM sqlite_master WHERE type = 'table' AND name = 'test_table';"
)

func TestDatabase(t *testing.T) {
	tests := []struct {
		name string
		cfg  core.DatabaseConfig
	}{
		{name: "sqlite", cfg: core.DatabaseConfig{Driver: "sqlite", DSN: ":memory:"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, err := core.NewDatabase(test.cfg)
			if err != nil {
				t.Fatalf("expected database, got %v", err)
			}
			defer db.Close()

			err = db.Migrate()
			if err != nil {
				t.Errorf("expected migration succes, got %v", err)
			}

			var query string
			switch test.cfg.Driver {
			case sqliteQuery:
				query = sqliteQuery
			}

			row := db.QueryRowContext(t.Context(), query)
			if row.Err() != nil {
				t.Fatalf("failed to query db: %v", err)
			}

			var got int
			err = row.Scan(&got)
			if err != nil {
				t.Errorf("failed to scan row: %v", err)
			}

			if got != 1 {
				t.Errorf("expected 1 from query, got %d", got)
			}
		})
	}
}
