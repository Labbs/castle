// Path: bootstrap/logger.go
// This is a generated file. Do not modify.
package bootstrap

import (
	"os"

	"github.com/rs/zerolog"

	"github.com/labbs/castle/config"
)

func InitLogger(c config.Config) zerolog.Logger {
	host, _ := os.Hostname()
	logger := zerolog.New(os.Stderr).With().
		Caller().
		Timestamp().
		Str("host", host).
		Str("version", c.Version).
		Logger()

	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if c.PrettyLogs {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return logger
}
