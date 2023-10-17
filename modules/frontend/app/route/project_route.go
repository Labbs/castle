package route

import (
	"github.com/labbs/castle/modules/frontend/app/controller"
	"github.com/labbs/castle/modules/frontend/bootstrap"
	"github.com/labbs/castle/modules/frontend/repository"
)

func NewProjectRoute(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing frontend project routes")
	lc := &controller.ProjectController{
		ProjectRepository: repository.NewProjectRepository(app.BusMessages),
		TaskRepository:    repository.NewTaskRepository(app.BusMessages),
		Logger:            app.Logger,
	}

	gr := app.FiberRouter.Group("/project")
	gr.Get("/", lc.ViewAll)
	gr.Get("/create", lc.Create)
	gr.Post("/create", lc.Create)
	gr.Get("/:id/view", lc.GetProjectById)
	gr.Get("/:id/edit", nil)
	gr.Post("/:id/edit", nil)
}
