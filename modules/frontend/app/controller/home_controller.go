package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/frontend/repository"
	"github.com/rs/zerolog"
)

type HomeController struct {
	ProjectRepository repository.ProjectRepository
	TaskRepository    repository.TaskRepository
	Logger            zerolog.Logger
}

func (hc *HomeController) Home(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "Home"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	projects, err := hc.ProjectRepository.GetAllProjects()
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.project.view_all").Msg("failed to get projects")
		d["Error"] = "Internal server error - failed to get projects"
	}
	d["Projects"] = projects

	return c.Render("templates/home", d, "templates/layouts/main")
}
