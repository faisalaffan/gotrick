// Simulasi pola gRPC tanpa protobuf + protoc.
//
// Di production:
//   - .proto file → protoc → generated Go code (struct, interface, client stub, server skeleton)
//   - Transport: HTTP/2 + protobuf (binary, schema-driven)
//   - Code generation menjamin type safety antara server dan client
//
// Simulasi ini:
//   - Define service interface manual (mirip generated interface dari .proto)
//   - Request/Response struct manual (mirip generated message dari .proto)
//   - Transport: HTTP/1.1 + JSON (mirip unary RPC, tanpa streaming)
//
// Install dependencies: none (standard library only).
// Cara run (dari root project):
//
//	Terminal 1 (server): go run -tags=server ./networking/grpc_basics/
//	Terminal 2 (client): go run -tags=client ./networking/grpc_basics/
package main

// User merepresentasikan data user.
// Di gRPC: message User { string id = 1; string name = 2; string email = 3; }
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// --- Request / Response structs ---
// Di gRPC: semua ini di-generate dari message di .proto.

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User User `json:"user"`
}

type ListUsersRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListUsersResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
}

// UserService adalah service interface.
// Di gRPC: di-generate dari service UserService { rpc GetUser(..) returns (..); }
// Pattern: server implement interface, client panggil via interface.
type UserService interface {
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
	ListUsers(req *ListUsersRequest) (*ListUsersResponse, error)
}
