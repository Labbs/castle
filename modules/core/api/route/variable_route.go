package route

import (
	"github.com/labbs/castle/modules/core/api/controller"
	"github.com/labbs/castle/modules/core/bootstrap"
	"github.com/labbs/castle/modules/core/repository"
)

func NewVariableRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing task routes")
	pr := repository.NewCoreRepository(app.Db)
	pc := &controller.VariableController{
		Repository: pr,
		Logger:     app.Logger,
	}

	gr := app.Fiber.Group("/api/variable")

	gr.Get("/project/:id", pc.GetAllVariablesByProjectId)
	gr.Get("/environment/:id", pc.GetAllVariablesByEnvironmentId)
	gr.Get("/task/:id", pc.GetAllVariablesByTaskId)
	gr.Post("/create", pc.CreateVariable)
	gr.Post("/update/:id", pc.EditVariable)
	gr.Delete("/delete/:id", pc.DeleteVariable)
	gr.Get("/:id", pc.GetVariableById)
}
