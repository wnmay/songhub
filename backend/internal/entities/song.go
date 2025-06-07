package entities

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Title 		 string    `gorm:"not null"`
	Genre        SongGenre `gorm:"type:varchar(20);not null"`
	MainArtistID uint      `gorm:"not null"`
	MainArtist   Artist    `gorm:"foreignKey:MainArtistID;constraint:OnDelete:CASCADE;"`

	FeaturingArtists []Artist `gorm:"many2many:song_featuring"`
}
