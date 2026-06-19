package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Networking ===")
	fmt.Println()
	fmt.Println("HTTP Retry:")
	fmt.Println("  go run ./networking/http_retry/ http://url-untuk-di-test")
	fmt.Println()
	fmt.Println("gRPC Simulation:")
	fmt.Println("  Terminal 1: go run ./networking/grpc_basics/server/")
	fmt.Println("  Terminal 2: go run ./networking/grpc_basics/client/")
	fmt.Println()
	fmt.Println("WebSocket Simulation:")
	fmt.Println("  Terminal 1: go run ./networking/websocket_basics/server/")
	fmt.Println("  Terminal 2: go run ./networking/websocket_basics/client/")
}
