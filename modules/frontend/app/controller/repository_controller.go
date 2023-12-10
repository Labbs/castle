package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/frontend/domain"
	"github.com/labbs/castle/modules/frontend/repository"
	"github.com/rs/zerolog"
)

type RepositoryController struct {
	RepositoryRepository repository.RepositoryRepository
	TaskRepository       repository.TaskRepository
	Logger               zerolog.Logger
}

func (hc *RepositoryController) ViewAll(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "Projects list"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	repositories, err := hc.RepositoryRepository.GetAllRepositories()
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
		repository.SSHKeyPassphrase = c.FormValue("ssh_key_passphrase")

		err := hc.RepositoryRepository.CreateRepository(repository)
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

	return c.Render("templates/repository_new", d, "templates/layouts/main")
}

func (hc *RepositoryController) UpdateRepositoryById(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "New repository"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	repo, err := hc.RepositoryRepository.GetRepositoryById(c.Params("id"))
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.repository.update_repository_by_id").Msg("failed to get repository")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to get repository",
		})
		return c.Redirect("/app/repository")
	}

	if c.Method() == "POST" {
		repo.Name = c.FormValue("name")
		repo.Description = c.FormValue("description")
		repo.Url = c.FormValue("url")
		repo.SSHKey = c.FormValue("ssh_key")
		repo.SSHKeyPassphrase = c.FormValue("ssh_key_passphrase")

		err := hc.RepositoryRepository.UpdateRepository(repo)
		if err != nil {
			hc.Logger.Error().Err(err).Str("event", "frontend.repository.update_repository_by_id").Msg("failed to update repository")
			d["Error"] = "Internal server error - failed to update repository"
			return c.Render("templates/repository_new_update", d, "templates/layouts/main")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "success-flash",
			Value: "Repository updated successfully",
		})

		return c.Redirect("/app/repository")
	}

	d["Repository"] = repo

	return c.Render("templates/repository_update", d, "templates/layouts/main")
}

func (hc *RepositoryController) DeleteRepositoryById(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	repoId := c.Params("id")

	count, err := hc.TaskRepository.CountTasksByRepositoryId(repoId)
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.repository.delete_repository_by_id").Msg("failed to count tasks")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to get count tasks using this repository",
		})
		return c.Redirect("/app/repository")
	}

	if count > 0 {
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Repository is being used by tasks",
		})
		return c.Redirect("/app/repository")
	}

	err = hc.RepositoryRepository.DeleteRepository(repoId)
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.repository.delete_repository_by_id").Msg("failed to delete repository")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to delete repository",
		})
		return c.Redirect("/app/repository")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "success-flash",
		Value: "Repository deleted successfully",
	})

	return c.Redirect("/app/repository")
}

func (hc *RepositoryController) TestCloneRepositoryById(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	repoId := c.Params("id")
	err := hc.RepositoryRepository.CloneTestRepository(repoId)
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.repository.test_clone_repository_by_id").Msg("failed to clone repository")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to clone repository",
		})
		return c.Redirect("/app/repository")
	}

	c.Cookie(&fiber.Cookie{
		Name:  "success-flash",
		Value: "Repository cloned successfully",
	})

	return c.Redirect("/app/repository")
}
