package api

import (
	"encoding/json"
	"fmt"
	"snapshop/config"
	"snapshop/internal"
	"snapshop/models"
	"sync"
)

func InitList(token string) (*models.MinimalResponse, error) {
	url := fmt.Sprintf("%s/search/v1?lat=%s&lng=%s", config.BaseURL, config.Latitude, config.Longitude)

	requestBody := models.RequestBody{
		Slug:   "gwWRMg",
		Render: 3,
	}

	responseBody, err := internal.FetchData(url, token, requestBody)
	if err != nil {
		return nil, err
	}

	var minimalResponse models.MinimalResponse
	if err := json.Unmarshal(responseBody, &minimalResponse); err != nil {
		fmt.Println("Error decoding JSON:", err)
		fmt.Println("Raw response body:", string(responseBody))
		return nil, err
	}

	if len(minimalResponse.Data.Structure) == 0 {
		return nil, fmt.Errorf("empty response structure")
	}

	return &minimalResponse, nil
}

func GetList(token string, uuid string, skip int, wg *sync.WaitGroup, results chan<- *models.MinimalResponse) {
	defer wg.Done()

	url := fmt.Sprintf("%s/search/v1?lat=%s&lng=%s", config.BaseURL, config.Latitude, config.Longitude)
	requestBody := models.RequestBody{
		Slug:   "gwWRMg",
		Render: 3,
		UUID:   uuid,
		Skip:   skip,
	}

	responseBody, err := internal.FetchData(url, token, requestBody)
	if err != nil {
		fmt.Println("Fetch error:", err)
		return
	}

	var minimalResponse models.MinimalResponse
	if err := json.Unmarshal(responseBody, &minimalResponse); err != nil {
		fmt.Println("JSON decode error:", err)
		fmt.Println("Raw response body:", string(responseBody))
		return
	}

	results <- &minimalResponse
}
