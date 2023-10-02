package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/labbs/castle/modules/core/domain"
	"github.com/labbs/castle/modules/core/repository"
	"github.com/rs/zerolog"
)

type EnvironmentController struct {
	Repository repository.CoreRepository
	Logger     zerolog.Logger
}

func (ec *EnvironmentController) GetAllEnvironmentsByProjectId(c *fiber.Ctx) error {
	projectId := c.Params("id")
	environments, err := ec.Repository.GetAllEnvironmentsByProjectId(projectId)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.get_all_by_project_id").Msg("failed to get all environments by project id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all environments by project id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"environments": environments,
	})
}

func (ec *EnvironmentController) CreateEnvironment(c *fiber.Ctx) error {
	environment := new(domain.Environment)
	if err := c.BodyParser(environment); err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.create").Msg("failed to parse environment")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse environment", "status": "error"})
	}

	environment.Id = utils.UUID()

	err := ec.Repository.CreateEnvironment(*environment)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.create").Msg("failed to create environment")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create environment", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"environment": environment,
	})
}

func (ec *EnvironmentController) EditEnvironment(c *fiber.Ctx) error {
	environmentId := c.Params("id")
	request := new(domain.Environment)
	if err := c.BodyParser(request); err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.edit").Msg("failed to parse environment")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse environment", "status": "error"})
	}

	environment, err := ec.Repository.GetEnvironmentById(environmentId)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.edit").Msg("failed to get environment by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get environment by id", "status": "error"})
	}

	environment.Name = request.Name
	environment.Description = request.Description

	err = ec.Repository.EditEnvironment(environment)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.edit").Msg("failed to edit environment")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit environment", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"environment": environment,
	})
}

func (ec *EnvironmentController) DeleteEnvironment(c *fiber.Ctx) error {
	id := c.Params("id")

	err := ec.Repository.DeleteEnvironment(id)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.delete").Msg("failed to delete environment")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete environment", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Environment deleted successfully",
	})
}

func (ec *EnvironmentController) GetEnvironmentById(c *fiber.Ctx) error {
	id := c.Params("id")

	environment, err := ec.Repository.GetEnvironmentById(id)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.get_by_id").Msg("failed to get environment by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get environment by id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"environment": environment,
	})
}
