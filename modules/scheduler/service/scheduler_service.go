package service

import (
	"github.com/labbs/castle/modules/scheduler/domain"
)

type schedulerService struct {
	schedulerRepository domain.SchedulerRepository
}

func NewSchedulerService(schedulerRepository domain.SchedulerRepository) domain.SchedulerService {
	return &schedulerService{schedulerRepository: schedulerRepository}
}

func (s *schedulerService) AddSchedulerCronJob(task domain.Task) error {
	return s.schedulerRepository.AddSchedulerCronJob(task)
}

func (s *schedulerService) RemoveSchedulerCronJob(jobId string) error {
	return s.schedulerRepository.RemoveSchedulerCronJob(jobId)
}

func (s *schedulerService) GetAllSchedulerCronJobs() ([]domain.Task, error) {
	return s.schedulerRepository.GetAllSchedulerCronJobs()
}

func (s *schedulerService) RunSchedulerCronJob(jobId string) error {
	return s.schedulerRepository.RunSchedulerCronJob(jobId)
}
