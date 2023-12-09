package service

import "github.com/labbs/castle/modules/project/domain"

type environmentService struct {
	environmentRepository domain.EnvironmentRepository
}

func NewEnvironmentService(environmentRepository domain.EnvironmentRepository) domain.EnvironmentService {
	return &environmentService{environmentRepository: environmentRepository}
}

func (s *environmentService) GetAllEnvironmentsByProjectId(projectId string) ([]domain.Environment, error) {
	return s.environmentRepository.GetAllEnvironmentsByProjectId(projectId)
}

func (s *environmentService) CreateEnvironment(environment domain.Environment) error {
	return s.environmentRepository.CreateEnvironment(environment)
}

func (s *environmentService) GetEnvironmentById(id string) (domain.Environment, error) {
	return s.environmentRepository.GetEnvironmentById(id)
}

func (s *environmentService) EditEnvironment(environment domain.Environment) error {
	return s.environmentRepository.EditEnvironment(environment)
}

func (s *environmentService) DeleteEnvironment(id string) error {
	return s.environmentRepository.DeleteEnvironment(id)
}
