package services

import (
	"fmt"
	"snapshop/api"
)

func GetAuthToken() (string, error) {
	token, err := api.GetToken()
	if err != nil || token == "" {
		return "", fmt.Errorf("failed to retrieve token: %v", err)
	}
	return token, nil
}
