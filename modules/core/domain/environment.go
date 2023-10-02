package domain

import "time"

type Environment struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"primaryKey" json:"name"`
	Description string `json:"description,omitempty"`

	ProjectId string     `json:"project_id"`
	Variables []Variable `json:"variables,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}
