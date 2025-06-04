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

func (r *GormAuthRepository) Create(user entities.User) error{
	return r.db.Create(&user).Error
}