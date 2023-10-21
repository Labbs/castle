package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Application struct {
	Db             *gorm.DB
	FiberApiRouter fiber.Router
	FiberAppRouter fiber.Router
	Fiber          *fiber.App
	Logger         zerolog.Logger
	Bus            *Module
	BusMessages    chan Message
	AppConfig      *config.Config
}

func App(c *config.Config) *Application {
	app := &Application{}
	app.AppConfig = c
	app.Logger = InitLogger(*app.AppConfig)
	app.Db = InitDatabase(app.Logger, *app.AppConfig)
	app.Fiber = InitFiber(app.Logger, *app.AppConfig)
	app.FiberApiRouter = app.Fiber.Group("/api")
	app.FiberAppRouter = app.Fiber.Group("/app")
	app.Bus, app.BusMessages = InitBusComm()

	return app
}
