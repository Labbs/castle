package route

import (
	"github.com/labbs/castle/modules/frontend/app/controller"
	"github.com/labbs/castle/modules/frontend/bootstrap"
	"github.com/labbs/castle/modules/frontend/repository"
)

func NewRepositoryRoute(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing frontend repository routes")
	lc := &controller.RepositoryController{
		Repository: repository.NewRepositoryRepository(app.BusMessages),
		Logger:     app.Logger,
	}

	gr := app.FiberRouter.Group("/repository")
	gr.Get("/", lc.ViewAll)
	gr.Get("/create", lc.Create)
	gr.Post("/create", lc.Create)
}
