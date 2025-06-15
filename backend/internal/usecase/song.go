package usecase

import "github.com/wnmay/songhub/backend/internal/entities"

type SongUsecase interface {
	CreateSong(song entities.Song) error
}

type SongRepository interface {
	CreateSong(song entities.Song) error
	GetArtistByUserID(userID uint) (entities.Artist, error)
}

type SongService struct{
	repo SongRepository
}

func NewSongService(repo SongRepository) SongUsecase{
	return &SongService{repo : repo}
}

func (s *SongService) CreateSong(song entities.Song) error{
	mainArtist, err := s.repo.GetArtistByUserID(song.MainArtistID)
	if err != nil {
		return err
	}

	featuringList := []entities.Artist{}
	for _, artist := range song.FeaturingArtists {
		if artist.UserID != mainArtist.UserID { 
			featuringList = append(featuringList, artist)
		}
	}
	song.FeaturingArtists = featuringList

	return s.repo.CreateSong(song)
}
