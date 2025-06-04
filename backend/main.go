package main

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/wnmay/songhub/backend/internal/config"
	"github.com/wnmay/songhub/backend/internal/handler"
	"github.com/wnmay/songhub/backend/internal/migration"
	"github.com/wnmay/songhub/backend/internal/repository"
	"github.com/wnmay/songhub/backend/internal/usecase"
)

func main() {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Println("No .env file found")
	}

	db := config.ConnectDB()
	migration.Migrate(db)


	app := fiber.New()

	authRepo := repository.NewGormAuthRepository(db)
	authService := usecase.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	app.Post("api/auth/register", authHandler.Register)
	
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
