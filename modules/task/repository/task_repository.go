package repository

import (
	pb "github.com/labbs/castle/gen/task"
	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(database *gorm.DB) TaskRepository {
	return TaskRepository{database: database}
}

func (d *TaskRepository) GetAllTasksByProjectId(projectId string) ([]pb.Task, error) {
	var tasks []pb.Task
	r := d.database.Where("project_id = ?", projectId).Find(&tasks)
	return tasks, r.Error
}

func (d *TaskRepository) GetTasksByRepositoryId(id string) error {
	var tasks []pb.Task
	r := d.database.Where("repository_id = ?", id).Find(&tasks)
	return r.Error
}

func (d *TaskRepository) CountTasksByRepositoryId(id string) (int64, error) {
	var count int64
	r := d.database.Model(&pb.Task{}).Where("repository_id = ?", id).Count(&count)
	return count, r.Error
}

func (d *TaskRepository) GetTaskById(id string) (*pb.Task, error) {
	r := pb.Task{}
	res := d.database.Where("id = ?", id).First(&r)
	return &r, res.Error
}

func (d *TaskRepository) CreateTask(task *pb.Task) error {
	r := d.database.Create(&task)
	return r.Error
}

func (d *TaskRepository) EditTask(task *pb.Task) error {
	r := d.database.Save(&task)
	return r.Error
}

func (d *TaskRepository) DeleteTask(id string) error {
	r := d.database.Where("id = ?", id).Delete(&pb.Task{})
	return r.Error
}

func (d *TaskRepository) GetAllEnabledCronTasks() ([]pb.Task, error) {
	var tasks []pb.Task
	r := d.database.Where("enabled = ? AND task_type = ?", true, "cron").Find(&tasks)
	return tasks, r.Error
}

func (d *TaskRepository) GetAllTasks() ([]pb.Task, error) {
	var tasks []pb.Task
	r := d.database.Find(&tasks)
	return tasks, r.Error
}
