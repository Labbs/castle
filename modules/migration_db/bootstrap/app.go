package bootstrap

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	initBootstrap "github.com/labbs/castle/bootstrap"
)

type Application struct {
	Db     *gorm.DB
	Logger zerolog.Logger
}

func App(initBootstrapApp *initBootstrap.Application) {
	app := &Application{}
	app.Db = initBootstrapApp.Db
	app.Logger = InitLogger(initBootstrapApp.Logger)
	MigrateDatabase(*app, *initBootstrapApp.AppConfig)
}
