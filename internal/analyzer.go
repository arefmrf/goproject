package internal

import (
	"fmt"
	"snapshop/models"
	"sync"
)

func AnalyzeResponseWorker(results <-chan *models.MinimalResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		if result != nil {
			analyzeResponse(result)
		}
	}
}

func analyzeResponse(minimalResponse *models.MinimalResponse) {
	for _, item := range minimalResponse.Data.Structure[0].Items {
		fmt.Println(item.Title, item.Price.Discount, item.Price.DiscountedPrice)
	}
}
