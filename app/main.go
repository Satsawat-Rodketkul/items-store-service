package main

import (
	"encoding/json"
	"items-store-service/internal/config"
	"items-store-service/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	app := fiber.New()

	app.Get("/items", func(c *fiber.Ctx) error {
		url := config.GetValue("URL")
		resp, err := http.Get(url)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON("error")
		}
		defer resp.Body.Close()

		response := []models.ItemsResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON("error")
		}

		return c.Status(fiber.StatusOK).JSON(response)
	})

	app.Listen(":8081")
}
