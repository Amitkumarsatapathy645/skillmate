package routes

import (
	"github.com/amit645/skillmate-backend/controllers"
	"github.com/amit645/skillmate-backend/middlewares"
	"github.com/amit645/skillmate-backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupServiceRoutes(app *fiber.App, serviceService *services.ServiceService) {
	serviceController := controllers.NewServiceController(serviceService)

	// Protected routes for freelancers
	api := app.Group("/api/services", middlewares.AuthFreelancer)
	api.Post("/", serviceController.CreateService)
	api.Get("/", serviceController.GetServices)
}
