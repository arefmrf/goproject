package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"snapshop/models"
)

func FetchData(url string, token string, requestBody models.RequestBody) ([]byte, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	//req.Header.Set("Host", "apix.snapshop.ir")
	//req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:136.0) Gecko/20100101 Firefox/136.0")
	//req.Header.Set("Accept", "*/*")
	//req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	//req.Header.Set("Referer", "https://snapshop.ir/")
	//req.Header.Set("s-device-source", "shop")
	//req.Header.Set("uuid", "fe14f497-60a2-4497-8cbf-7e3572a31cca")
	//req.Header.Set("S-Session-Id", "5bf16222-aa54-4530-a54c-79905c07c618")
	//req.Header.Set("s-device", "DESKTOP")
	//req.Header.Set("X-Origin", "https://snapshop.ir")
	//req.Header.Set("Origin", "https://snapshop.ir")
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Set("Sec-Fetch-Dest", "empty")
	//req.Header.Set("Sec-Fetch-Mode", "cors")
	//req.Header.Set("Sec-Fetch-Site", "same-site")
	//req.Header.Set("Priority", "u=4")
	//req.Header.Set("TE", "trailers")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making POST request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error closing body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: status code %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var reader io.Reader
	reader = resp.Body
	//switch resp.Header.Get("Content-Encoding") {
	////req.Header.Set("Accept-Encoding", "gzip, br, deflate, zstd")
	//case "br":
	//
	//	//"github.com/andybalholm/brotli"
	//	reader = brotli.NewReader(resp.Body)
	//case "gzip":
	//	gzipReader, err := gzip.NewReader(resp.Body)
	//	if err != nil {
	//		return nil, fmt.Errorf("error creating gzip reader: %v", err)
	//	}
	//	defer func(gzipReader *gzip.Reader) {
	//		err := gzipReader.Close()
	//		if err != nil {
	//			fmt.Println("error closing gzip reader:", err)
	//		}
	//	}(gzipReader)
	//	reader = gzipReader
	//default:
	//	reader = resp.Body
	//}

	responseBody, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return responseBody, nil
}
