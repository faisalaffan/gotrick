//go:build !http_retry

// Package networking berisi contoh HTTP retry, gRPC, dan WebSocket.
//
// HTTP Retry:
//
//	go run -tags=http_retry ./networking/ http://url-untuk-di-test
//
// gRPC basics (simulasi tanpa protoc):
//
//	Terminal 1: go run -tags=server ./networking/grpc_basics/
//	Terminal 2: go run -tags=client ./networking/grpc_basics/
//
// WebSocket basics (simulasi tanpa library eksternal):
//
//	Terminal 1: go run -tags=server ./networking/websocket_basics/
//	Terminal 2: go run -tags=client ./networking/websocket_basics/
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Networking ===")
	fmt.Println()
	fmt.Println("HTTP Retry:")
	fmt.Println("  go run -tags=http_retry ./networking/ http://url-untuk-di-test")
	fmt.Println()
	fmt.Println("gRPC Simulation:")
	fmt.Println("  Terminal 1: go run -tags=server ./networking/grpc_basics/")
	fmt.Println("  Terminal 2: go run -tags=client ./networking/grpc_basics/")
	fmt.Println()
	fmt.Println("WebSocket Simulation:")
	fmt.Println("  Terminal 1: go run -tags=server ./networking/websocket_basics/")
	fmt.Println("  Terminal 2: go run -tags=client ./networking/websocket_basics/")
}
