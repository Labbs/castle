package repository

import (
	pb "github.com/labbs/castle/gen/project"
	"gorm.io/gorm"
)

type EnvironmentRepository struct {
	database *gorm.DB
}

func NewEnvironmentRepository(database *gorm.DB) EnvironmentRepository {
	return EnvironmentRepository{database: database}
}

func (d *EnvironmentRepository) GetAllEnvironmentsByProjectId(projectId string) ([]pb.Environment, error) {
	var environments []pb.Environment
	r := d.database.Where("project_id = ?", projectId).Find(&environments)
	return environments, r.Error
}

func (d *EnvironmentRepository) CreateEnvironment(environment *pb.Environment) error {
	r := d.database.Create(&environment)
	return r.Error
}

func (d *EnvironmentRepository) GetEnvironmentById(id string) (*pb.Environment, error) {
	r := pb.Environment{}
	res := d.database.Where("id = ?", id).First(&r)
	return &r, res.Error
}

func (d *EnvironmentRepository) EditEnvironment(environment *pb.Environment) error {
	r := d.database.Save(&environment)
	return r.Error
}

func (d *EnvironmentRepository) DeleteEnvironment(id string) error {
	r := d.database.Where("id = ?", id).Delete(&pb.Environment{})
	return r.Error
}
