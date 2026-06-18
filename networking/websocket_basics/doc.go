//go:build !server && !client

// Package websocket_basics adalah simulasi WebSocket tanpa library eksternal.
// Menunjukkan upgrade handshake, read/write loop, dan connection hijacking.
//
// Jalankan (2 terminal):
//
//	Terminal 1: go run -tags=server ./networking/websocket_basics/
//	Terminal 2: go run -tags=client ./networking/websocket_basics/
//
// File:
//   - server.go (build tag: server): WebSocket echo server
//   - client.go (build tag: client): WebSocket client
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: WebSocket Basics (Simulasi) ===")
	fmt.Println()
	fmt.Println("Jalankan dengan build tag:")
	fmt.Println("  Terminal 1: go run -tags=server ./networking/websocket_basics/")
	fmt.Println("  Terminal 2: go run -tags=client ./networking/websocket_basics/")
}
