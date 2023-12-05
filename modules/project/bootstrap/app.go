package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	initBootstrap "github.com/labbs/castle/bootstrap"
)

type Application struct {
	Db             *gorm.DB
	Logger         zerolog.Logger
	FiberApiRouter fiber.Router
	Fiber          *fiber.App
	Bus            *initBootstrap.Module
}

func App(initBootstrapApp *initBootstrap.Application) Application {
	app := &Application{}
	app.Db = initBootstrapApp.Db
	app.Logger = InitLogger(initBootstrapApp.Logger)
	app.Fiber = initBootstrapApp.Fiber
	app.FiberApiRouter = initBootstrapApp.FiberApiRouter
	app.Bus = initBootstrapApp.Bus

	return *app
}
