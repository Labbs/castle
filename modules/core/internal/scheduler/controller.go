package scheduler

import (
	"errors"

	"github.com/go-co-op/gocron"
	"github.com/labbs/castle/modules/core/domain"
	"github.com/labbs/castle/modules/core/repository"
	"github.com/rs/zerolog"
)

type SchedulerController struct {
	Repository repository.CoreRepository
	Scheduler  *gocron.Scheduler
	Logger     zerolog.Logger
}

func (sc *SchedulerController) Start() {
	sc.Logger.Info().Str("event", "scheduler.controller.start").Msg("starting scheduler")
	tasks, err := sc.Repository.GetAllEnabledTasks()
	if err != nil {
		sc.Logger.Error().Err(err).Str("event", "scheduler.controller.start").Msg("failed to get all enabled tasks")
		return
	}

	sc.Logger.Info().Str("event", "scheduler.controller.start").Msg("adding tasks to scheduler")
	for _, task := range tasks {
		err := sc.scheduleCron(task)
		if err != nil {
			sc.Logger.Error().Err(err).Str("event", "scheduler.controller.start").Str("task_id", task.Id).Msg("failed to schedule task")
		}
	}
}

func (sc *SchedulerController) Stop() {
	sc.Logger.Info().Str("event", "scheduler.controller.stop").Msg("stopping scheduler")
	sc.Scheduler.Stop()
}

func (sc *SchedulerController) Add(task domain.Task) error {
	sc.Logger.Info().Str("event", "scheduler.controller.add").
		Str("task", task.Id).Msg("adding task to scheduler")
	return sc.scheduleCron(task)
}

func (sc *SchedulerController) Remove(id string) error {
	sc.Logger.Info().Str("event", "scheduler.controller.remove").
		Str("task", id).Msg("removing task from scheduler")
	return sc.Scheduler.RemoveByTag(id)
}

func (sc *SchedulerController) scheduleCron(task domain.Task) error {
	switch task.TaskType {
	case "ansible":
		_, err := sc.Scheduler.Cron(task.Cron).Tag(task.Id).Do(sc.ansibleTask, task)
		return err
	case "terraform":
		_, err := sc.Scheduler.Cron(task.Cron).Tag(task.Id).Do(sc.terraformTask, task)
		return err
	default:
		return errors.New("invalid task type")
	}
}
