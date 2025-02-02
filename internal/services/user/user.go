package user

import (
	"project-control-system/internal/entities"
	"project-control-system/internal/repositories/user"
)

type UserService interface {
	GetUser(username string) (entities.User, error)
	GetUserTickets(username string) ([]entities.UserTicket, error)
}

type Service struct {
	repo *user.UserRepo
}

var _ UserService = (*Service)(nil)

func New(repo *user.UserRepo) *Service {
	return &Service{repo}
}

func (s *Service) GetUser(username string) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetUserTickets(username string) ([]entities.UserTicket, error) {
	//TODO implement me
	panic("implement me")
}
