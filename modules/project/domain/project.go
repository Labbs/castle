package domain

import (
	"time"
)

type Project struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Environments []Environment `json:"environments,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Tasks        []interface{} `json:"tasks,omitempty" gorm:"-"`
	Variables    VariablesList `json:"variables,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

func (p *Project) TableName() string {
	return "project"
}

type ProjectRepository interface {
	GetProjectById(id string) (Project, error)
	GetAllProjects() ([]Project, error)
	CreateProject(project Project) error
	EditProject(project Project) error
	DeleteProject(id string) error
}

type ProjectService interface {
	GetProjectById(id string) (Project, error)
	GetAllProjects() ([]Project, error)
	CreateProject(project Project) error
	EditProject(project Project) error
	DeleteProject(id string) error
}
