package kafka

import (
	"fmt"
	"items-store-service/internal/config"
	"log"

	"github.com/IBM/sarama"
)

var Producer sarama.SyncProducer

func KafkaConnection() {
	kafkaHost := config.GetValue("KAFKA_HOST")
	kafkaPort := config.GetValue("KAFKA_PORT")

	server := []string{fmt.Sprintf("%s:%s", kafkaHost, kafkaPort)}

	producer, err := sarama.NewSyncProducer(server, nil)
	if err != nil {
		log.Fatal("Failed to connect kafka server:", err)
	}
	Producer = producer

	log.Println("Connect to kafka server success")
}
