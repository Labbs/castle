package route

import (
	"github.com/labbs/castle/modules/core/api/controller"
	"github.com/labbs/castle/modules/core/bootstrap"
	"github.com/labbs/castle/modules/core/repository"
)

func NewRepositoryRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing repository routes")
	pr := repository.NewCoreRepository(app.Db)
	pc := &controller.RepositoryController{
		Repository: pr,
		Logger:     app.Logger,
	}

	gr := app.Fiber.Group("/api/repository")

	gr.Get("/all", pc.GetAllRepositories)
	gr.Get("/:id", pc.GetRepositoryById)
	gr.Post("/create", pc.CreateRepository)
	gr.Post("/update/:id", pc.EditRepository)
	gr.Delete("/delete/:id", pc.DeleteRepository)
}
