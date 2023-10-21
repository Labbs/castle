package domain

import "time"

type BusGetUserByUsernameResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	DarkMode string `json:"dark_mode"`
}

type BusProjectsResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BusDeleteProjectResponse struct {
	Id string `json:"id"`
}

type BusRepositoryResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	SSHKey      string `json:"ssh_key"`
	SSHKeyPass  string `json:"ssh_key_pass"`
}

type BusTaskResponse struct {
	Id             string        `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	ProjectId      string        `json:"project_id"`
	EnvironmentId  string        `json:"environment_id"`
	Variables      []interface{} `json:"variables,omitempty"`
	RepositoryId   string        `json:"repository_id"`
	RepositoryPath string        `json:"repository_path"`
	Enabled        bool          `json:"enabled"`
	TaskType       string        `json:"task_type"`
	Cron           string        `json:"cron,omitempty"`
	ExecType       string        `json:"exec_type"`
	AnsibleTask    AnsibleTask   `json:"ansible_task,omitempty"`
	TerraformTask  TerraformTask `json:"terraform_task,omitempty"`
	ScriptTask     ScriptTask    `json:"script_task,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	DeleteAt       time.Time     `json:"-" gorm:"index"`
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

type ScriptTask struct {
	Path string `json:"path"`
}

type BusEnvironmentResponse struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	ProjectId   string        `json:"project_id"`
	Variables   []interface{} `json:"variables,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeleteAt    time.Time     `json:"-"`
}
