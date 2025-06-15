package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wnmay/songhub/backend/internal/entities"
	"github.com/wnmay/songhub/backend/internal/usecase"
)

type SongHandler struct {
	songUseCase usecase.SongUsecase
}

func NewSongHandler(useCase usecase.SongUsecase) *SongHandler {
	return &SongHandler{songUseCase: useCase}
}

func (h *SongHandler) CreateSong(c *fiber.Ctx) error {
	type SongRequest struct {
		Title              string             `json:"title"`
		Genre              entities.SongGenre `json:"genre"`
		FeaturingArtist    []uint             `json:"featuring_artist"`
	}

	var req SongRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	userIDFromLocal := c.Locals("user_id")
	userID, ok := userIDFromLocal.(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	featuringArtists := make([]entities.Artist, 0, len(req.FeaturingArtist))
	for _, id := range req.FeaturingArtist {
		featuringArtists = append(featuringArtists, entities.Artist{UserID: id})
	}

	song := entities.Song{
		Title:            req.Title,
		Genre:            req.Genre,
		MainArtistID:     uint(userID),
		FeaturingArtists: featuringArtists,
	}

	if err := h.songUseCase.CreateSong(song); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Song created successfully"})
}

