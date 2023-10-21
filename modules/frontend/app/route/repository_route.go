package route

import (
	"github.com/labbs/castle/modules/frontend/app/controller"
	"github.com/labbs/castle/modules/frontend/bootstrap"
	"github.com/labbs/castle/modules/frontend/repository"
)

func NewRepositoryRoute(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing frontend repository routes")
	lc := &controller.RepositoryController{
		RepositoryRepository: repository.NewRepositoryRepository(app.BusMessages),
		TaskRepository:       repository.NewTaskRepository(app.BusMessages),
		Logger:               app.Logger,
	}

	gr := app.FiberRouter.Group("/repository")
	gr.Get("/", lc.ViewAll)
	gr.Get("/create", lc.Create)
	gr.Post("/create", lc.Create)
	gr.Get("/:id/update", lc.UpdateRepositoryById)
	gr.Post("/:id/update", lc.UpdateRepositoryById)
	// gr.Get("/:id/delete", lc.DeleteRepositoryById)
	gr.Get("/:id/test_clone", lc.TestCloneRepositoryById)
}
