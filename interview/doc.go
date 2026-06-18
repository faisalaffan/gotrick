//go:build !transfer && !rate_limiter && !graceful_shutdown && !retry_backoff && !idempotency

// Package interview berisi contoh-contoh soal interview Go untuk fintech:
// concurrent transfer, rate limiter, graceful shutdown, retry backoff, idempotency.
//
// Setiap file memiliki build tag sendiri. Jalankan dengan:
//
//	go run -tags=<tag> ./interview/
//
// Tag yang tersedia:
//
//	transfer          - Concurrent transfer + deadlock avoidance
//	rate_limiter      - Token bucket rate limiter
//	graceful_shutdown - HTTP server graceful shutdown
//	retry_backoff     - Exponential backoff + jitter
//	idempotency       - Idempotency key pattern
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Interview Prep (Fintech) ===")
	fmt.Println()
	fmt.Println("Gunakan build tag untuk menjalankan contoh spesifik:")
	fmt.Println()
	fmt.Println("  go run -tags=transfer          ./interview/   # Concurrent transfer")
	fmt.Println("  go run -tags=rate_limiter      ./interview/   # Token bucket")
	fmt.Println("  go run -tags=graceful_shutdown ./interview/   # Graceful shutdown")
	fmt.Println("  go run -tags=retry_backoff     ./interview/   # Retry + backoff")
	fmt.Println("  go run -tags=idempotency       ./interview/   # Idempotency key")
}
