package repository

import (
	"fmt"

	"github.com/goccy/go-json"
	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/frontend/domain"
)

type EnvironmentRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewEnvironmentRepository(busMessages chan initBootstrap.Message) EnvironmentRepository {
	return EnvironmentRepository{BusMessages: busMessages}
}

func (d *EnvironmentRepository) GetAllEnvironmentsByProjectId(projectId string) ([]domain.BusEnvironmentResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "environment:get_all", Data: projectId, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return []domain.BusEnvironmentResponse{}, response.Error
	}

	var environments []domain.BusEnvironmentResponse
	err := json.Unmarshal(response.Data.([]byte), &environments)
	if err != nil {
		return []domain.BusEnvironmentResponse{}, err
	}

	return environments, nil
}

func (d *EnvironmentRepository) GetEnvironmentById(id string) (domain.BusEnvironmentResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "environment:get_by_id", Data: id, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return domain.BusEnvironmentResponse{}, response.Error
	}

	var environment domain.BusEnvironmentResponse
	err := json.Unmarshal(response.Data.([]byte), &environment)
	if err != nil {
		return domain.BusEnvironmentResponse{}, err
	}

	return environment, nil
}

func (d *EnvironmentRepository) CreateEnvironment(environment domain.BusEnvironmentResponse) error {
	responseChan := make(chan initBootstrap.Message)

	data, err := json.Marshal(environment)
	if err != nil {
		return err
	}

	d.BusMessages <- initBootstrap.Message{Action: "environment:create", Data: data, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "environment created" {
		return fmt.Errorf("environment not created")
	}

	return nil
}

func (d *EnvironmentRepository) UpdateEnvironment(environment domain.BusEnvironmentResponse) error {
	responseChan := make(chan initBootstrap.Message)

	data, err := json.Marshal(environment)
	if err != nil {
		return err
	}

	d.BusMessages <- initBootstrap.Message{Action: "environment:update", Data: data, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "environment updated" {
		return fmt.Errorf("environment not updated")
	}

	return nil
}

func (d *EnvironmentRepository) DeleteEnvironment(id string) error {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "environment:delete", Data: id, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "environment deleted" {
		return fmt.Errorf("environment not deleted")
	}

	return nil
}
