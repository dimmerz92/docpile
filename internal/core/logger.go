package core

import (
	"fmt"
	"log/slog"
	"maps"
	"os"
)

var supportedLoggerFormats = map[string]struct{}{
	"json": {},
	"text": {},
}

var supportedLoggerLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

const (
	LoggerDefaultFormat = "json"
	LoggerDefaultLevel  = "warn"
)

var (
	ErrLoggerFormat = fmt.Errorf("logger format not supported, use any of %v", maps.Keys(supportedLoggerFormats))
	ErrLoggerLevel  = fmt.Errorf("logger level not supported, use any of %v", maps.Keys(supportedLoggerLevels))
)

type LoggerConfig struct {
	Format string `toml:"format"`
	Level  string `toml:"level"`
}

func (lc *LoggerConfig) Validate() error {
	lc.Format = Coalesce(lc.Format, LoggerDefaultFormat)
	if _, ok := supportedLoggerFormats[lc.Format]; !ok {
		return ErrLoggerFormat
	}

	lc.Level = Coalesce(lc.Level, LoggerDefaultLevel)
	if _, ok := supportedLoggerLevels[lc.Level]; !ok {
		return ErrLoggerLevel
	}

	return nil
}

func (lc *LoggerConfig) Init() error {
	if err := lc.Validate(); err != nil {
		return err
	}

	opts := &slog.HandlerOptions{Level: supportedLoggerLevels[lc.Level]}

	var handler slog.Handler
	switch lc.Format {
	case "text":
		handler = slog.NewTextHandler(os.Stdout, opts)
	default:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))

	return nil
}
