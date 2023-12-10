package route

import (
	"github.com/labbs/castle/modules/project/api/controller"
	"github.com/labbs/castle/modules/project/bootstrap"
	"github.com/labbs/castle/modules/project/repository"
	"github.com/labbs/castle/modules/project/service"
)

func NewEnvironmentRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing environment routes")
	pr := repository.NewEnvironmentRepository(app.Db)
	pc := &controller.EnvironmentController{
		Service: service.NewEnvironmentService(pr),
		Logger:  app.Logger,
	}

	gr := app.FiberApiRouter.Group("/environment")

	gr.Get("/all", pc.GetAllEnvironments)
	gr.Post("/create", pc.CreateEnvironment)
}
