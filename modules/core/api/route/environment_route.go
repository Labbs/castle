package route

import (
	"github.com/labbs/castle/modules/core/api/controller"
	"github.com/labbs/castle/modules/core/bootstrap"
	"github.com/labbs/castle/modules/core/repository"
)

func NewEnvironmentRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing environment routes")
	pr := repository.NewCoreRepository(app.Db)
	pc := &controller.EnvironmentController{
		Repository: pr,
		Logger:     app.Logger,
	}

	gr := app.Fiber.Group("/api/environment")

	gr.Get("/project/:id", pc.GetAllEnvironmentsByProjectId)
	gr.Post("/create", pc.CreateEnvironment)
	gr.Post("/update/:id", pc.EditEnvironment)
	gr.Delete("/delete/:id", pc.DeleteEnvironment)
}
