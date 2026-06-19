// Simulasi gRPC Server — implementasi UserService via HTTP/1.1 + JSON.
// Di production: .proto → protoc → generated server skeleton, HTTP/2 + protobuf.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// --- Shared types (di production: generated dari .proto) ---

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

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

type UserService interface {
	GetUser(req *GetUserRequest) (*GetUserResponse, error)
	ListUsers(req *ListUsersRequest) (*ListUsersResponse, error)
}

// --- Server implementation ---

type userServiceImpl struct{}

func (s *userServiceImpl) GetUser(req *GetUserRequest) (*GetUserResponse, error) {
	return &GetUserResponse{
		User: User{
			ID:    req.ID,
			Name:  "Budi Santoso",
			Email: "budi@example.com",
		},
	}, nil
}

func (s *userServiceImpl) ListUsers(req *ListUsersRequest) (*ListUsersResponse, error) {
	users := []User{
		{ID: "1", Name: "Budi Santoso", Email: "budi@example.com"},
		{ID: "2", Name: "Siti Rahayu", Email: "siti@example.com"},
	}
	return &ListUsersResponse{Users: users, Total: len(users)}, nil
}

// Transport layer (simulasi gRPC HTTP/2 handler)
type userServiceHTTP struct {
	svc UserService
}

func (h *userServiceHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "hanya POST yang diizinkan", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	switch r.URL.Path {
	case "/UserService/GetUser":
		var req GetUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
			return
		}
		resp, err := h.svc.GetUser(&req)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)

	case "/UserService/ListUsers":
		var req ListUsersRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
			return
		}
		resp, err := h.svc.ListUsers(&req)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err.Error()), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, `{"error":"unknown method"}`, http.StatusNotFound)
	}
}

func main() {
	svc := &userServiceImpl{}
	handler := &userServiceHTTP{svc: svc}

	addr := ":50051"
	fmt.Printf("gRPC-like server listening di %s\n", addr)
	fmt.Println("Endpoint:")
	fmt.Println("  POST /UserService/GetUser")
	fmt.Println("  POST /UserService/ListUsers")
	fmt.Println()
	fmt.Println("Jalankan client di terminal lain:")
	fmt.Println("  go run ./networking/grpc_basics/client/")
	log.Fatal(http.ListenAndServe(addr, handler))
}
