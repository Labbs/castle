package route

import (
	"github.com/labbs/castle/modules/task/api/controller"
	"github.com/labbs/castle/modules/task/bootstrap"
	"github.com/labbs/castle/modules/task/repository"
	"github.com/labbs/castle/modules/task/service"
)

func NewTaskRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing task routes")

	pr := repository.NewTaskRepository(app.Db)
	pc := &controller.TaskController{
		Service: service.NewTaskService(pr),
		Logger:  app.Logger,
		// Scheduler:  app.Scheduler,
	}

	gr := app.FiberApiRouter.Group("/task")

	gr.Get("/project/:id", pc.GetAllTasksByProjectId)
	gr.Post("/create", pc.CreateTask)
}
