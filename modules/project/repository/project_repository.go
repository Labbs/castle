package repository

import (
	"github.com/labbs/castle/modules/project/domain"
	"gorm.io/gorm"
)

type projectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(database *gorm.DB) *projectRepository {
	return &projectRepository{database: database}
}

func (d *projectRepository) GetAllProjects() ([]domain.Project, error) {
	var projects []domain.Project
	r := d.database.Find(&projects)
	return projects, r.Error
}

func (d *projectRepository) GetProjectById(id string) (domain.Project, error) {
	p := domain.Project{}
	r := d.database.Where("id = ?", id).First(&p)
	return p, r.Error
}

func (d *projectRepository) CreateProject(project domain.Project) error {
	r := d.database.Create(&project)
	return r.Error
}

func (d *projectRepository) EditProject(project domain.Project) error {
	r := d.database.Save(&project)
	return r.Error
}

func (d *projectRepository) DeleteProject(id string) error {
	r := d.database.Where("id = ?", id).Delete(&domain.Project{})
	return r.Error
}
