package usecase

import (
	"github.com/wnmay/songhub/backend/internal/entities"
)


type AuthUsecase interface {
	Register(user entities.User) error
}

type AuthRepository interface{
	Create(user entities.User) error
}

type AuthService struct{
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthUsecase{
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user entities.User) error{
	return s.repo.Create(user)
}