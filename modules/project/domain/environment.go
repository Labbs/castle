package domain

import "time"

type Environment struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	ProjectId string        `json:"project_id,omitempty" gorm:"index"`
	Variables VariablesList `json:"variables,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"index"`
}

func (e *Environment) TableName() string {
	return "environment"
}

type EnvironmentRepository interface {
	GetEnvironmentById(id string) (Environment, error)
	GetAllEnvironmentsByProjectId(projectId string) ([]Environment, error)
	CreateEnvironment(environment Environment) error
	EditEnvironment(environment Environment) error
	DeleteEnvironment(id string) error
}

type EnvironmentService interface {
	GetEnvironmentById(id string) (Environment, error)
	GetAllEnvironmentsByProjectId(projectId string) ([]Environment, error)
	CreateEnvironment(environment Environment) error
	EditEnvironment(environment Environment) error
	DeleteEnvironment(id string) error
}
