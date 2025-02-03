package auth

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"project-control-system/internal/entities"
	"project-control-system/internal/repositories/user"
	"project-control-system/pkg/jwt_utils"
)

var (
	userWithEmailExist           = errors.New("user with email exist")
	userWithUsernameExist        = errors.New("user with username exist")
	userWithUsernameDoesNotExist = errors.New("user with username does not")
	invalidPasswordError         = errors.New("invalid password")
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("[auth][register]: cannot hash password: %v", err)
	}

	newUser := entities.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(newUser)
}

func (s *Service) Login(dto entities.UserLoginDTO) (entities.LoginResponse, error) {
	isExist, err := s.repo.ExistsByUsername(dto.Username)
	if err != nil {
		return entities.LoginResponse{}, fmt.Errorf("[auth][login]: cannot check if userFromDb exists by username: %v", err)
	}

	if !isExist {
		return entities.LoginResponse{}, userWithUsernameDoesNotExist
	}

	userFromDb, err := s.repo.GetByUsername(dto.Username)
	if err != nil {
		return entities.LoginResponse{}, fmt.Errorf("[auth][login]: cannot get userFromDb by username: %v", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.LoginResponse{}, fmt.Errorf("[auth][login]: cannot hash password: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), hashedPassword)
	if err != nil {
		return entities.LoginResponse{}, invalidPasswordError
	}

	token, err := jwt_utils.GenerateAccessToken(dto.Username)
	if err != nil {
		return entities.LoginResponse{}, fmt.Errorf("[auth][login]: cannot generate access token: %v", err)
	}

	return entities.LoginResponse{Token: token}, nil
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
