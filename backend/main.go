package main

import (
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/wnmay/songhub/backend/internal/config"
)

func main() {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Println("No .env file found")
	}

	db := config.ConnectDB()
	_ = db


	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Backend is alive")
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
