package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/frontend/domain"
	"github.com/labbs/castle/modules/frontend/repository"
	"github.com/rs/zerolog"
)

type RepositoryController struct {
	Repository repository.RepositoryRepository
	Logger     zerolog.Logger
}

func (hc *RepositoryController) ViewAll(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "Projects list"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	repositories, err := hc.Repository.GetAllRepositories()
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.repository.view_all").Msg("failed to get repositories")
		d["Error"] = "Internal server error - failed to get repositories"
	}

	d["Repositories"] = repositories

	return c.Render("templates/repository_view_list", d, "templates/layouts/main")
}

func (hc *RepositoryController) Create(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "New repository"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	if c.Method() == "POST" {
		var repository domain.BusRepositoryResponse
		repository.Name = c.FormValue("name")
		repository.Description = c.FormValue("description")
		repository.Url = c.FormValue("url")
		repository.SSHKey = c.FormValue("ssh_key")
		repository.SSHKeyPass = c.FormValue("ssh_key_pass")

		err := hc.Repository.CreateRepository(repository)
		if err != nil {
			hc.Logger.Error().Err(err).Str("event", "frontend.repository.create").Msg("failed to create repository")
			d["Error"] = "Internal server error - failed to create repository"
			return c.Render("templates/repository_new_update", d, "templates/layouts/main")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "success-flash",
			Value: "Repository created successfully",
		})

		return c.Redirect("/app/repository")
	}

	return c.Render("templates/repository_new_update", d, "templates/layouts/main")
}
