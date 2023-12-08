package bus

import (
	"github.com/labbs/castle/modules/repository/bootstrap"
	"github.com/labbs/castle/modules/repository/domain"
	"github.com/labbs/castle/modules/repository/repository"
	"github.com/labbs/castle/modules/repository/service"
)

type RepositoryController struct {
	Service domain.RepositoryService
}

func Setup(app bootstrap.Application) {
	ur := repository.NewRepostioryRepository(app.Db)
	uc := &RepositoryController{
		Service: service.NewRepositoryService(ur),
	}

	app.Bus.AddHandler("repository:get_by_id", uc.GetRepositoryById)
	app.Bus.AddHandler("repository:get_all", uc.GetAllRepositories)
	app.Bus.AddHandler("repository:create", uc.CreateRepository)
	app.Bus.AddHandler("repository:update", uc.UpdateRepository)
	app.Bus.AddHandler("repository:delete", uc.DeleteRepository)
	app.Bus.AddHandler("repository:clone", uc.CloneRepository)
	app.Bus.AddHandler("repository:clone_test", uc.CloneTestRepository)
}
