package service

import (
	"encoding/json"
	"items-store-service/internal/config"
	"items-store-service/internal/models"
	"log"
	"net/http"
)

func GetAllItems() (*[]models.ItemListResponse, error) {
	url := config.GetValue("URL")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error get item from webclient")
		return nil, err
	}
	defer resp.Body.Close()

	response := &[]models.ItemListResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal("Error map response from webclient")
		return nil, err
	}
	return response, nil
}
