package service

import "github.com/labbs/castle/modules/repository/domain"

type repositoryService struct {
	repositoryRepository domain.RepositoryRepository
}

func NewRepositoryService(repositoryRepository domain.RepositoryRepository) domain.RepositoryService {
	return &repositoryService{repositoryRepository: repositoryRepository}
}

func (s *repositoryService) GetRepositoryById(id string) (domain.Repository, error) {
	return s.repositoryRepository.GetRepositoryById(id)
}

func (s *repositoryService) GetRepositoryByName(name string) (domain.Repository, error) {
	return s.repositoryRepository.GetRepositoryByName(name)
}

func (s *repositoryService) GetAllRepositories() ([]domain.Repository, error) {
	return s.repositoryRepository.GetAllRepositories()
}

func (s *repositoryService) CreateRepository(repository domain.Repository) error {
	return s.repositoryRepository.CreateRepository(repository)
}

func (s *repositoryService) UpdateRepository(repository domain.Repository) error {
	return s.repositoryRepository.UpdateRepository(repository)
}

func (s *repositoryService) DeleteRepository(id string) error {
	return s.repositoryRepository.DeleteRepository(id)
}
