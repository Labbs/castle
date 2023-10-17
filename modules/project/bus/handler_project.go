package bus

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/labbs/castle/modules/project/domain"
)

// Bus handler for project:get_by_id
func (uc *Controller) GetProjectById(data interface{}) interface{} {
	id := data.(string)
	r, err := uc.ProjectRepository.GetProjectById(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(r)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}

// Bus handler for project:get_all
func (uc *Controller) GetAllProjects(data interface{}) interface{} {
	r, err := uc.ProjectRepository.GetAllProjects()
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(r)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}

// Bus handler for project:create
func (uc *Controller) CreateProject(data interface{}) interface{} {
	var project domain.Project
	err := json.Unmarshal(data.([]byte), &project)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	project.Id = utils.UUID()
	project.Variables = domain.VariablesList{}

	err = uc.ProjectRepository.CreateProject(project)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "project created"}
}

// Bus handler for project:update
func (uc *Controller) UpdateProject(data interface{}) interface{} {
	var project domain.Project
	err := json.Unmarshal(data.([]byte), &project)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	err = uc.ProjectRepository.EditProject(project)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "project updated"}
}

// Bus handler for project:delete
func (uc *Controller) DeleteProject(data interface{}) interface{} {
	id := data.(string)
	err := uc.ProjectRepository.DeleteProject(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "project deleted"}
}
