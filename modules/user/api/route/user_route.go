package route

import (
	"github.com/labbs/castle/modules/user/api/controller"
	"github.com/labbs/castle/modules/user/bootstrap"
	"github.com/labbs/castle/modules/user/repository"
	"github.com/labbs/castle/modules/user/service"
)

func NewUserRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing user routes")
	ur := repository.NewUserRepository(app.Db)
	pc := &controller.UserController{
		Service: service.NewUserService(ur),
		Logger:  app.Logger,
	}

	gr := app.FiberApiRouter.Group("/user")

	gr.Get("/", pc.Get)
	// too dangerous to have this endpoint
	// gr.Put("/email", pc.EditEmail)
	gr.Put("/avatar", pc.EditAvatar)
	gr.Put("/password", pc.EditPassword)
	gr.Put("/dark_mode", pc.EditDarkMode)
}
