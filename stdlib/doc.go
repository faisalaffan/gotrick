//go:build !http_server && !http_client && !json && !context_usage && !sync_primitives && !io_bufio && !time_pkg

// Package stdlib berisi contoh-contoh penggunaan Go standard library dan Fiber:
// Fiber HTTP framework, encoding/json, context, sync, io, bufio, time.
//
// Setiap file memiliki build tag sendiri. Jalankan dengan:
//
//	go run -tags=<tag> ./stdlib/
//
// Tag yang tersedia:
//
//	http_server     - HTTP server + middleware pattern
//	http_client     - HTTP client dengan timeout
//	json            - Marshal/Unmarshal, struct tags, custom marshaler
//	context_usage   - Context: cancel, timeout, deadline
//	sync_primitives - Mutex, RWMutex, WaitGroup, Once, sync.Map
//	io_bufio        - io.Reader/Writer, bufio.Scanner
//	time_pkg        - Time formatting, tickers, timers
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Standard Library ===")
	fmt.Println()
	fmt.Println("Gunakan build tag untuk menjalankan contoh spesifik:")
	fmt.Println()
	fmt.Println("  go run -tags=http_server     ./stdlib/   # Fiber HTTP server + middleware")
	fmt.Println("  go run -tags=http_client     ./stdlib/   # HTTP client + timeout")
	fmt.Println("  go run -tags=json            ./stdlib/   # JSON marshal/unmarshal")
	fmt.Println("  go run -tags=context_usage   ./stdlib/   # Context patterns")
	fmt.Println("  go run -tags=sync_primitives ./stdlib/   # Sync primitives")
	fmt.Println("  go run -tags=io_bufio        ./stdlib/   # I/O & bufio")
	fmt.Println("  go run -tags=time_pkg        ./stdlib/   # Time package")
}
