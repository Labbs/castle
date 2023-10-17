package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	initBootstrap "github.com/labbs/castle/bootstrap"
)

type Application struct {
	Logger      zerolog.Logger
	FiberRouter fiber.Router
	Fiber       *fiber.App
	BusMessages chan initBootstrap.Message
}

func App(initBootstrapApp *initBootstrap.Application) Application {
	app := &Application{}
	app.Logger = InitLogger(initBootstrapApp.Logger)
	app.Fiber = initBootstrapApp.Fiber
	app.FiberRouter = initBootstrapApp.Fiber.Group("/app")
	app.BusMessages = initBootstrapApp.BusMessages

	InitSession(app)

	return *app
}
