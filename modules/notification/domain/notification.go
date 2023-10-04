package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Notification struct {
	Id      string `json:"id" gorm:"primaryKey"`
	Message string `json:"message"`

	NotificationChannel NotificationChannel `json:"notification_channel"`

	TaskId        string `json:"task_id"`
	ProjectId     string `json:"project_id"`
	EnvironmentId string `json:"environment_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-" gorm:"index"`
}

func (sla *NotificationChannel) Scan(value interface{}) error {
	var skills NotificationChannel
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla NotificationChannel) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}
