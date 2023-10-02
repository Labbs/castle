package route

import (
	"github.com/labbs/castle/modules/auth/api/controller"
	"github.com/labbs/castle/modules/auth/bootstrap"
	"github.com/labbs/castle/modules/auth/repository"
)

func NewAuthRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing auth routes")
	ur := repository.NewAuthRepository(app.BusMessages)
	lc := &controller.AuthController{
		Repository: ur,
		Logger:     app.Logger,
	}

	app.Fiber.Post("/api/auth/login", lc.Login)
}
