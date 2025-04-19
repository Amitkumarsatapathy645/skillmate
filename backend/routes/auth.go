package routes

import (
	"github.com/amit645/skillmate-backend/controllers"
	"github.com/amit645/skillmate-backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, authService *services.AuthService) {
	authController := controllers.NewAuthController(authService)

	app.Post("/api/auth/register", authController.Register)
	app.Post("/api/auth/login", authController.Login)
}
