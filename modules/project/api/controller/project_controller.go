package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	pb "github.com/labbs/castle/gen/project"
	"github.com/labbs/castle/modules/project/repository"
	"github.com/rs/zerolog"
)

type ProjectController struct {
	Repository repository.ProjectRepository
	Logger     zerolog.Logger
}

func (pc *ProjectController) GetAllProjects(c *fiber.Ctx) error {
	projects, err := pc.Repository.GetAllProjects()
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.get_all").Msg("failed to get all projects")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all projects", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"projects": projects,
	})
}

func (pc *ProjectController) GetProjectById(c *fiber.Ctx) error {
	projectID := c.Params("id")
	project, err := pc.Repository.GetProjectById(projectID)
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.get_by_id").Msg("failed to get project by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get project by id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"project": &project,
	})
}

func (pc *ProjectController) CreateProject(c *fiber.Ctx) error {
	project := new(pb.Project)
	if err := c.BodyParser(project); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.create").Msg("failed to parse project")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse project", "status": "error"})
	}

	project.Id = utils.UUID()

	if err := pc.Repository.CreateProject(project); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.create").Msg("failed to create project")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create project", "status": "error"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"project": project,
		"message": "Project created",
	})
}

func (pc *ProjectController) UpdateProject(c *fiber.Ctx) error {
	projectId := c.Params("id")
	request := new(pb.Project)
	if err := c.BodyParser(request); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.update").Msg("failed to parse project")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse project", "status": "error"})
	}

	project, err := pc.Repository.GetProjectById(projectId)
	if err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.update").Msg("failed to get project by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get project by id", "status": "error"})
	}

	project.Name = request.Name
	project.Description = request.Description

	if err := pc.Repository.EditProject(&project); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.update").Msg("failed to update project")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update project", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"project": &project,
		"message": "Project updated",
	})
}

func (pc *ProjectController) DeleteProject(c *fiber.Ctx) error {
	projectID := c.Params("id")
	if err := pc.Repository.DeleteProject(projectID); err != nil {
		pc.Logger.Error().Err(err).Str("event", "api.controller.project.delete").Msg("failed to delete project")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete project", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Project deleted",
	})
}
