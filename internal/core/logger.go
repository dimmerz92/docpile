package core

import (
	"fmt"
	"log/slog"
	"maps"
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

var (
	ErrLoggerFormat = fmt.Errorf("logger format not supported, use any of %v", maps.Keys(supportedLoggerFormats))
	ErrLoggerLevel  = fmt.Errorf("logger level not supported, use any of %v", maps.Keys(supportedLoggerLevels))
)

type LoggerConfig struct {
	Format string `toml:"format"`
	Level  string `toml:"level"`
}

func (lc *LoggerConfig) Validate() error

func (lc *LoggerConfig) Init() error
