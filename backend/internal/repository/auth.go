package repository

import (
	"github.com/wnmay/songhub/backend/internal/entities"
	"github.com/wnmay/songhub/backend/internal/usecase"
	"gorm.io/gorm"
)

type GormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository (db *gorm.DB ) usecase.AuthRepository{
	return &GormAuthRepository{db:db}
}

func (r *GormAuthRepository) Create(user *entities.User) error{
	return r.db.Create(&user).Error
}

func (r *GormAuthRepository) GetEmail(email, password string) (*entities.User,error){
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err!=nil{
		return nil,err
	}
	return &user,nil
}

func (r *GormAuthRepository) CreateListener(listener *entities.Listener) error {
	return r.db.Create(listener).Error
}

func (r *GormAuthRepository) CreateArtist(artist *entities.Artist) error {
	return r.db.Create(artist).Error
}