package routes

import (
	"items-store-service/internal/handlers"
	"items-store-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ItemStoreRoute(app *fiber.App) {
	itemGroup := app.Group("/item", middleware.ValidateToken)
	itemGroup.Get("/list", handlers.ItemListHandler)
	itemGroup.Post("/buy", handlers.ItemBuyHandler)
}
