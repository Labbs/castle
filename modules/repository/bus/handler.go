package bus

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/utils"
	pb "github.com/labbs/castle/gen/repository"
	"github.com/labbs/castle/modules/repository/internal"
)

// Bus handler for repository:get_by_id
func (uc *RepositoryController) GetRepositoryById(data interface{}) interface{} {
	id := data.(string)
	repo, err := uc.Repository.GetRepositoryById(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(repo)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}

// Bus handler for repository:get_all
func (uc *RepositoryController) GetAllRepositories(data interface{}) interface{} {
	repos, err := uc.Repository.GetAllRepositories()
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	j, err := json.Marshal(repos)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return j
}

// Bus handler for repository:create
func (uc *RepositoryController) CreateRepository(data interface{}) interface{} {
	var repo pb.Repository
	err := json.Unmarshal(data.([]byte), &repo)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	repo.Id = utils.UUID()

	err = uc.Repository.CreateRepository(&repo)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "repository created"}
}

// Bus handler for repository:update
func (uc *RepositoryController) UpdateRepository(data interface{}) interface{} {
	var repo pb.Repository
	err := json.Unmarshal(data.([]byte), &repo)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	err = uc.Repository.UpdateRepository(&repo)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "repository updated"}
}

// Bus handler for repository:delete
func (uc *RepositoryController) DeleteRepository(data interface{}) interface{} {
	id := data.(string)
	err := uc.Repository.DeleteRepository(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "repository deleted"}
}

// Bus handler for repository:clone
func (uc *RepositoryController) CloneRepository(data interface{}) interface{} {
	id := data.(string)
	repo, err := uc.Repository.GetRepositoryById(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	err = internal.CloneRepository(repo, false)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "repository cloned"}
}

// Bus handler for repository:clone_test
func (uc *RepositoryController) CloneTestRepository(data interface{}) interface{} {
	id := data.(string)
	repo, err := uc.Repository.GetRepositoryById(id)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	err = internal.CloneRepository(repo, true)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}

	return map[string]string{"success": "repository cloned"}
}
