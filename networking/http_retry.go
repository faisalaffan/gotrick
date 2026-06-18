// HTTP Client dengan Retry Logic
//
// Exponential backoff: 1s, 2s, 4s -- max 3 retry.
// Jitter +-20% untuk menghindari thundering herd.
// Retry hanya untuk 5xx dan network error.
//
// Cara run:
//
//	go run ./networking/http_retry.go <url>
//
// Contoh URL:
//
//	go run ./networking/http_retry.go https://httpbin.org/status/500
//	go run ./networking/http_retry.go https://httpbin.org/delay/10
//	go run ./networking/http_retry.go http://localhost:9999  (connection refused)
//go:build http_retry

package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"time"
)

const (
	maxRetries  = 3
	baseBackoff = 1 * time.Second
	maxBackoff  = 8 * time.Second
	// totalTimeout adalah batas waktu keseluruhan untuk semua percobaan.
	totalTimeout = 30 * time.Second
	// jitterFactor = +-20% dari backoff.
	jitterFactor = 0.2
)

func retryableHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

// isRetryable mengembalikan true jika layak di-retry.
// Hanya 5xx (server error) dan network error yang di-retry.
// 4xx (client error) dan success (2xx) tidak di-retry.
func isRetryable(resp *http.Response, err error) bool {
	if err != nil {
		return true
	}
	return resp.StatusCode >= 500
}

// calculateBackoff menghitung delay exponensial dengan jitter.
// Formula: min(base * 2^attempt, maxBackoff) + jitter(+-20%).
func calculateBackoff(attempt int) time.Duration {
	backoff := baseBackoff * (1 << attempt) // 1s, 2s, 4s
	if backoff > maxBackoff {
		backoff = maxBackoff
	}
	// Jitter +-20%: thundering herd mitigation.
	jitter := time.Duration(float64(backoff) * jitterFactor * (2*rand.Float64() - 1))
	return backoff + jitter
}

// doRequestWithRetry menjalankan HTTP request dengan retry logic.
// Context digunakan untuk timeout keseluruhan.
func doRequestWithRetry(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	var lastErr error
	var lastResp *http.Response

	for attempt := range maxRetries + 1 {
		// Cek konteks sebelum tiap attempt.
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("konteks berakhir sebelum attempt %d: %w", attempt+1, ctx.Err())
		default:
		}

		resp, err := client.Do(req)

		if err != nil {
			// Network error: timeout, DNS failure, connection refused, dll.
			lastErr = err
			if resp != nil {
				resp.Body.Close()
			}
		} else if resp.StatusCode >= 500 {
			// 5xx: server error, retryable.
			lastResp = resp
			resp.Body.Close()
		} else {
			// Success atau 4xx (client error) -- jangan retry.
			return resp, nil
		}

		if attempt < maxRetries {
			backoff := calculateBackoff(attempt)
			status := "network error"
			if lastErr == nil && lastResp != nil {
				status = lastResp.Status
			}
			fmt.Printf("Attempt %d gagal (%s), retry dalam %v...\n", attempt+1, status, backoff)

			select {
			case <-ctx.Done():
				return nil, fmt.Errorf("konteks berakhir saat backoff: %w", ctx.Err())
			case <-time.After(backoff):
			}
		}
	}

	// Semua attempt habis.
	if lastErr != nil {
		return nil, fmt.Errorf("gagal setelah %d percobaan: %w", maxRetries+1, lastErr)
	}
	return lastResp, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run http_retry.go <url>")
		fmt.Println("Contoh:")
		fmt.Println("  go run ./networking/http_retry.go https://httpbin.org/status/500")
		fmt.Println("  go run ./networking/http_retry.go https://httpbin.org/delay/10")
		fmt.Println("  go run ./networking/http_retry.go http://localhost:9999")
		os.Exit(1)
	}

	url := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), totalTimeout)
	defer cancel()

	client := retryableHTTPClient()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("Gagal membuat request: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Memanggil %s ...\n", url)

	resp, err := doRequestWithRetry(ctx, client, req)
	if err != nil {
		fmt.Printf("Request gagal: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("Sukses! Status: %s\n", resp.Status)
}
