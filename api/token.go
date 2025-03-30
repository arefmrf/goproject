package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"snapshop/config"
	"snapshop/models"
)

func GetToken() (string, error) {
	url := fmt.Sprintf("%s/guest/v1/token?lat=%s&lng=%s", config.BaseURL, config.Latitude, config.Longitude)
	//url := "https://apix.snappshop.ir/guest/v1/token?lat=35.77331&lng=51.418591"
	//tokenURL := "https://apix.snapshop.ir/guest/v1/token?lat=35.77331&lng=51.418591"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("func GetToken error closing body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {

		body, _ := io.ReadAll(resp.Body) // Use io.ReadAll to read the response body
		return "", fmt.Errorf("error: status code: %v  response body: %v", resp.StatusCode, string(body))
	}

	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))

	if resp.ContentLength == 0 {
		fmt.Println("Error: Response body is empty")
		return "", fmt.Errorf("error: Response body is empty")
	}

	var response models.STokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("error: decoding JSON: %v  response body: %v", err, string(body))
	}
	return response.Data.Token, nil
}
