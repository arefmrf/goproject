package services

//import (
//	"fmt"
//	"snapshop/api"
//	"snapshop/models"
//)
//
//func SearchProducts(token string, slug string) (*models.MinimalResponse, error) {
//	request := models.RequestBody{
//		Slug:   slug,
//		Render: 3,
//	}
//
//	response, err := api.FetchSearchResults(token, request)
//	if err != nil {
//		return nil, fmt.Errorf("search failed: %v", err)
//	}
//
//	return response, nil
//}
