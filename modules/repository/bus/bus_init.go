package bus

import (
	"github.com/labbs/castle/modules/repository/bootstrap"
	"github.com/labbs/castle/modules/repository/repository"
)

type RepositoryController struct {
	Repository repository.RepositoryRepository
}

func Setup(app bootstrap.Application) {
	ur := repository.NewRepostioryRepository(app.Db)
	uc := &RepositoryController{
		Repository: ur,
	}

	app.Bus.AddHandler("repository:get_by_id", uc.GetRepositoryById)
	app.Bus.AddHandler("repository:get_all", uc.GetAllRepositories)
	app.Bus.AddHandler("repository:create", uc.CreateRepository)
	app.Bus.AddHandler("repository:update", uc.UpdateRepository)
	app.Bus.AddHandler("repository:delete", uc.DeleteRepository)
	app.Bus.AddHandler("repository:clone", uc.CloneRepository)
	app.Bus.AddHandler("repository:clone_test", uc.CloneTestRepository)
}
