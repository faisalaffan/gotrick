// File: http_client.go
// HTTP client: GET dengan DefaultClient, custom timeout, RequestWithContext,
// baca body, handle non-200.


package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// 1. HTTP GET pakai DefaultClient
	fmt.Println("=== 1. GET dengan http.DefaultClient ===")
	getWithDefaultClient("https://httpbin.org/get")

	// 2. Custom client dengan timeout
	fmt.Println("\n=== 2. Custom client dengan timeout 5s ===")
	customClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
	}
	getWithCustomClient(customClient, "https://httpbin.org/get")

	// 3. Request dengan context + timeout
	fmt.Println("\n=== 3. Request dengan context timeout 2s ===")
	getWithContext("https://httpbin.org/delay/1")

	// 4. Request yang akan timeout (delay 4s > timeout 2s)
	fmt.Println("\n=== 4. Request yang timeout ===")
	getWithContext("https://httpbin.org/delay/4")

	// 5. Handle non-200 status code
	fmt.Println("\n=== 5. Handle non-200 (404) ===")
	getWithDefaultClient("https://httpbin.org/status/404")

	// 6. Handle non-200 (500)
	fmt.Println("\n=== 6. Handle non-200 (500) ===")
	getWithDefaultClient("https://httpbin.org/status/500")
}

// getWithDefaultClient GET dengan DefaultClient dan handle response.
func getWithDefaultClient(url string) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Handle non-200
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Status: %s\nBody: %s\n", resp.Status, string(body))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ERROR baca body: %v\n", err)
		return
	}
	fmt.Printf("Status: %s\nBody (len=%d): %s...\n", resp.Status, len(body), string(body[:min(len(body), 200)]))
}

// getWithCustomClient GET dengan custom HTTP client.
func getWithCustomClient(client *http.Client, url string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("ERROR buat request: %v\n", err)
		return
	}
	req.Header.Set("User-Agent", "gotrick/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nHeaders: %v\nBody: %s...\n",
		resp.Status, resp.Header, string(body[:min(len(body), 150)]))
}

// getWithContext GET dengan context timeout.
func getWithContext(url string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("ERROR buat request: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %s\nBody: %s...\n",
		resp.Status, string(body[:min(len(body), 150)]))
}
