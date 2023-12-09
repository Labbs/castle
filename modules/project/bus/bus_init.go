package bus

import (
	"github.com/labbs/castle/modules/project/bootstrap"
	"github.com/labbs/castle/modules/project/domain"
	"github.com/labbs/castle/modules/project/repository"
)

type Controller struct {
	ProjectService     domain.ProjectService
	EnvironmentService domain.EnvironmentService
}

func Setup(app bootstrap.Application) {
	uc := &Controller{
		ProjectService:     repository.NewProjectRepository(app.Db),
		EnvironmentService: repository.NewEnvironmentRepository(app.Db),
	}

	app.Bus.AddHandler("project:get_by_id", uc.GetProjectById)
	app.Bus.AddHandler("project:get_all", uc.GetAllProjects)
	app.Bus.AddHandler("project:create", uc.CreateProject)
	app.Bus.AddHandler("project:update", uc.UpdateProject)
	app.Bus.AddHandler("project:delete", uc.DeleteProject)

	app.Bus.AddHandler("environment:get_by_id", uc.GetEnvironmentById)
	app.Bus.AddHandler("environment:get_all", uc.GetAllEnvironments)
	app.Bus.AddHandler("environment:create", uc.CreateEnvironment)
	app.Bus.AddHandler("environment:update", uc.UpdateEnvironment)
	app.Bus.AddHandler("environment:delete", uc.DeleteEnvironment)
}
