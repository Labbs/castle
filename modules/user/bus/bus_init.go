package bus

import (
	"github.com/labbs/castle/modules/user/bootstrap"
	"github.com/labbs/castle/modules/user/domain"
	"github.com/labbs/castle/modules/user/repository"
	"github.com/labbs/castle/modules/user/service"
	"github.com/rs/zerolog"
)

type UserController struct {
	Service domain.UserService
	Logger  zerolog.Logger
}

func Setup(app bootstrap.Application) {
	ur := repository.NewUserRepository(app.Db)
	uc := &UserController{
		Service: service.NewUserService(ur),
		Logger:  app.Logger,
	}

	app.Bus.AddHandler("user:get_by_email", uc.GetByEmail)
}
