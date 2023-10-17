package repository

import (
	"github.com/labbs/castle/modules/project/domain"
	"gorm.io/gorm"
)

type EnvironmentRepository struct {
	database *gorm.DB
}

func NewEnvironmentRepository(database *gorm.DB) EnvironmentRepository {
	return EnvironmentRepository{database: database}
}

func (d *EnvironmentRepository) GetAllEnvironmentsByProjectId(projectId string) ([]domain.Environment, error) {
	var environments []domain.Environment
	r := d.database.Where("project_id = ?", projectId).Find(&environments)
	return environments, r.Error
}

func (d *EnvironmentRepository) CreateEnvironment(environment domain.Environment) error {
	r := d.database.Create(&environment)
	return r.Error
}

func (d *EnvironmentRepository) GetEnvironmentById(id string) (domain.Environment, error) {
	r := domain.Environment{}
	res := d.database.Where("id = ?", id).First(&r)
	return r, res.Error
}

func (d *EnvironmentRepository) EditEnvironment(environment domain.Environment) error {
	r := d.database.Save(&environment)
	return r.Error
}

func (d *EnvironmentRepository) DeleteEnvironment(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Environment{})
	return r.Error
}
