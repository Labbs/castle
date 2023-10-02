package bootstrap

import (
	"github.com/labbs/castle/modules/core/domain"
)

func InitOrMigrateDatabase(app Application) error {
	db := app.Db
	logger := app.Logger
	err := db.AutoMigrate(
		&domain.Environment{},
		&domain.Project{},
		&domain.Task{},
		&domain.Variable{},
		&domain.TaskHistory{},
		&domain.Repository{},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error on migrate database")
	}

	return nil
}
