package controllers

import (
	"github.com/amit645/skillmate-backend/services"
	"github.com/gofiber/fiber/v2"
)

type ClientController struct {
	clientService *services.ClientService
}

func NewClientController(clientService *services.ClientService) *ClientController {
	return &ClientController{clientService: clientService}
}

func (c *ClientController) CreateRequest(ctx *fiber.Ctx) error {
	type RequestInput struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Skills      []string `json:"skills"`
		Budget      float64  `json:"budget"`
		City        string   `json:"city"`
	}

	var input RequestInput
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	userID := ctx.Locals("user_id").(string)
	request, err := c.clientService.CreateRequest(input.Title, input.Description, input.City, input.Budget, input.Skills, userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(request)
}

func (c *ClientController) BrowseServices(ctx *fiber.Ctx) error {
	skill := ctx.Query("skill")
	city := ctx.Query("city")
	minPrice := ctx.QueryFloat("min_price", 0)
	maxPrice := ctx.QueryFloat("max_price", 0)

	services, err := c.clientService.BrowseServices(skill, city, minPrice, maxPrice)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(services)
}
