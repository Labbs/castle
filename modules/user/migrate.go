package bootstrap

import (
	"github.com/labbs/castle/config"
	pb "github.com/labbs/castle/gen/task"
)

func InitOrMigrateDatabase(app Application, c config.Config) error {
	db := app.Db
	logger := app.Logger
	err := db.AutoMigrate(
		&pb.Task{},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error on migrate database")
	}

	return nil
}
