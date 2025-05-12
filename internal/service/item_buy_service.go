package service

import (
	kafkaModel "items-store-service/internal/kafka/models"
	"items-store-service/internal/kafka/producer"
	"items-store-service/internal/models"
	"items-store-service/internal/utils"
	"log"
	"time"
)

func ItemBuyService(request models.ItemBuyRequest) error {

	//Suppose buy item success

	txnMessage := kafkaModel.TransactionMessage{
		TransactionId: utils.GenerateTransactionId(),
		UserId:        "",
		Title:         request.Title,
		Price:         request.Price,
		Categoty:      request.Categoty,
		CreateDate:    time.Now(),
	}

	err := producer.ProduceMessageTransaction(txnMessage)
	if err != nil {
		log.Fatal("Error produce message process")
		return err
	}

	return nil
}
