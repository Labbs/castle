package domain

type Task struct {
	Id       string `json:"id"`
	TaskType string `json:"task_type"`
	Cron     string `json:"cron,omitempty"`
	ExecType string `json:"exec_type"`
}
