package route

import (
	"github.com/labbs/castle/modules/repository/api/controller"
	"github.com/labbs/castle/modules/repository/bootstrap"
	"github.com/labbs/castle/modules/repository/repository"
	"github.com/labbs/castle/modules/repository/service"
)

func NewRepositoryRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing repository routes")
	pr := repository.NewRepostioryRepository(app.Db)
	pc := &controller.RepositoryController{
		Service: service.NewRepositoryService(pr),
		Logger:  app.Logger,
	}

	gr := app.FiberApiRouter.Group("/repository")

	gr.Get("/all", pc.GetAllRepositories)
	gr.Get("/:id", pc.GetRepositoryById)
	gr.Post("/create", pc.CreateRepository)
	gr.Post("/update/:id", pc.EditRepository)
	// gr.Delete("/delete/:id", pc.DeleteRepository)
}
