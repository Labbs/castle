package bootstrap

import (
	"github.com/labbs/castle/config"
	pb "github.com/labbs/castle/gen/repository"
)

func InitOrMigrateDatabase(app Application, c config.Config) error {
	db := app.Db
	logger := app.Logger
	err := db.AutoMigrate(
		&pb.Repository{},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error on migrate database")
	}

	return nil
}
