package service

import (
	client "items-store-service/internal/client/service"
	"items-store-service/internal/models"
	"log"
)

func ItemListService() (*[]models.ItemListResponse, error) {
	response, err := client.GetAllItems()
	if err != nil {
		log.Fatal("Error get item list process")
		return nil, err
	}

	return response, nil
}
