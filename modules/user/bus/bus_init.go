package bus

import (
	"github.com/labbs/castle/modules/user/bootstrap"
	"github.com/labbs/castle/modules/user/repository"
	"github.com/rs/zerolog"
)

type UserController struct {
	Repository repository.UserRepository
	Logger     zerolog.Logger
}

func Setup(app bootstrap.Application) {
	ur := repository.NewUserRepository(app.Db)
	uc := &UserController{
		Repository: ur,
		Logger:     app.Logger,
	}

	app.Bus.AddHandler("user:get_by_email", uc.GetByEmail)
}
