package repository

import (
	pb "github.com/labbs/castle/gen/project"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(database *gorm.DB) ProjectRepository {
	return ProjectRepository{database: database}
}

func (d *ProjectRepository) GetAllProjects() ([]pb.Project, error) {
	var projects []pb.Project
	r := d.database.Find(&projects)
	return projects, r.Error
}

func (d *ProjectRepository) GetProjectById(id string) (pb.Project, error) {
	p := pb.Project{}
	r := d.database.Where("id = ?", id).First(&p)
	return p, r.Error
}

func (d *ProjectRepository) CreateProject(project *pb.Project) error {
	r := d.database.Create(&project)
	return r.Error
}

func (d *ProjectRepository) EditProject(project *pb.Project) error {
	r := d.database.Save(&project)
	return r.Error
}

func (d *ProjectRepository) DeleteProject(id string) error {
	r := d.database.Where("id = ?", id).Delete(&pb.Project{})
	return r.Error
}
