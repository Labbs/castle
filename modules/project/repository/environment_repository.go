package repository

import (
	"github.com/labbs/castle/modules/project/domain"
	"gorm.io/gorm"
)

type environmentRepository struct {
	database *gorm.DB
}

func NewEnvironmentRepository(database *gorm.DB) *environmentRepository {
	return &environmentRepository{database: database}
}

func (d *environmentRepository) GetAllEnvironmentsByProjectId(projectId string) ([]domain.Environment, error) {
	var environments []domain.Environment
	r := d.database.Where("project_id = ?", projectId).Find(&environments)
	return environments, r.Error
}

func (d *environmentRepository) CreateEnvironment(environment domain.Environment) error {
	r := d.database.Create(&environment)
	return r.Error
}

func (d *environmentRepository) GetEnvironmentById(id string) (domain.Environment, error) {
	r := domain.Environment{}
	res := d.database.Where("id = ?", id).First(&r)
	return r, res.Error
}

func (d *environmentRepository) EditEnvironment(environment domain.Environment) error {
	r := d.database.Save(&environment)
	return r.Error
}

func (d *environmentRepository) DeleteEnvironment(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Environment{})
	return r.Error
}
