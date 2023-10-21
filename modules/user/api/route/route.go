package route

import "github.com/labbs/castle/modules/user/bootstrap"

func Setup(app bootstrap.Application) {
	NewUserRouter(app)
}
