package auth

import (
	"errors"
	"fmt"
	"project-control-system/internal/entities"
	"project-control-system/internal/repositories/user"
)

var (
	userWithEmailExist    = errors.New("user with email exist")
	userWithUsernameExist = errors.New("user with username exist")
)

type AuthService interface {
	Register(user entities.UserDTO) error
	Login(user entities.UserLoginDTO) (entities.LoginResponse, error)
	GetCurrentUser() (entities.UserDTO, error)
	HasRole(user entities.UserDTO, projectId, role string) bool
	HasAnyRole(user entities.UserDTO, projectId string, role []string) bool
}

type Service struct {
	repo user.UserRepo
}

var _ AuthService = (*Service)(nil)

func New(repo user.UserRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(user entities.UserDTO) error {
	isExist, err := s.repo.ExistsByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("[auth][register]: cannot check if user exists by email: %v", err)
	}

	if isExist {
		return userWithEmailExist
	}

	isExist, err = s.repo.ExistsByUsername(user.Username)
	if err != nil {
		return fmt.Errorf("[auth][register]: cannot check if user exists by username: %v", err)
	}

	if isExist {
		return userWithUsernameExist
	}

	newUser := entities.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	return s.repo.Create(newUser)
}

func (s *Service) Login(user entities.UserLoginDTO) (entities.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetCurrentUser() (entities.UserDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) HasRole(user entities.UserDTO, projectId, role string) bool {
	//TODO implement me
	panic("implement me")
}

func (s *Service) HasAnyRole(user entities.UserDTO, projectId string, role []string) bool {
	//TODO implement me
	panic("implement me")
}
