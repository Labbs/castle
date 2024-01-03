package repository

import (
	"github.com/go-co-op/gocron"
	"github.com/labbs/castle/modules/scheduler/domain"
	"github.com/labbs/castle/modules/scheduler/internal"
	"github.com/rs/zerolog"
)

type schedulerRepository struct {
	cronScheduler *gocron.Scheduler
	logger        zerolog.Logger
}

func NewSchedulerRepository(cronScheduler *gocron.Scheduler, logger zerolog.Logger) *schedulerRepository {
	return &schedulerRepository{cronScheduler: cronScheduler, logger: logger}
}

func (r *schedulerRepository) AddSchedulerCronJob(task domain.Task) error {
	_, err := r.cronScheduler.Cron(task.Cron).Tag(task.Id).SingletonMode().Do(internal.Execute, task.Id, r.logger)
	return err
}

func (r *schedulerRepository) RemoveSchedulerCronJob(jobId string) error {
	return r.cronScheduler.RemoveByTag(jobId)
}

func (r *schedulerRepository) GetAllSchedulerCronJobs() ([]domain.Task, error) {
	return nil, nil
}

func (r *schedulerRepository) RunSchedulerCronJob(jobId string) error {
	return r.cronScheduler.RunByTag(jobId)
}
