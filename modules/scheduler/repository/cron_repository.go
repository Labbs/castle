package repository

import "github.com/go-co-op/gocron"

type CronRepository struct {
	cronScheduler *gocron.Scheduler
}

func NewCronRepository(cronScheduler *gocron.Scheduler) CronRepository {
	return CronRepository{cronScheduler: cronScheduler}
}
