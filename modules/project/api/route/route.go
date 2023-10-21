package route

import "github.com/labbs/castle/modules/project/bootstrap"

func Setup(app bootstrap.Application) {
	NewProjectRouter(app)
}
