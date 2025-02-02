package project

import (
	"project-control-system/internal/entities"
	"project-control-system/internal/repositories/project"
)

type ProjectService interface {
	CreateProject(dto entities.CreateProjectDTO) (entities.ProjectDTO, error)
	GetUserProject(username string) (entities.UserProjectDTO, error)
	AssignDeveloper(projectId string, dto entities.AssignUserDTO) error
	AssignTeamLead(projectId string, dto entities.AssignUserDTO) error
	AssignQA(projectId string, dto entities.AssignUserDTO) error
	TestProject(projectId string) error
}

type Service struct {
	repo *project.ProjectRepo
}

var _ ProjectService = (*Service)(nil)

func New(repo *project.ProjectRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProject(dto entities.CreateProjectDTO) (entities.ProjectDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetUserProject(username string) (entities.UserProjectDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AssignDeveloper(projectId string, dto entities.AssignUserDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AssignTeamLead(projectId string, dto entities.AssignUserDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AssignQA(projectId string, dto entities.AssignUserDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) TestProject(projectId string) error {
	//TODO implement me
	panic("implement me")
}
