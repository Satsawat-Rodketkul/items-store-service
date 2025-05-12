package handlers

import (
	"items-store-service/internal/models"
	"items-store-service/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func ItemListHandler(c *fiber.Ctx) error {
	response, err := service.ItemListService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func ItemBuyHandler(c *fiber.Ctx) error {
	request := models.ItemBuyRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		log.Fatal("Error BodyParser")
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	err = service.ItemBuyService(request)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("error")
	}

	return c.Status(fiber.StatusOK).JSON("success")
}
