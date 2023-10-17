package repository

import (
	"fmt"

	"github.com/goccy/go-json"
	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/frontend/domain"
)

type TaskRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewTaskRepository(busMessages chan initBootstrap.Message) TaskRepository {
	return TaskRepository{BusMessages: busMessages}
}

func (d *TaskRepository) GetAllTasksByProjectId(id string) ([]domain.BusTaskResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "task:get_all_by_project_id", Data: id, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return []domain.BusTaskResponse{}, response.Error
	}

	var tasks []domain.BusTaskResponse
	err := json.Unmarshal(response.Data.([]byte), &tasks)
	if err != nil {
		return []domain.BusTaskResponse{}, err
	}

	return tasks, nil
}

func (d *TaskRepository) CreateTask(task domain.BusTaskResponse) error {
	responseChan := make(chan initBootstrap.Message)

	data, err := json.Marshal(task)
	if err != nil {
		return err
	}

	d.BusMessages <- initBootstrap.Message{Action: "task:create", Data: data, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "task created" {
		return fmt.Errorf("task not created")
	}

	return nil
}
