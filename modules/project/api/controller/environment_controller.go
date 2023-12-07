package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/labbs/castle/modules/project/domain"
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

func (ec *EnvironmentController) CreateEnvironment(c *fiber.Ctx) error {
	var environment domain.Environment
	err := c.BodyParser(&environment)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.create").Msg("failed to parse environment")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to parse environment", "status": "error"})
	}

	environment.Id = utils.UUID()

	err = ec.Repository.CreateEnvironment(environment)
	if err != nil {
		ec.Logger.Error().Err(err).Str("event", "api.controller.environment.create").Msg("failed to create environment")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create environment", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Environment created",
		"environment": environment,
	})
}
