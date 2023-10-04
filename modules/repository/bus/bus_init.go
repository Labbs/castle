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
}
