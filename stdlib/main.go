package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Standard Library & Gin ===")
	fmt.Println()
	fmt.Println("go run ./stdlib/httpserver/      # Gin HTTP server + middleware")
	fmt.Println("go run ./stdlib/httpclient/      # HTTP client + timeout")
	fmt.Println("go run ./stdlib/json/            # JSON marshal/unmarshal")
	fmt.Println("go run ./stdlib/context_usage/   # Context patterns")
	fmt.Println("go run ./stdlib/sync_primitives/ # Sync primitives")
	fmt.Println("go run ./stdlib/io_bufio/        # I/O & bufio")
	fmt.Println("go run ./stdlib/time_pkg/        # Time package")
}
