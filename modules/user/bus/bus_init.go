package bus

import (
	"github.com/labbs/castle/modules/user/bootstrap"
	"github.com/labbs/castle/modules/user/repository"
)

type UserController struct {
	Repository repository.UserRepository
}

func Setup(app bootstrap.Application) {
	ur := repository.NewUserRepository(app.Db)
	uc := &UserController{
		Repository: ur,
	}

	app.Bus.AddHandler("user:get_by_email", uc.GetByEmail)
}
