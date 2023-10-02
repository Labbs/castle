package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type TaskHistory struct {
	Id            string     `gorm:"primaryKey" json:"id"`
	Task          TaskStruct `json:"task"`
	ProjectId     string     `json:"project_id"`
	EnvironmentId string     `json:"environment_id"`
	RepositoryId  string     `json:"repository_id"`

	ExecutionOutput string `json:"execution_return"`
	ExecutionError  string `json:"execution_error"`
	ExecutionTime   string `json:"execution_time"`

	Parameters ParameterStruct `json:"parameters"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}

type ParameterStruct []Parameter
type TaskStruct Task

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (sla *ParameterStruct) Scan(value interface{}) error {
	var skills []Parameter
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla ParameterStruct) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

func (sla *TaskStruct) Scan(value interface{}) error {
	skill := Task{}
	err := json.Unmarshal([]byte(value.(string)), &skill)
	if err != nil {
		return err
	}
	*sla = TaskStruct(skill)
	return nil
}

func (sla TaskStruct) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}
