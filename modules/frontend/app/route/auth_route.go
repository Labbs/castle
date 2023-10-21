package route

import (
	"github.com/labbs/castle/modules/frontend/app/controller"
	"github.com/labbs/castle/modules/frontend/bootstrap"
	"github.com/labbs/castle/modules/frontend/repository"
)

func NewAuthRoute(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing frontend auth routes")
	ur := repository.NewUserRepository(app.BusMessages)
	lc := &controller.AuthController{
		Repository: ur,
		Logger:     app.Logger,
	}

	gr := app.FiberRouter.Group("/auth")
	gr.Get("/login", lc.Login)
	gr.Post("/login", lc.Login)
}
