package usecase

import (
	"github.com/wnmay/songhub/backend/internal/entities"
	"github.com/wnmay/songhub/backend/pkg/hash"
	"github.com/wnmay/songhub/backend/pkg/jwt"
)


type AuthUsecase interface {
	Register(user entities.User) error
	Login(email,password string) (string, error)

}

type AuthRepository interface{
	Create(user *entities.User) error
	GetEmail(email,password string) (*entities.User, error)
	CreateListener(listener *entities.Listener) error
	CreateArtist(artist *entities.Artist) error
}

type AuthService struct{
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthUsecase{
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user entities.User) error {
	hashed, err := hash.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	if err := s.repo.Create(&user); err != nil {
		return err
	}
	
	switch user.Role {
	case entities.RoleListener:
		listener := &entities.Listener{UserID: user.ID}
		if err := s.repo.CreateListener(listener); err != nil {
			return err
		}
	case entities.RoleArtist:
		artist := &entities.Artist{UserID: user.ID}
		if err := s.repo.CreateArtist(artist); err != nil {
			return err
		}
	}

	return nil
}


func (s *AuthService) Login(email,password string) (string, error){
	user, err := s.repo.GetEmail(email,password)
	if err!=nil{
		return "",err
	}

	if err := hash.CheckPassword(user.Password, password); err != nil {
		return "", err
	}
	return jwt.GenerateToken(user.ID, string(user.Role))
}
