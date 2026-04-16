package service

import (
	"context"

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

func (s *UserService) Get(req context.Context, id int) (*entities.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAll(
	req context.Context,
) ([]entities.User, error) {

	return s.repo.FindAll()
}
