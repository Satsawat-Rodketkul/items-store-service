package main

import (
	client "items-store-service/internal/client/service"
	"items-store-service/internal/config"
	kafkaconnect "items-store-service/internal/kafka"
	kafkaModel "items-store-service/internal/kafka/models"
	"items-store-service/internal/kafka/producer"
	service "items-store-service/internal/service/models"
	"items-store-service/internal/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	kafkaconnect.KafkaConnection()
	app := fiber.New()

	app.Get("/items", func(c *fiber.Ctx) error {
		response, err := client.GetAllItems()
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusInternalServerError).JSON("error")
		}
		return c.Status(fiber.StatusOK).JSON(response)
	})

	app.Post("/item/buy", func(c *fiber.Ctx) error {
		request := service.ItemBuyRequest{}
		err := c.BodyParser(&request)
		if err != nil {
			log.Fatal(err)
			return err
		}
		txnMessage := kafkaModel.TransactionMessage{
			TransactionId: utils.GenerateTransactionId(),
			UserId:        "",
			Title:         request.Title,
			Price:         request.Price,
			Categoty:      request.Categoty,
			CreateDate:    time.Now(),
		}

		err = producer.ProduceMessageTransaction(txnMessage)
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusInternalServerError).JSON("error")
		}

		return c.Status(fiber.StatusOK).JSON("success")
	})

	app.Listen(":8081")
}
