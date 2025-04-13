package internal

import (
	"fmt"
	"snapshop/models"
	"sync"
	"time"
)

func AnalyzeResponseWorker(results <-chan *models.MinimalResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		if result != nil {
			analyzeResponse(result)
			addDB(result)
		}
	}
}

func analyzeResponse(minimalResponse *models.MinimalResponse) {
	for _, item := range minimalResponse.Data.Structure[0].Items {
		fmt.Println(item.Price.Discount, item.Price.DiscountedPrice, item.Title)
	}
}

func addDB(minimalResponse *models.MinimalResponse) {
	db, err := DBSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	var products []models.DBProduct
	for _, item := range minimalResponse.Data.Structure[0].Items {
		endAt, err := time.Parse(models.TimeLayout, item.Price.EndAt)
		if err != nil {
			fmt.Println("Failed to parse EndAt:", err)
			continue // Skip this item if parsing fails
		}
		products = append(products, models.DBProduct{
			Title:           item.Title,
			Discount:        item.Price.Discount,
			Price:           item.Price.Price,
			DiscountedPrice: item.Price.DiscountedPrice,
			EndAt:           endAt,
			ProductID:       item.ID,
		})
	}
	if len(products) > 0 {
		if err := db.Create(&products).Error; err != nil {
			fmt.Println("Failed to insert products:", err)
		}
	} else {
		fmt.Println("*** len(products) < 0 ***")
	}
}
