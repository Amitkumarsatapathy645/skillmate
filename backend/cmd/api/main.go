package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amit645/skillmate-backend/config"
	"github.com/amit645/skillmate-backend/routes"
	"github.com/amit645/skillmate-backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	db := config.ConnectDB()

	// Initialize services
	authService := services.NewAuthService(db)

	// Initialize Fiber
	app := fiber.New()

	// Setup routes
	routes.SetupAuthRoutes(app, authService)

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to SkillMate API!",
		})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
