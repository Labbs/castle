package route

import (
	"github.com/labbs/castle/modules/project/api/controller"
	"github.com/labbs/castle/modules/project/bootstrap"
	"github.com/labbs/castle/modules/project/repository"
)

func NewProjectRouter(app bootstrap.Application) {
	app.Logger.Info().Msg("Initializing project routes")
	pr := repository.NewProjectRepository(app.Db)
	pc := &controller.ProjectController{
		Repository: pr,
		Logger:     app.Logger,
	}

	gr := app.Fiber.Group("/api/project")

	gr.Get("/all", pc.GetAllProjects)
	gr.Get("/:id", pc.GetProjectById)
	gr.Post("/update/:id", pc.UpdateProject)
	// gr.Post("/create", pc.CreateProject)
}
