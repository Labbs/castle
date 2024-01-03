package domain

type SchedulerService interface {
	AddSchedulerCronJob(task Task) error
	RemoveSchedulerCronJob(jobId string) error
	GetAllSchedulerCronJobs() ([]Task, error)
	RunSchedulerCronJob(jobId string) error
}

type SchedulerRepository interface {
	AddSchedulerCronJob(task Task) error
	RemoveSchedulerCronJob(jobId string) error
	GetAllSchedulerCronJobs() ([]Task, error)
	RunSchedulerCronJob(jobId string) error
}
