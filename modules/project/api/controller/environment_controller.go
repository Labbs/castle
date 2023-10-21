package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/castle/modules/project/repository"
	"github.com/rs/zerolog"
)

type EnvironmentController struct {
	Repository repository.EnvironmentRepository
	Logger     zerolog.Logger
}

func (ec *EnvironmentController) GetAllEnvironments(c *fiber.Ctx) error {
	projectId := c.Params("id")
	environments, err := ec.Repository.GetAllEnvironmentsByProjectId(projectId)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.get_all").Msg("failed to get all environments")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all environments", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"environments": environments,
	})
}

func (ec *EnvironmentController) GetEnvironmentById(c *fiber.Ctx) error {
	environmentId := c.Params("id")
	environment, err := ec.Repository.GetEnvironmentById(environmentId)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.get_by_id").Msg("failed to get environment by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get environment by id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"environment": environment,
	})
}
