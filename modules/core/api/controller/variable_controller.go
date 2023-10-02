package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/labbs/castle/modules/core/domain"
	"github.com/labbs/castle/modules/core/repository"
	"github.com/rs/zerolog"
)

type VariableController struct {
	Repository repository.CoreRepository
	Logger     zerolog.Logger
}

func (vc *VariableController) GetAllVariablesByProjectId(c *fiber.Ctx) error {
	projectId := c.Params("id")
	variables, err := vc.Repository.GetAllVariablesByProjectId(projectId)
	if err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.get_all_by_project_id").Msg("failed to get all variables by project id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all variables by project id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"variables": variables,
	})
}

func (vc *VariableController) GetAllVariablesByEnvironmentId(c *fiber.Ctx) error {
	environmentId := c.Params("id")
	variables, err := vc.Repository.GetAllVariablesByEnvironmentId(environmentId)
	if err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.get_all_by_environment_id").Msg("failed to get all variables by environment id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all variables by environment id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"variables": variables,
	})
}

func (vc *VariableController) GetAllVariablesByTaskId(c *fiber.Ctx) error {
	taskId := c.Params("id")
	variables, err := vc.Repository.GetAllVariablesByTaskId(taskId)
	if err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.get_all_by_task_id").Msg("failed to get all variables by task id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all variables by task id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"variables": variables,
	})
}

func (vc *VariableController) CreateVariable(c *fiber.Ctx) error {
	variable := new(domain.Variable)
	if err := c.BodyParser(variable); err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.create").Msg("failed to parse variable")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse variable", "status": "error"})
	}

	variable.Id = utils.UUID()

	if err := vc.Repository.CreateVariable(*variable); err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.create").Msg("failed to create variable")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create variable", "status": "error"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"variable": variable,
		"message":  "Variable created",
	})
}

func (vc *VariableController) GetVariableById(c *fiber.Ctx) error {
	variableId := c.Params("id")
	variable, err := vc.Repository.GetVariableById(variableId)
	if err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.get_by_id").Msg("failed to get variable by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get variable by id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"variable": variable,
	})
}

func (vc *VariableController) EditVariable(c *fiber.Ctx) error {
	request := new(domain.Variable)
	if err := c.BodyParser(request); err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.edit").Msg("failed to parse variable")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse variable", "status": "error"})
	}

	variable, err := vc.Repository.GetVariableById(request.Id)
	if err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.edit").Msg("failed to get variable by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get variable by id", "status": "error"})
	}

	variable.Name = request.Name
	variable.Value = request.Value

	if err := vc.Repository.EditVariable(variable); err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.edit").Msg("failed to edit variable")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit variable", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Variable edited",
	})
}

func (vc *VariableController) DeleteVariable(c *fiber.Ctx) error {
	variableId := c.Params("id")
	if err := vc.Repository.DeleteVariable(variableId); err != nil {
		vc.Logger.Error().Err(err).Str("event", "api.controller.variable.delete").Msg("failed to delete variable")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete variable", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Variable deleted",
	})
}
