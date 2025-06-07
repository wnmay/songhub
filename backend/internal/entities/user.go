package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string   `gorm:"unique;not null" json:"username"`
	Email    string   `gorm:"unique;not null" json:"email"`
	Password string   `gorm:"not null" json:"password"`
	Role     UserRole `gorm:"type:varchar(20);not null;default:'customer'" json:"role"`

	Artist   *Artist   `gorm:"foreignKey:UserID" json:"artist,omitempty"`
	Listener *Listener `gorm:"foreignKey:UserID" json:"listener,omitempty"`
}