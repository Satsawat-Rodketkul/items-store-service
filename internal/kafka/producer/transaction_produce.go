package producer

import (
	"encoding/json"
	"items-store-service/internal/config"
	"items-store-service/internal/kafka"
	"items-store-service/internal/kafka/models"
	"log"

	"github.com/IBM/sarama"
)

func ProduceMessageTransaction(message models.TransactionMessage) error {
	topic := config.GetValue("KAFKA_TOPIC")
	value, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Error marshal message")
		return err
	}

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := kafka.Producer.SendMessage(&msg)
	if err != nil {
		log.Fatal("Error produce message")
		return err
	}

	log.Printf("partition: %v, offset: %v \n", partition, offset)

	defer kafka.Producer.Close()

	log.Println("produce success")

	return nil
}
