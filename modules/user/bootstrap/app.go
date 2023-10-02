package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	initBootstrap "github.com/labbs/castle/bootstrap"
)

type Application struct {
	Db     *gorm.DB
	Logger zerolog.Logger
	Fiber  *fiber.App
	Bus    *initBootstrap.Module
}

func App(initBootstrapApp *initBootstrap.Application) Application {
	app := &Application{}
	app.Db = initBootstrapApp.Db
	app.Logger = InitLogger(initBootstrapApp.Logger)
	InitOrMigrateDatabase(*app)
	app.Fiber = initBootstrapApp.Fiber
	app.Bus = initBootstrapApp.Bus

	return *app
}
