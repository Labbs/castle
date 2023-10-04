package route

import "github.com/labbs/castle/modules/task/bootstrap"

func Setup(app bootstrap.Application) {
	NewTaskRouter(app)
}
