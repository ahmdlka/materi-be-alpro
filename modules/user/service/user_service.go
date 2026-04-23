package service

import (
	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/Mobilizes/materi-be-alpro/modules/user/dto"
	"github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	"github.com/Mobilizes/materi-be-alpro/pkg/helpers"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*entities.User, error) {
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	err = s.repo.Create(user)
	return user, err
}

func (s *UserService) Get(id int) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
	return response, err
}

func (s *UserService) GetAll() ([]dto.UserResponse, error) {
	users, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}
	var responses []dto.UserResponse

	for _, user := range users {
		userRes := dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		}
		responses = append(responses, userRes)
	}
	return responses, err
}
