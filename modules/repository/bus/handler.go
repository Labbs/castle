package bus

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/utils"
	"github.com/labbs/castle/modules/repository/domain"
	"github.com/labbs/castle/modules/repository/internal"
)

// Bus handler for repository:get_by_id
func (uc *RepositoryController) GetRepositoryById(data interface{}) (interface{}, error) {
	id := data.(string)
	repo, err := uc.Repository.GetRepositoryById(id)
	if err != nil {
		return map[string]string{}, err
	}

	j, err := json.Marshal(repo)
	if err != nil {
		return map[string]string{}, err
	}

	return j, nil
}

// Bus handler for repository:get_all
func (uc *RepositoryController) GetAllRepositories(data interface{}) (interface{}, error) {
	repos, err := uc.Repository.GetAllRepositories()
	if err != nil {
		return map[string]string{}, err
	}

	j, err := json.Marshal(repos)
	if err != nil {
		return map[string]string{}, err
	}

	return j, nil
}

// Bus handler for repository:create
func (uc *RepositoryController) CreateRepository(data interface{}) (interface{}, error) {
	var repo domain.Repository
	err := json.Unmarshal(data.([]byte), &repo)
	if err != nil {
		return map[string]string{}, err
	}

	repo.Id = utils.UUID()

	err = uc.Repository.CreateRepository(repo)
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"success": "repository created"}, nil
}

// Bus handler for repository:update
func (uc *RepositoryController) UpdateRepository(data interface{}) (interface{}, error) {
	var repo domain.Repository
	err := json.Unmarshal(data.([]byte), &repo)
	if err != nil {
		return map[string]string{}, err
	}

	err = uc.Repository.UpdateRepository(repo)
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"success": "repository updated"}, nil
}

// Bus handler for repository:delete
func (uc *RepositoryController) DeleteRepository(data interface{}) (interface{}, error) {
	id := data.(string)
	err := uc.Repository.DeleteRepository(id)
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"success": "repository deleted"}, nil
}

// Bus handler for repository:clone
func (uc *RepositoryController) CloneRepository(data interface{}) (interface{}, error) {
	id := data.(string)
	repo, err := uc.Repository.GetRepositoryById(id)
	if err != nil {
		return map[string]string{}, err
	}

	err = internal.CloneRepository(repo, false)
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"success": "repository cloned"}, nil
}

// Bus handler for repository:clone_test
func (uc *RepositoryController) CloneTestRepository(data interface{}) (interface{}, error) {
	id := data.(string)
	repo, err := uc.Repository.GetRepositoryById(id)
	if err != nil {
		return map[string]string{}, err
	}

	err = internal.CloneRepository(repo, true)
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"success": "repository cloned"}, nil
}
