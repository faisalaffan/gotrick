//go:build server

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// userServiceImpl adalah implementasi konkrit dari UserService.
// Di gRPC: generated server skeleton (interface) + user implementation.
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

// userServiceHTTP adalah transport layer untuk mengekspos service via HTTP.
// Di gRPC: ini HTTP/2 + protobuf marshaling, di-generate otomatis.
// Di simulasi ini: HTTP/1.1 + JSON marshaling, ditulis manual.
type userServiceHTTP struct {
	svc UserService
}

func (h *userServiceHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "hanya POST yang diizinkan", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Routing berdasarkan path (mirip /package.Service/Method di gRPC).
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
	fmt.Println("")
	fmt.Println("Coba jalankan client di terminal lain:")
	fmt.Println("  go run -tags=client ./networking/grpc_basics/")
	log.Fatal(http.ListenAndServe(addr, handler))
}
