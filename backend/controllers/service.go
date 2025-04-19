package controllers

import (
	"github.com/amit645/skillmate-backend/services"
	"github.com/gofiber/fiber/v2"
)

type ServiceController struct {
	serviceService *services.ServiceService
}

func NewServiceController(serviceService *services.ServiceService) *ServiceController {
	return &ServiceController{serviceService: serviceService}
}

func (c *ServiceController) CreateService(ctx *fiber.Ctx) error {
	type ServiceRequest struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Price       float64  `json:"price"`
		Tags        []string `json:"tags"`
		City        string   `json:"city"`
	}

	var req ServiceRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID := ctx.Locals("user_id").(string)
	service, err := c.serviceService.CreateService(req.Title, req.Description, req.City, req.Price, req.Tags, userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(service)
}

func (c *ServiceController) GetServices(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(string)
	services, err := c.serviceService.GetServicesByFreelancer(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(services)
}
