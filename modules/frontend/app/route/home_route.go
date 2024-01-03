package route

import (
	"github.com/labbs/castle/modules/frontend/app/controller"
	"github.com/labbs/castle/modules/frontend/bootstrap"
	"github.com/labbs/castle/modules/frontend/repository"
)

func NewHomeRoute(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing frontend home routes")
	lc := &controller.HomeController{
		ProjectRepository: repository.NewProjectRepository(app.BusMessages, app.Logger),
		TaskRepository:    repository.NewTaskRepository(app.BusMessages),
		Logger:            app.Logger,
	}

	app.FiberRouter.Get("/", lc.Home)
}
