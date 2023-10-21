package bootstrap

import (
	"github.com/labbs/castle/config"
	"github.com/labbs/castle/modules/repository/domain"
)

func InitOrMigrateDatabase(app Application, c config.Config) error {
	db := app.Db
	logger := app.Logger
	err := db.AutoMigrate(
		&domain.Repository{},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error on migrate database")
	}

	return nil
}
