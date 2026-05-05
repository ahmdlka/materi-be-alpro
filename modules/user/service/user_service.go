package service

import (
	"github.com/Mobilizes/materi-be-alpro/database/entities"
	"github.com/Mobilizes/materi-be-alpro/modules/user/dto"
	"github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	"github.com/Mobilizes/materi-be-alpro/pkg/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (s *UserService) Get(id uuid.UUID) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return s.mapToUserResponse(user), nil
}

func (s *UserService) GetAll() ([]dto.UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.UserResponse
	for _, user := range users {
		responses = append(responses, *s.mapToUserResponse(&user))
	}
	return responses, nil
}

func (s *UserService) UpdateProfile(userID uuid.UUID, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	profile, err := s.repo.GetProfileByUserID(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		profile = &entities.UserProfile{
			UserID: userID.String(),
		}
	}

	profile.Age = req.Age
	profile.WeightKg = req.WeightKg
	profile.HeightCm = req.HeightCm
	profile.Gender = req.Gender
	profile.ActivityLevel = req.ActivityLevel
	profile.Goal = req.Goal
	profile.Allergies = req.Allergies
	profile.Diseases = req.Diseases
	profile.FoodPreferences = req.FoodPreferences

	if err := s.repo.UpdateProfile(profile); err != nil {
		return nil, err
	}

	return s.Get(userID)
}

func (s *UserService) mapToUserResponse(user *entities.User) *dto.UserResponse {
	res := &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if user.Profile.ID != uuid.Nil {
		res.Profile = &dto.ProfileResponse{
			Age:             user.Profile.Age,
			WeightKg:        user.Profile.WeightKg,
			HeightCm:        user.Profile.HeightCm,
			Gender:          user.Profile.Gender,
			ActivityLevel:   user.Profile.ActivityLevel,
			Goal:            user.Profile.Goal,
			Allergies:       user.Profile.Allergies,
			Diseases:        user.Profile.Diseases,
			FoodPreferences: user.Profile.FoodPreferences,
		}
	}

	return res
}
