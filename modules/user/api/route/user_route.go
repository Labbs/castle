package route

import (
	"github.com/labbs/castle/modules/user/api/controller"
	"github.com/labbs/castle/modules/user/bootstrap"
	"github.com/labbs/castle/modules/user/repository"
)

func NewUserRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing user routes")
	ur := repository.NewUserRepository(app.Db)
	pc := &controller.UserController{
		Repository: ur,
		Logger:     app.Logger,
	}

	gr := app.FiberApiRouter.Group("/user")

	gr.Get("/", pc.Get)
	gr.Put("/email", pc.EditEmail)
	gr.Put("/avatar", pc.EditAvatar)
	gr.Put("/password", pc.EditPassword)
	gr.Put("/dark_mode", pc.EditDarkMode)
}
