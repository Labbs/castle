package repository

import (
	"github.com/labbs/castle/modules/task/domain"
	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(database *gorm.DB) TaskRepository {
	return TaskRepository{database: database}
}

func (d *TaskRepository) GetAllTasksByProjectId(projectId string) ([]domain.Task, error) {
	var tasks []domain.Task
	r := d.database.Where("project_id = ?", projectId).Find(&tasks)
	return tasks, r.Error
}

func (d *TaskRepository) GetTasksByRepositoryId(id string) error {
	var tasks []domain.Task
	r := d.database.Where("repository_id = ?", id).Find(&tasks)
	return r.Error
}

func (d *TaskRepository) CountTasksByRepositoryId(id string) (int64, error) {
	var count int64
	r := d.database.Model(&domain.Task{}).Where("repository_id = ?", id).Count(&count)
	return count, r.Error
}

func (d *TaskRepository) GetTaskById(id string) (domain.Task, error) {
	r := domain.Task{}
	res := d.database.Where("id = ?", id).First(&r)
	return r, res.Error
}

func (d *TaskRepository) CreateTask(task domain.Task) error {
	r := d.database.Create(&task)
	return r.Error
}

func (d *TaskRepository) EditTask(task domain.Task) error {
	r := d.database.Save(&task)
	return r.Error
}

func (d *TaskRepository) DeleteTask(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Task{})
	return r.Error
}

func (d *TaskRepository) GetAllEnabledCronTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	r := d.database.Where("enabled = ? AND task_type = ?", true, "cron").Find(&tasks)
	return tasks, r.Error
}

func (d *TaskRepository) GetAllTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	r := d.database.Find(&tasks)
	return tasks, r.Error
}
