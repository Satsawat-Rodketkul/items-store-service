package service

import (
	"encoding/json"
	"items-store-service/internal/client/models"
	"items-store-service/internal/config"
	"log"
	"net/http"
)

func GetAllItems() ([]models.ItemsResponse, error) {
	url := config.GetValue("URL")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	response := []models.ItemsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return response, nil
}
