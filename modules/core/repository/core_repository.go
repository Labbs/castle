package repository

import (
	"github.com/labbs/castle/modules/core/domain"
	"gorm.io/gorm"
)

type CoreRepository struct {
	database *gorm.DB
}

func NewCoreRepository(database *gorm.DB) CoreRepository {
	return CoreRepository{database: database}
}

func (d *CoreRepository) GetAllProjects() ([]domain.Project, error) {
	var projects []domain.Project
	r := d.database.Find(&projects)
	return projects, r.Error
}

func (d *CoreRepository) GetProjectById(id string) (domain.Project, error) {
	p := domain.Project{}
	r := d.database.Where("id = ?", id).First(&p)
	return p, r.Error
}

func (d *CoreRepository) CreateProject(project domain.Project) error {
	r := d.database.Create(&project)
	return r.Error
}

func (d *CoreRepository) EditProject(project domain.Project) error {
	r := d.database.Save(&project)
	return r.Error
}

func (d *CoreRepository) DeleteProject(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Project{})
	return r.Error
}

func (d *CoreRepository) GetAllTasksByProjectId(projectId string) ([]domain.Task, error) {
	var tasks []domain.Task
	r := d.database.Where("project_id = ?", projectId).Find(&tasks)
	return tasks, r.Error
}

func (d *CoreRepository) GetAllRepositories() ([]domain.Repository, error) {
	var repositories []domain.Repository
	r := d.database.Find(&repositories)
	return repositories, r.Error
}

func (d *CoreRepository) CreateRepository(repository domain.Repository) error {
	r := d.database.Create(&repository)
	return r.Error
}

func (d *CoreRepository) GetRepositoryById(id string) (domain.Repository, error) {
	r := domain.Repository{}
	res := d.database.Where("id = ?", id).First(&r)
	return r, res.Error
}

func (d *CoreRepository) EditRepository(repository domain.Repository) error {
	r := d.database.Save(&repository)
	return r.Error
}

func (d *CoreRepository) DeleteRepository(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Repository{})
	return r.Error
}

func (d *CoreRepository) GetTasksByRepositoryId(id string) error {
	var tasks []domain.Task
	r := d.database.Where("repository_id = ?", id).Find(&tasks)
	return r.Error
}

func (d *CoreRepository) CountTasksByRepositoryId(id string) (int64, error) {
	var count int64
	r := d.database.Model(&domain.Task{}).Where("repository_id = ?", id).Count(&count)
	return count, r.Error
}

func (d *CoreRepository) GetAllVariablesByProjectId(projectId string) ([]domain.Variable, error) {
	var variables []domain.Variable
	r := d.database.Where("project_id = ?", projectId).Find(&variables)
	return variables, r.Error
}

func (d *CoreRepository) GetAllVariablesByEnvironmentId(environmentId string) ([]domain.Variable, error) {
	var variables []domain.Variable
	r := d.database.Where("environment_id = ?", environmentId).Find(&variables)
	return variables, r.Error
}

func (d *CoreRepository) GetAllVariablesByTaskId(taskId string) ([]domain.Variable, error) {
	var variables []domain.Variable
	r := d.database.Where("task_id = ?", taskId).Find(&variables)
	return variables, r.Error
}

func (d *CoreRepository) CreateVariable(variable domain.Variable) error {
	r := d.database.Create(&variable)
	return r.Error
}

func (d *CoreRepository) GetVariableById(id string) (domain.Variable, error) {
	r := domain.Variable{}
	res := d.database.Where("id = ?", id).First(&r)
	return r, res.Error
}

func (d *CoreRepository) EditVariable(variable domain.Variable) error {
	r := d.database.Save(&variable)
	return r.Error
}

func (d *CoreRepository) DeleteVariable(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Variable{})
	return r.Error
}

func (d *CoreRepository) GetAllEnvironmentsByProjectId(projectId string) ([]domain.Environment, error) {
	var environments []domain.Environment
	r := d.database.Where("project_id = ?", projectId).Find(&environments)
	return environments, r.Error
}

func (d *CoreRepository) CreateEnvironment(environment domain.Environment) error {
	r := d.database.Create(&environment)
	return r.Error
}

func (d *CoreRepository) GetEnvironmentById(id string) (domain.Environment, error) {
	r := domain.Environment{}
	res := d.database.Where("id = ?", id).First(&r)
	return r, res.Error
}

func (d *CoreRepository) EditEnvironment(environment domain.Environment) error {
	r := d.database.Save(&environment)
	return r.Error
}

func (d *CoreRepository) DeleteEnvironment(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Environment{})
	return r.Error
}

func (d *CoreRepository) GetTaskById(id string) (domain.Task, error) {
	r := domain.Task{}
	res := d.database.Where("id = ?", id).First(&r)
	return r, res.Error
}

func (d *CoreRepository) CreateTask(task domain.Task) error {
	r := d.database.Create(&task)
	return r.Error
}

func (d *CoreRepository) EditTask(task domain.Task) error {
	r := d.database.Save(&task)
	return r.Error
}

func (d *CoreRepository) DeleteTask(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Task{})
	return r.Error
}

func (d *CoreRepository) GetAllEnabledTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	r := d.database.Where("enabled = ?", true).Find(&tasks)
	return tasks, r.Error
}

func (d *CoreRepository) CreateTaskHistory(taskHistory domain.TaskHistory) error {
	r := d.database.Create(&taskHistory)
	return r.Error
}
