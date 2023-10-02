package route

import (
	"github.com/labbs/castle/modules/auth/api/middleware"
	"github.com/labbs/castle/modules/auth/bootstrap"
)

func Setup(app bootstrap.Application) {
	app.Fiber.Use(middleware.JwtAuthMiddleware())

	NewAuthRouter(app)
}
