package logger

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/5verst/config"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zlog *zerolog.Logger

// nolint: gofumpt
func MustLogger(cfg config.Config) *zerolog.Logger {
	var writers []io.Writer

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if cfg.LogFile == "" {
		console := zerolog.ConsoleWriter{Out: os.Stdout}
		console.FormatLevel = func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("%-6s|", i)) }
		console.FormatMessage = func(i interface{}) string { return strings.ToUpper(fmt.Sprintf("%-6s|", i)) }
		writers = append(writers, console)
	} else {
		logFile, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			panic(fmt.Errorf("open log file: %v", err))
		}

		file := &lumberjack.Logger{
			Filename: logFile.Name(),
			MaxSize:  10, // nolint:gomnd
		}

		writers = append(writers, file)
	}

	zerolog.MessageFieldName = "MESSAGE"
	zerolog.LevelFieldName = "LEVEL"

	mw := io.MultiWriter(writers...)
	logger := zerolog.New(mw).With().Logger()
	logger.Info().
		Str("PROGRAM", "go-launcher-api").
		Str("filename", cfg.LogFile).
		Int("maxSizeMB", 10).
		Msg("logging configured")

	zlog = &logger

	return zlog
}
