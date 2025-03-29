package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/andybalholm/brotli"
)

const worker = 8

func getToken() string {
	tokenURL := "https://apix.snappshop.ir/guest/v1/token?lat=35.77331&lng=51.418591"
	resp, err := http.Get(tokenURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)            // Use io.ReadAll to read the response body
		fmt.Println("Response body:", string(body)) // Print the body to understand the error
		return ""
	}

	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))

	if resp.ContentLength == 0 {
		fmt.Println("Error: Response body is empty")
		return ""
	}

	var response STokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Error decoding JSON:", err)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("Response body:", string(body))
		return ""
	}
	return response.Data.Token
}

func fetchData(url string, token string, requestBody map[string]interface{}) ([]byte, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Host", "apix.snappshop.ir")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:136.0) Gecko/20100101 Firefox/136.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Referer", "https://snappshop.ir/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("s-device-source", "shop")
	req.Header.Set("uuid", "fe14f497-60a2-4497-8cbf-7e3572a31cca")
	req.Header.Set("S-Session-Id", "5bf16222-aa54-4530-a54c-79905c07c618")
	req.Header.Set("s-device", "DESKTOP")
	req.Header.Set("X-Origin", "https://snappshop.ir")
	req.Header.Set("Origin", "https://snappshop.ir")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Priority", "u=4")
	req.Header.Set("TE", "trailers")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: status code %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var reader io.Reader
	switch resp.Header.Get("Content-Encoding") {
	case "br":
		reader = brotli.NewReader(resp.Body)
	case "gzip":
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error creating gzip reader: %v", err)
		}
		defer func(gzipReader *gzip.Reader) {
			err := gzipReader.Close()
			if err != nil {
				fmt.Println("error closing gzip reader:", err)
			}
		}(gzipReader)
		reader = gzipReader
	default:
		reader = resp.Body
	}

	responseBody, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return responseBody, nil
}

func initList(token string) (*MinimalResponse, error) {
	url := "https://apix.snappshop.ir/search/v1?lat=35.77331&lng=51.418591"

	requestBody := map[string]interface{}{
		"slug":   "gwWRMg",
		"render": 3,
	}

	responseBody, err := fetchData(url, token, requestBody)
	if err != nil {
		return nil, err
	}

	var minimalResponse MinimalResponse
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

func getList(token string, uuid string, skip int, wg *sync.WaitGroup, results chan<- *MinimalResponse) {
	defer wg.Done()

	url := "https://apix.snappshop.ir/search/v1?lat=35.77331&lng=51.418591"
	requestBody := map[string]interface{}{
		"slug":   "gwWRMg",
		"render": 3,
		"uuid":   uuid,
		"skip":   skip,
	}

	responseBody, err := fetchData(url, token, requestBody)
	if err != nil {
		fmt.Println("Fetch error:", err)
		return
	}

	var minimalResponse MinimalResponse
	if err := json.Unmarshal(responseBody, &minimalResponse); err != nil {
		fmt.Println("JSON decode error:", err)
		fmt.Println("Raw response body:", string(responseBody))
		return
	}

	results <- &minimalResponse
}

func analyzeResponseWorker(results <-chan *MinimalResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		if result != nil {
			analyzeResponse(result)
		}
	}
}

func analyzeResponse(minimalResponse *MinimalResponse) {
	for _, item := range minimalResponse.Data.Structure[0].Items {
		fmt.Println(item.Title, item.Price.Discount, item.Price.DiscountedPrice)
	}
}

func main() {
	token := getToken()
	if token == "" {
		fmt.Println("Failed to retrieve token. Exiting...")
		return
	}

	results := make(chan *MinimalResponse, worker)
	var fetchWg sync.WaitGroup
	var analyzeWg sync.WaitGroup

	for i := 0; i < worker; i++ {
		analyzeWg.Add(1)
		go analyzeResponseWorker(results, &analyzeWg)
	}

	initResponse, err := initList(token)
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
		go getList(token, uuid, page, &fetchWg, results)
	}

	fetchWg.Wait()
	close(results)
	analyzeWg.Wait()
}
