package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/frontend/domain"
	"github.com/labbs/castle/modules/frontend/repository"
	"github.com/rs/zerolog"
)

type ProjectController struct {
	ProjectRepository repository.ProjectRepository
	TaskRepository    repository.TaskRepository
	Logger            zerolog.Logger
}

func (hc *ProjectController) ViewAll(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "Projects list"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	projects, err := hc.ProjectRepository.GetAllProjects()
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.project.view_all").Msg("failed to get projects")
		d["Error"] = "Internal server error - failed to get projects"
	}
	d["Projects"] = projects

	return c.Render("templates/project_view_list", d, "templates/layouts/main")
}

func (hc *ProjectController) Create(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "New project"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	if c.Method() == "POST" {
		var project domain.BusProjectsResponse
		project.Name = c.FormValue("name")
		project.Description = c.FormValue("description")

		err := hc.ProjectRepository.CreateProject(project)
		if err != nil {
			hc.Logger.Error().Err(err).Str("event", "frontend.project.create").Msg("failed to create project")
			d["Error"] = "Internal server error - failed to create project"
			return c.Render("templates/project_new_update", d, "templates/layouts/main")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "success-flash",
			Value: "Project created successfully",
		})

		return c.Redirect("/app/project")
	}

	return c.Render("templates/project_new", d, "templates/layouts/main")
}

func (hc *ProjectController) GetProjectById(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "Project details"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	project, err := hc.ProjectRepository.GetProjectById(c.Params("id"))
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.project.get_project_by_id").Msg("failed to get project")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to get project",
		})
		return c.Redirect("/app/project")
	}

	d["Project"] = project

	tasks, err := hc.TaskRepository.GetAllTasksByProjectId(project.Id)
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.project.get_project_by_id").Msg("failed to get tasks")
		d["Error"] = "Internal server error - failed to get tasks"
	}

	d["Tasks"] = tasks
	d["Histories"] = []interface{}{}

	return c.Render("templates/project_view_details", d, "templates/layouts/main")
}

func (hc *ProjectController) UpdateProjectById(c *fiber.Ctx) error {
	d := c.Locals("latestUserData").(fiber.Map)
	d["Title"] = "Update project"
	d["Error"] = c.Cookies("error-flash")
	d["Success"] = c.Cookies("success-flash")
	c.ClearCookie("error-flash", "success-flash")

	project, err := hc.ProjectRepository.GetProjectById(c.Params("id"))
	if err != nil {
		hc.Logger.Error().Err(err).Str("event", "frontend.project.update_project_by_id").Msg("failed to get project")
		c.Cookie(&fiber.Cookie{
			Name:  "error-flash",
			Value: "Internal server error - failed to get project",
		})
		return c.Redirect("/app/project")
	}

	if c.Method() == "POST" {
		project.Name = c.FormValue("name")
		project.Description = c.FormValue("description")

		err := hc.ProjectRepository.UpdateProject(project)
		if err != nil {
			hc.Logger.Error().Err(err).Str("event", "frontend.project.update_project_by_id").Msg("failed to update project")
			d["Error"] = "Internal server error - failed to update project"
			return c.Render("templates/project_new_update", d, "templates/layouts/main")
		}

		c.Cookie(&fiber.Cookie{
			Name:  "success-flash",
			Value: "Project updated successfully",
		})

		return c.Redirect("/app/project/" + project.Id + "/view")
	}

	d["Project"] = project

	return c.Render("templates/project_update", d, "templates/layouts/main")
}
