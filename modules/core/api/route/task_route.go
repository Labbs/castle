package route

import (
	"github.com/labbs/castle/modules/core/api/controller"
	"github.com/labbs/castle/modules/core/bootstrap"
	"github.com/labbs/castle/modules/core/repository"
)

func NewTaskRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing task routes")

	pr := repository.NewCoreRepository(app.Db)
	pc := &controller.TaskController{
		Repository: pr,
		Logger:     app.Logger,
		Scheduler:  app.Scheduler,
	}

	gr := app.Fiber.Group("/api/task")

	gr.Get("/project/:id", pc.GetAllTasksByProjectId)
	gr.Post("/create", pc.CreateTask)
}
