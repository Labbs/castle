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
	DeleteAt  time.Time `json:"-" gorm:"index"`
}
