package main

import (
	"fmt"
	"snapshop/api"
	"snapshop/config"
	"snapshop/internal"
	"snapshop/models"
	"snapshop/services"
	"sync"
)

func main() {
	token, err := services.GetAuthToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	results := make(chan *models.MinimalResponse, config.Worker)
	var fetchWg sync.WaitGroup
	var analyzeWg sync.WaitGroup

	for i := 0; i < config.Worker; i++ {
		analyzeWg.Add(1)
		go internal.AnalyzeResponseWorker(results, &analyzeWg)
	}

	initResponse, err := api.InitList(token)
	if err != nil {
		fmt.Println("Failed to fetch initial list:", err)
		close(results)
		analyzeWg.Wait()
		return
	}

	fmt.Println("Status:", initResponse.Status)
	results <- initResponse

	totalPages := initResponse.Data.Structure[0].Pagination.TotalPages
	uuid := initResponse.Data.Structure[0].UUID

	for page := 1; page < totalPages; page++ {
		fetchWg.Add(1)
		go api.GetList(token, uuid, page, &fetchWg, results)
	}

	fetchWg.Wait()
	close(results)
	analyzeWg.Wait()
}
