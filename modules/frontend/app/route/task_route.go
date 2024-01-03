package route

import (
	"github.com/labbs/castle/modules/frontend/app/controller"
	"github.com/labbs/castle/modules/frontend/bootstrap"
	"github.com/labbs/castle/modules/frontend/repository"
)

func NewTaskRoute(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing frontend task routes")
	lc := &controller.TaskController{
		ProjectRepository:     repository.NewProjectRepository(app.BusMessages, app.Logger),
		TaskRepository:        repository.NewTaskRepository(app.BusMessages),
		RespositoryRepository: repository.NewRepositoryRepository(app.BusMessages),
		Logger:                app.Logger,
	}

	gr := app.FiberRouter.Group("/task")
	gr.Get("/create/project/:id", lc.Create)
	gr.Post("/create/project/:id", lc.Create)
}
