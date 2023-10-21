package bus

import (
	"encoding/json"

	"github.com/gofiber/utils"
	"github.com/labbs/castle/modules/task/domain"
)

// Bus handler for task:get_by_id
func (uc *Controller) GetTaskById(data interface{}) interface{} {
	task := data.(string)
	r, err := uc.Repository.GetTaskById(task)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return r
}

// Bus handler for task:get_all
func (uc *Controller) GetAllTasks(data interface{}) interface{} {
	r, err := uc.Repository.GetAllTasks()
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return r
}

// Bus handler for task:get_all_by_project_id
func (uc *Controller) GetAllTasksByProjectId(data interface{}) interface{} {
	id := data.(string)
	r, err := uc.Repository.GetAllTasksByProjectId(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(r)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}

// Bus handler for task:create
func (uc *Controller) CreateTask(data interface{}) interface{} {
	var task domain.Task
	err := json.Unmarshal(data.([]byte), &task)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	task.Id = utils.UUID()

	err = uc.Repository.CreateTask(task)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "task created"}
}

// Bus handler for task:count_by_repository_id
func (uc *Controller) CountTasksByRepositoryId(data interface{}) interface{} {
	id := data.(string)
	r, err := uc.Repository.CountTasksByRepositoryId(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]int64{"count": r}
}
