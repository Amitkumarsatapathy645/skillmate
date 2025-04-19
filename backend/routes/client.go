package routes

import (
	"github.com/amit645/skillmate-backend/controllers"
	"github.com/amit645/skillmate-backend/middlewares"
	"github.com/amit645/skillmate-backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupClientRoutes(app *fiber.App, clientService *services.ClientService) {
	clientController := controllers.NewClientController(clientService)

	// Public route for browsing services
	app.Get("/api/services", clientController.BrowseServices)

	// Protected route for posting requests
	api := app.Group("/api/requests", middlewares.AuthClient)
	api.Post("/", clientController.CreateRequest)
}
