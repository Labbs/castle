package repository

import (
	"fmt"

	"github.com/goccy/go-json"
	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/frontend/domain"
)

type RepositoryRepository struct {
	BusMessages chan initBootstrap.Message
}

func NewRepositoryRepository(busMessages chan initBootstrap.Message) RepositoryRepository {
	return RepositoryRepository{BusMessages: busMessages}
}

func (d *RepositoryRepository) GetAllRepositories() ([]domain.BusRepositoryResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "repository:get_all", Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return []domain.BusRepositoryResponse{}, response.Error
	}

	var repositories []domain.BusRepositoryResponse
	err := json.Unmarshal(response.Data.([]byte), &repositories)
	if err != nil {
		return []domain.BusRepositoryResponse{}, err
	}

	return repositories, nil
}

func (d *RepositoryRepository) GetRepositoryById(id string) (domain.BusRepositoryResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "repository:get_by_id", Data: id, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return domain.BusRepositoryResponse{}, response.Error
	}

	var repository domain.BusRepositoryResponse
	err := json.Unmarshal(response.Data.([]byte), &repository)
	if err != nil {
		return domain.BusRepositoryResponse{}, err
	}

	return repository, nil
}

func (d *RepositoryRepository) CreateRepository(repository domain.BusRepositoryResponse) error {
	responseChan := make(chan initBootstrap.Message)

	data, err := json.Marshal(repository)
	if err != nil {
		return err
	}

	d.BusMessages <- initBootstrap.Message{Action: "repository:create", Data: data, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "repository created" {
		return fmt.Errorf("repository not created")
	}

	return nil
}

func (d *RepositoryRepository) UpdateRepository(repository domain.BusRepositoryResponse) error {
	responseChan := make(chan initBootstrap.Message)

	data, err := json.Marshal(repository)
	if err != nil {
		return err
	}

	d.BusMessages <- initBootstrap.Message{Action: "repository:update", Data: data, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "repository updated" {
		return fmt.Errorf("repository not updated")
	}

	return nil
}

func (d *RepositoryRepository) DeleteRepository(id string) error {
	responseChan := make(chan initBootstrap.Message)

	d.BusMessages <- initBootstrap.Message{Action: "repository:delete", Data: id, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "repository deleted" {
		return fmt.Errorf("repository not deleted")
	}

	return nil
}
