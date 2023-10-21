package bus

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/labbs/castle/modules/project/domain"
)

// Bus handler for environment:get_by_id
func (uc *Controller) GetEnvironmentById(data interface{}) interface{} {
	id := data.(string)
	r, err := uc.EnvironmentRepository.GetEnvironmentById(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(r)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}

// Bus handler for environment:get_all
func (uc *Controller) GetAllEnvironments(data interface{}) interface{} {
	id := data.(string)
	r, err := uc.EnvironmentRepository.GetAllEnvironmentsByProjectId(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(r)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}

// Bus handler for environment:create
func (uc *Controller) CreateEnvironment(data interface{}) interface{} {
	var environment domain.Environment
	err := json.Unmarshal(data.([]byte), &environment)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	environment.Id = utils.UUID()

	err = uc.EnvironmentRepository.CreateEnvironment(environment)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "environment created"}
}

// Bus handler for environment:update
func (uc *Controller) UpdateEnvironment(data interface{}) interface{} {
	var environment domain.Environment
	err := json.Unmarshal(data.([]byte), &environment)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	err = uc.EnvironmentRepository.EditEnvironment(environment)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "environment updated"}
}

// Bus handler for environment:delete
func (uc *Controller) DeleteEnvironment(data interface{}) interface{} {
	id := data.(string)
	err := uc.EnvironmentRepository.DeleteEnvironment(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "environment deleted"}
}
