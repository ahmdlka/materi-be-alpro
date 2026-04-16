package service

import (
	"errors"


	"github.com/Mobilizes/materi-be-alpro/modules/auth/dto"
	userRepo "github.com/Mobilizes/materi-be-alpro/modules/user/repository"
	"github.com/Mobilizes/materi-be-alpro/pkg/helpers"
)

type AuthService struct {
	userRepo   *userRepo.UserRepository
	jwtService *JWTService
}

func NewAuthService(userRepo *userRepo.UserRepository, jwtService *JWTService) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *AuthService) Login(req *dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	if !helpers.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("email atau password salah")
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return "", errors.New("gagal membuat sesi")
	}

	return token, nil
}
