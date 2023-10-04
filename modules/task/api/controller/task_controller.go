package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/labbs/castle/modules/task/domain"
	"github.com/labbs/castle/modules/task/repository"
	"github.com/rs/zerolog"
)

type TaskController struct {
	Repository repository.TaskRepository
	Logger     zerolog.Logger
	// Scheduler  scheduler.SchedulerController
}

func (tc *TaskController) GetAllTasksByProjectId(c *fiber.Ctx) error {
	projectId := c.Params("id")
	tasks, err := tc.Repository.GetAllTasksByProjectId(projectId)
	if err != nil {
		tc.Logger.Error().Err(err).Str("event", "api.controller.task.get_all_by_project_id").Msg("failed to get all tasks by project id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get all tasks by project id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tasks": tasks,
	})
}

func (tc *TaskController) GetTaskById(c *fiber.Ctx) error {
	taskId := c.Params("id")
	task, err := tc.Repository.GetTaskById(taskId)
	if err != nil {
		tc.Logger.Error().Err(err).Str("event", "api.controller.task.get_by_id").Msg("failed to get task by id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get task by id", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"task": task,
	})
}

func (tc *TaskController) CreateTask(c *fiber.Ctx) error {
	task := new(domain.Task)
	if err := c.BodyParser(task); err != nil {
		tc.Logger.Error().Err(err).Str("event", "api.controller.task.create").Msg("failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse body", "status": "error"})
	}

	task.Id = utils.UUID()

	if err := tc.Repository.CreateTask(*task); err != nil {
		tc.Logger.Error().Err(err).Str("event", "api.controller.task.create").Msg("failed to create task")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create task", "status": "error"})
	}

	// if task.Enabled && task.TaskType == "cron" {
	// 	if err := tc.Scheduler.Add(*task); err != nil {
	// 		tc.Logger.Error().Err(err).Str("event", "api.controller.task.create").Msg("failed to add task to scheduler")
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to add task to scheduler", "status": "error"})
	// 	}
	// }

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created",
	})
}

func (tc *TaskController) EditTask(c *fiber.Ctx) error {
	request := new(domain.Task)
	if err := c.BodyParser(request); err != nil {
		tc.Logger.Error().Err(err).Str("event", "api.controller.task.edit").Msg("failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse body", "status": "error"})
	}

	// task, err := tc.Repository.GetTaskById(request.Id)
	// if err != nil {
	// 	tc.Logger.Error().Err(err).Str("event", "api.controller.task.edit").Msg("failed to get task by id")
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get task by id", "status": "error"})
	// }

	// if task.Enabled && !request.Enabled {
	// 	if err := tc.Scheduler.Remove(request.Id); err != nil {
	// 		tc.Logger.Error().Err(err).Str("event", "api.controller.task.edit").Msg("failed to remove task from scheduler")
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to remove task from scheduler", "status": "error"})
	// 	}
	// } else if !task.Enabled && request.Enabled {
	// 	if err := tc.Scheduler.Add(*request); err != nil {
	// 		tc.Logger.Error().Err(err).Str("event", "api.controller.task.edit").Msg("failed to add task to scheduler")
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to add task to scheduler", "status": "error"})
	// 	}
	// }

	if err := tc.Repository.EditTask(*request); err != nil {
		tc.Logger.Error().Err(err).Str("event", "api.controller.task.edit").Msg("failed to edit task")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to edit task", "status": "error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task edited",
	})
}

func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := tc.Repository.DeleteTask(id); err != nil {
		tc.Logger.Error().Err(err).Str("event", "api.controller.task.delete").Msg("failed to delete task")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete task", "status": "error"})
	}

	// if err := tc.Scheduler.Remove(id); err != nil {
	// 	tc.Logger.Error().Err(err).Str("event", "api.controller.task.delete").Msg("failed to remove task from scheduler")
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to remove task from scheduler", "status": "error"})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
