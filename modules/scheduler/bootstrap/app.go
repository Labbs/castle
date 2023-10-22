package bootstrap

import (
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"

	initBootstrap "github.com/labbs/castle/bootstrap"
)

type Application struct {
	Logger         zerolog.Logger
	FiberApiRouter fiber.Router
	Fiber          *fiber.App
	Bus            *initBootstrap.Module
	CronScheduler  *gocron.Scheduler
}

func App(initBootstrapApp *initBootstrap.Application) Application {
	app := &Application{}
	app.Logger = InitLogger(initBootstrapApp.Logger)
	app.Fiber = initBootstrapApp.Fiber
	app.FiberApiRouter = initBootstrapApp.FiberApiRouter
	app.Bus = initBootstrapApp.Bus
	app.CronScheduler = InitCronScheduler()

	return *app
}
