package domain

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Task struct {
	Id          string `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`

	ProjectId     string     `json:"project_id"`
	EnvironmentId string     `json:"environment_id"`
	Variables     []Variable `json:"variables,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	RepositoryId   string `json:"repository_id"`
	RepositoryPath string `json:"repository_path"`

	Enabled bool `json:"enabled"`

	Cron string `json:"cron,omitempty"`

	TaskType      string        `json:"task_type"`
	AnsibleTask   AnsibleTask   `json:"ansible_task,omitempty"`
	TerraformTask TerraformTask `json:"terraform_task,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}

type AnsibleTask struct {
	DryRun        bool   `json:"dry_run"`
	Debug         bool   `json:"debug"`
	Verbose       bool   `json:"verbose"`
	PlaybookPath  string `json:"playbook_path"`
	InventoryPath string `json:"inventory_path"`
}

type TerraformTask struct {
	Workspace string `json:"workspace"`
	PlanOnly  bool   `json:"plan_only"`
}

func (sla *AnsibleTask) Scan(value interface{}) error {
	var skills AnsibleTask
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla AnsibleTask) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

func (sla *TerraformTask) Scan(value interface{}) error {
	var skills TerraformTask
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla TerraformTask) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}
