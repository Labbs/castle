package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/frontend/bootstrap"
)

func Setup(app bootstrap.Application) {
	app.Fiber.All("/", func(c *fiber.Ctx) error {
		return c.Redirect("/app/auth/login")
	})
	NewAuthRoute(app)
	NewHomeRoute(app)
	NewProjectRoute(app)
	NewRepositoryRoute(app)
	NewTaskRoute(app)
}
