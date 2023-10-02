package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Application struct {
	Db          *gorm.DB
	Fiber       *fiber.App
	Logger      zerolog.Logger
	Bus         *Module
	BusMessages chan Message
}

func App() *Application {
	app := &Application{}
	app.Logger = InitLogger()
	app.Db = InitDatabase(app.Logger)
	app.Fiber = InitFiber(app.Logger)
	app.Bus, app.BusMessages = InitBusComm()

	return app
}
