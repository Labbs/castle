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
	DeleteAt  time.Time `json:"-" gorm:"index"`
}
