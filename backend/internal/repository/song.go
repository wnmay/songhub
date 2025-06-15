package repository

import (
	"github.com/wnmay/songhub/backend/internal/entities"
	"github.com/wnmay/songhub/backend/internal/usecase"
	"gorm.io/gorm"
)

type GormSongRepository struct {
	db *gorm.DB
}

func NewGormSongRepository (db *gorm.DB) usecase.SongRepository{
	return &GormSongRepository{db:db}
}

func (r *GormSongRepository) CreateSong(song entities.Song) error {
	return r.db.Create(&song).Error
}

func (r *GormSongRepository) GetArtistByUserID(userID uint) (entities.Artist, error) {
	var artist entities.Artist
	if err := r.db.Where("user_id = ?", userID).First(&artist).Error; err != nil {
		return entities.Artist{}, err
	}
	return artist, nil
}
