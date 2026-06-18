//go:build !server && !client

// Package grpc_basics adalah simulasi pola gRPC tanpa protobuf+protoc.
// Menunjukkan service interface contract, server implementation, dan client stub.
//
// Jalankan (2 terminal):
//
//	Terminal 1: go run -tags=server ./networking/grpc_basics/
//	Terminal 2: go run -tags=client ./networking/grpc_basics/
//
// File:
//   - service.go: type User, interface UserService, request/response struct
//   - server.go (build tag: server): implementasi service + HTTP handler
//   - client.go (build tag: client): client stub via HTTP
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: gRPC Basics (Simulasi) ===")
	fmt.Println()
	fmt.Println("Jalankan dengan build tag:")
	fmt.Println("  Terminal 1: go run -tags=server ./networking/grpc_basics/")
	fmt.Println("  Terminal 2: go run -tags=client ./networking/grpc_basics/")
}
