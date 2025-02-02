package ticket

import (
	"project-control-system/internal/entities"
	"project-control-system/internal/repositories/ticket"
)

type TicketService interface {
	Create(createTicket entities.CreateTicket) (entities.TicketDTO, error)
	AssignUserToTicket(ticketId string, dto entities.AssignTicketDTO) error
	GetTicketStatus(ticketId string) (entities.TicketStatusDTO, error)
	UpdateTicketStatus(ticketId string, status entities.TicketStatusDTO) error
}

type Service struct {
	repo *ticket.TicketRepo
}

var _ TicketService = (*Service)(nil)

func New(repo *ticket.TicketRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(createTicket entities.CreateTicket) (entities.TicketDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) AssignUserToTicket(ticketId string, dto entities.AssignTicketDTO) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetTicketStatus(ticketId string) (entities.TicketStatusDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateTicketStatus(ticketId string, status entities.TicketStatusDTO) error {
	//TODO implement me
	panic("implement me")
}
