package route

import "github.com/labbs/castle/modules/repository/bootstrap"

func Setup(app bootstrap.Application) {
	NewRepositoryRouter(app)
}
