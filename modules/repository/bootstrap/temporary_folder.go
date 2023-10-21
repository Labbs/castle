package bootstrap

import (
	"os"
	"strings"

	"github.com/labbs/castle/config"
	"github.com/rs/zerolog"
)

func InitTemporaryFolder(logger zerolog.Logger) {
	path := config.AppConfig.TemporaryFolder
	if strings.HasSuffix(path, "/") {
		path = path + "repos"
	} else {
		path = path + "/repos"
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	} else {
		logger.Info().Str("event", "bootstrap.init_temporary_folder").Msg("temporary folder already exists")
	}
}
