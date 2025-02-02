package milestone

import (
	"project-control-system/internal/entities"
	"project-control-system/internal/repositories/milestone"
)

type MilestoneService interface {
	CreateMilestone(dto entities.CreateMilestoneDTO, projectId string) (entities.MilestoneDTO, error)
	UpdateMilestone(milestoneId string, dto entities.MilestoneUpdateStatusDTO) error
}

type Service struct {
	repo *milestone.MilestoneRepo
}

var _ MilestoneService = (*Service)(nil)

func New(repo *milestone.MilestoneRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateMilestone(dto entities.CreateMilestoneDTO, projectId string) (entities.MilestoneDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateMilestone(milestoneId string, dto entities.MilestoneUpdateStatusDTO) error {
	//TODO implement me
	panic("implement me")
}
