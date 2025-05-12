package main

import (
	"items-store-service/internal/config"
	kafkaconnect "items-store-service/internal/kafka"
	"items-store-service/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	kafkaconnect.KafkaConnection()

	app := fiber.New()
	routes.ItemStoreRoute(app)
	app.Listen(":8081")
}
