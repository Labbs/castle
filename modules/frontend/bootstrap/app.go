package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	initBootstrap "github.com/labbs/castle/bootstrap"
)

type Application struct {
	Logger      zerolog.Logger
	Fiber       *fiber.App
	BusMessages chan initBootstrap.Message
}

func App(initBootstrapApp *initBootstrap.Application) Application {
	app := &Application{}
	app.Logger = InitLogger(initBootstrapApp.Logger)
	app.Fiber = initBootstrapApp.Fiber
	app.BusMessages = initBootstrapApp.BusMessages

	return *app
}
