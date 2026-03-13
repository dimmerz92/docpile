package core_test

import (
	"bytes"
	"docpile/internal/core"
	"log/slog"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestLoggerConfig(t *testing.T) {
	msg := "hello world"

	var loggerLevels = map[string][]string{
		"debug": {"debug", "info", "warn", "error"},
		"info":  {"info", "warn", "error"},
		"warn":  {"warn", "error"},
		"error": {"error"},
	}

	catchStdOut := func(f func()) string {
		t.Helper()

		stdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		f()

		_ = w.Close()
		os.Stdout = stdout

		var buf bytes.Buffer
		if _, err := buf.ReadFrom(r); err != nil {
			t.Fatalf("failed to read from stdout: %v", err)
		}

		return buf.String()
	}

	tests := []struct {
		name string
		cfg  core.LoggerConfig
	}{
		{name: "default logger (json/warn)"},
		{name: "json/debug", cfg: core.LoggerConfig{Level: "debug"}},
		{name: "json/info", cfg: core.LoggerConfig{Level: "info"}},
		{name: "json/warn", cfg: core.LoggerConfig{Level: "warn"}},
		{name: "json/error", cfg: core.LoggerConfig{Level: "error"}},
		{name: "text defaul level (text/warn)", cfg: core.LoggerConfig{Format: "text"}},
		{name: "text/debug", cfg: core.LoggerConfig{Format: "text", Level: "debug"}},
		{name: "text/info", cfg: core.LoggerConfig{Format: "text", Level: "info"}},
		{name: "text/warn", cfg: core.LoggerConfig{Format: "text", Level: "warn"}},
		{name: "text/error", cfg: core.LoggerConfig{Format: "text", Level: "error"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, level := range []string{"debug", "info", "warn", "error"} {
				got := catchStdOut(func() {
					err := test.cfg.Init()
					if err != nil {
						t.Fatalf("failed to init with error: %v", err)
					}

					switch level {
					case "debug":
						slog.Debug(msg)
					case "info":
						slog.Info(msg)
					case "warn":
						slog.Warn(msg)
					case "error":
						slog.Error(msg)
					}
				})

				_ = test.cfg.Validate()
				if slices.Contains(loggerLevels[test.cfg.Level], level) {
					if !strings.Contains(got, msg) {
						t.Errorf("expected %s, got %s", msg, got)
					}
				} else if got != "" {
					println(got)
					t.Errorf("expected no output for level %s with logger level %s", test.cfg.Level, level)
				}
			}
		})
	}
}
