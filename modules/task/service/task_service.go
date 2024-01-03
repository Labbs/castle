package service

import "github.com/labbs/castle/modules/task/domain"

type taskService struct {
	taskRepository domain.TaskRepository
}

func NewTaskService(taskRepository domain.TaskRepository) domain.TaskService {
	return &taskService{taskRepository: taskRepository}
}

func (s *taskService) GetAllTasksByProjectId(projectId string) ([]domain.Task, error) {
	return s.taskRepository.GetAllTasksByProjectId(projectId)
}

func (s *taskService) GetTasksByRepositoryId(id string) error {
	return s.taskRepository.GetTasksByRepositoryId(id)
}

func (s *taskService) CountTasksByRepositoryId(id string) (int64, error) {
	return s.taskRepository.CountTasksByRepositoryId(id)
}

func (s *taskService) GetTaskById(id string) (domain.Task, error) {
	return s.taskRepository.GetTaskById(id)
}

func (s *taskService) CreateTask(task domain.Task) error {
	return s.taskRepository.CreateTask(task)
}

func (s *taskService) EditTask(task domain.Task) error {
	return s.taskRepository.EditTask(task)
}

func (s *taskService) DeleteTask(id string) error {
	return s.taskRepository.DeleteTask(id)
}

func (s *taskService) GetAllEnabledCronTasks() ([]domain.Task, error) {
	return s.taskRepository.GetAllEnabledCronTasks()
}

func (s *taskService) GetAllTasks() ([]domain.Task, error) {
	return s.taskRepository.GetAllTasks()
}

func (s *taskService) CreateSchedulerTask(task domain.Task) error {
	return s.taskRepository.CreateSchedulerTask(task)
}
