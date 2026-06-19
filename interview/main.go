package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Interview Prep (Fintech) ===")
	fmt.Println()
	fmt.Println("go run ./interview/transfer/          # Concurrent transfer")
	fmt.Println("go run ./interview/rate_limiter/      # Token bucket")
	fmt.Println("go run ./interview/graceful_shutdown/ # Graceful shutdown")
	fmt.Println("go run ./interview/retry_backoff/     # Retry + backoff")
	fmt.Println("go run ./interview/idempotency/       # Idempotency key")
}
