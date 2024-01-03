package repository

import (
	"fmt"

	"github.com/goccy/go-json"
	initBootstrap "github.com/labbs/castle/bootstrap"
	"github.com/labbs/castle/modules/frontend/domain"
	"github.com/rs/zerolog"
)

type ProjectRepository struct {
	BusMessages chan initBootstrap.Message
	Logger      zerolog.Logger
}

func NewProjectRepository(busMessages chan initBootstrap.Message, logger zerolog.Logger) ProjectRepository {
	return ProjectRepository{BusMessages: busMessages, Logger: logger}
}

func (d *ProjectRepository) GetAllProjects() ([]domain.BusProjectsResponse, error) {
	d.Logger.Debug().Str("event", "repository.project.get_all").Msg("requesting projects from bus")

	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "project:get_all", Response: responseChan}
	response := <-responseChan
	d.Logger.Debug().Str("event", "repository.project.get_all").Interface("response", response).Msg("response received")
	if response.Error != nil {
		return []domain.BusProjectsResponse{}, response.Error
	}
	d.Logger.Debug().Str("event", "repository.project.get_all").Interface("projects", response.Data).Msg("projects found")

	var projects []domain.BusProjectsResponse
	err := json.Unmarshal(response.Data.([]byte), &projects)
	if err != nil {
		return []domain.BusProjectsResponse{}, err
	}
	d.Logger.Debug().Str("event", "repository.project.get_all").Interface("projects", projects).Msg("projects unmarshalled")

	return projects, nil
}

func (d *ProjectRepository) GetProjectById(id string) (domain.BusProjectsResponse, error) {
	responseChan := make(chan initBootstrap.Message)
	d.BusMessages <- initBootstrap.Message{Action: "project:get_by_id", Data: id, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return domain.BusProjectsResponse{}, response.Error
	}

	var project domain.BusProjectsResponse
	err := json.Unmarshal(response.Data.([]byte), &project)
	if err != nil {
		return domain.BusProjectsResponse{}, err
	}

	return project, nil
}

func (d *ProjectRepository) CreateProject(project domain.BusProjectsResponse) error {
	responseChan := make(chan initBootstrap.Message)

	data, err := json.Marshal(project)
	if err != nil {
		return err
	}

	d.BusMessages <- initBootstrap.Message{Action: "project:create", Data: data, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "project created" {
		return fmt.Errorf("project not created")
	}

	return nil
}

func (d *ProjectRepository) UpdateProject(project domain.BusProjectsResponse) error {
	responseChan := make(chan initBootstrap.Message)

	data, err := json.Marshal(project)
	if err != nil {
		return err
	}

	d.BusMessages <- initBootstrap.Message{Action: "project:update", Data: data, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "project updated" {
		return fmt.Errorf("project not updated")
	}

	return nil
}

func (d *ProjectRepository) DeleteProject(id string) error {
	responseChan := make(chan initBootstrap.Message)

	d.BusMessages <- initBootstrap.Message{Action: "project:delete", Data: id, Response: responseChan}
	response := <-responseChan
	if response.Error != nil {
		return response.Error
	}

	resp := response.Data.(map[string]string)

	if resp["success"] != "project deleted" {
		return fmt.Errorf("project not deleted")
	}

	return nil
}
