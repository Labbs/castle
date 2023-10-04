package route

import (
	"github.com/labbs/castle/modules/task/api/controller"
	"github.com/labbs/castle/modules/task/bootstrap"
	"github.com/labbs/castle/modules/task/repository"
)

func NewTaskRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing task routes")

	pr := repository.NewTaskRepository(app.Db)
	pc := &controller.TaskController{
		Repository: pr,
		Logger:     app.Logger,
		// Scheduler:  app.Scheduler,
	}

	gr := app.Fiber.Group("/api/task")

	gr.Get("/project/:id", pc.GetAllTasksByProjectId)
	gr.Post("/create", pc.CreateTask)
}
