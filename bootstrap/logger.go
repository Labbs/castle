// Path: bootstrap/logger.go
// This is a generated file. Do not modify.
package bootstrap

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/labbs/castle/config"
)

func InitLogger() zerolog.Logger {
	host, _ := os.Hostname()
	logger := zerolog.New(os.Stderr).With().
		Caller().
		Timestamp().
		Str("host", host).
		Str("version", config.Version).
		Logger()

	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if config.PrettyLogs {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return logger
}
