package service

import "github.com/labbs/castle/modules/project/domain"

type projectService struct {
	projectRepository domain.ProjectRepository
}

func NewProjectService(projectRepository domain.ProjectRepository) domain.ProjectService {
	return &projectService{projectRepository: projectRepository}
}

func (s *projectService) GetProjectById(id string) (domain.Project, error) {
	return s.projectRepository.GetProjectById(id)
}

func (s *projectService) GetAllProjects() ([]domain.Project, error) {
	return s.projectRepository.GetAllProjects()
}

func (s *projectService) CreateProject(project domain.Project) error {
	return s.projectRepository.CreateProject(project)
}

func (s *projectService) EditProject(project domain.Project) error {
	return s.projectRepository.EditProject(project)
}

func (s *projectService) DeleteProject(id string) error {
	return s.projectRepository.DeleteProject(id)
}
