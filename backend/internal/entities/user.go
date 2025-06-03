package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `gorm:"unique;not null"`
	Email    string   `gorm:"unique;not null"`
	Password string   `gorm:"not null"`
	Role     UserRole `gorm:"type:varchar(20);not null;default:'customer'"`
}