package scheduler

import (
	"github.com/gofiber/fiber/v2/utils"
	"github.com/labbs/castle/modules/core/domain"
)

func (sc *SchedulerController) terraformTask(task domain.Task) {
	sc.Logger.Info().Str("event", "scheduler.controller.terraform_task").Str("task_id", task.Id).Msg("running terraform task")
	history := domain.TaskHistory{
		Id:            utils.UUID(),
		Task:          domain.TaskStruct(task),
		ProjectId:     task.ProjectId,
		EnvironmentId: task.EnvironmentId,
		RepositoryId:  task.RepositoryId,
		Status:        "started",
	}

	if err := sc.Repository.CreateTaskHistory(history); err != nil {
		sc.Logger.Error().Err(err).Str("event", "scheduler.controller.terraform_task").Str("task_id", task.Id).Msg("failed to create task history")
		return
	}
}
