package bugReport

import (
	"project-control-system/internal/entities"
	"project-control-system/internal/repositories/bugReport"
)

type BugReportService interface {
	Create(dto entities.CreateBugReportDTO, projectId string) (entities.BugReportDTO, error)
	UpdateBugReportStatus(dto entities.UpdateBugReportDTO, projectId string) error
}

type Service struct {
	repo bugReport.BugReportRepo
}

var _ BugReportService = (*Service)(nil)

func New(repo *bugReport.BugReportRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(dto entities.CreateBugReportDTO, projectId string) (entities.BugReportDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateBugReportStatus(dto entities.UpdateBugReportDTO, projectId string) error {
	//TODO implement me
	panic("implement me")
}
