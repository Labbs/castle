package bus

import (
	"github.com/labbs/castle/modules/project/bootstrap"
	"github.com/labbs/castle/modules/project/repository"
)

type Controller struct {
	Repository repository.ProjectRepository
}

func Setup(app bootstrap.Application) {
	ur := repository.NewProjectRepository(app.Db)
	uc := &Controller{
		Repository: ur,
	}

	app.Bus.AddHandler("project:get_by_id", uc.GetProjectById)
	app.Bus.AddHandler("project:get_all", uc.GetAllProjects)
}
