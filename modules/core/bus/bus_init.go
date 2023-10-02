package bus

import (
	"github.com/labbs/castle/modules/core/bootstrap"
	"github.com/labbs/castle/modules/core/repository"
)

type CoreController struct {
	Repository repository.CoreRepository
}

func Setup(app bootstrap.Application) {
	// ur := repository.NewCoreRepository(app.Db)
	// uc := &CoreController{
	// 	Repository: ur,
	// }

	// app.Bus.AddHandler("user:get_by_username", uc.GetByUsername)
}
