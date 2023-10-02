package route

import "github.com/labbs/castle/modules/core/bootstrap"

func Setup(app bootstrap.Application) {
	NewProjectRouter(app)
	NewRepositoryRouter(app)
	NewTaskRouter(app)
	NewVariableRouter(app)
	NewEnvironmentRouter(app)
}
