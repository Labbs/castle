package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/core/internal/scheduler"
)

type Application struct {
	Db        *gorm.DB
	Logger    zerolog.Logger
	Fiber     *fiber.App
	Bus       *initBootstrap.Module
	Scheduler scheduler.SchedulerController
}

func App(initBootstrapApp *initBootstrap.Application) Application {
	app := &Application{}
	app.Db = initBootstrapApp.Db
	app.Logger = InitLogger(initBootstrapApp.Logger)
	InitOrMigrateDatabase(*app)
	app.Fiber = initBootstrapApp.Fiber
	app.Bus = initBootstrapApp.Bus
	app.Scheduler = InitScheduler(*app)

	app.Scheduler.Start()

	return *app
}
