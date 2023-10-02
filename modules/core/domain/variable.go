package domain

import "time"

type Variable struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`

	EnvironmentId string `json:"environment_id,omitempty"`
	TaskId        string `json:"task_id,omitempty"`
	ProjectId     string `json:"project_id,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}
