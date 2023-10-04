package domain

import "time"

type Project struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	Environments []interface{} `json:"environments,omitempty" gorm:"-"`
	Tasks        []interface{} `json:"tasks,omitempty" gorm:"-"`
	Variables    []interface{} `json:"variables,omitempty" gorm:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}
