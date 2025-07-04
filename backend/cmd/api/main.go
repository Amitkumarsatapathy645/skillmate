package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amit645/skillmate-backend/config"
	"github.com/amit645/skillmate-backend/routes"
	"github.com/amit645/skillmate-backend/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	serviceService := services.NewServiceService(db)
	clientService := services.NewClientService(db)

	// Initialize Fiber
	app := fiber.New()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Authorization, Content-Type",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	// Setup routes
	routes.SetupAuthRoutes(app, authService)
	routes.SetupServiceRoutes(app, serviceService)
	routes.SetupClientRoutes(app, clientService)

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
