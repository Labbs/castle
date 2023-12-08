package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/labbs/castle/modules/repository/domain"
	"github.com/rs/zerolog"
)

type RepositoryController struct {
	Service domain.RepositoryService
	Logger  zerolog.Logger
}

func (rc *RepositoryController) GetAllRepositories(c *fiber.Ctx) error {
	repositories, err := rc.Service.GetAllRepositories()
	if err != nil {
		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.get_all").Msg("failed to get all repositories")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all repositories", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"repositories": repositories,
	})
}

func (rc *RepositoryController) CreateRepository(c *fiber.Ctx) error {
	repository := new(domain.Repository)
	if err := c.BodyParser(repository); err != nil {
		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.create").Msg("failed to parse repository")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse repository", "status": "error"})
	}

	repository.Id = utils.UUID()

	if err := rc.Service.CreateRepository(*repository); err != nil {
		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.create").Msg("failed to create repository")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create repository", "status": "error"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"repository": repository,
		"message":    "Repository created",
	})
}

func (rc *RepositoryController) GetRepositoryById(c *fiber.Ctx) error {
	repositoryID := c.Params("id")
	repository, err := rc.Service.GetRepositoryById(repositoryID)
	if err != nil {
		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.get_by_id").Msg("failed to get repository by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get repository by id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"repository": repository,
	})
}

func (rc *RepositoryController) EditRepository(c *fiber.Ctx) error {
	repository := new(domain.Repository)
	if err := c.BodyParser(repository); err != nil {
		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.edit").Msg("failed to parse repository")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse repository", "status": "error"})
	}

	if err := rc.Service.UpdateRepository(*repository); err != nil {
		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.edit").Msg("failed to edit repository")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit repository", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"repository": repository,
		"message":    "Repository edited",
	})
}

// func (rc *RepositoryController) DeleteRepository(c *fiber.Ctx) error {
// 	repositoryID := c.Params("id")

// 	taskCount, err := rc.Repository.CountTasksByRepositoryId(repositoryID)
// 	if err != nil {
// 		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.delete").Msg("failed to count tasks by repository id")
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to check if the repository is used", "status": "error"})
// 	}

// 	if taskCount > 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Repository is used by tasks",
// 		})
// 	}

// 	if err := rc.Repository.DeleteRepository(repositoryID); err != nil {
// 		rc.Logger.Error().Err(err).Str("event", "api.controller.repository.delete").Msg("failed to delete repository")
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete repository", "status": "error"})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "Repository deleted",
// 	})
// }
