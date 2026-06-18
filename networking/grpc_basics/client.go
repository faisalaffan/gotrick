//go:build client

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// userServiceClient adalah client stub.
// Di gRPC: di-generate oleh protoc, method-method panggil server via HTTP/2.
//
// Client ini mengimplementasi interface UserService yang SAMA dengan server.
// Caller tidak perlu tahu apakah service lokal atau remote — inilah essence dari gRPC.
type userServiceClient struct {
	baseURL string
	client  *http.Client
}

func newUserServiceClient(addr string) *userServiceClient {
	return &userServiceClient{
		baseURL: fmt.Sprintf("http://%s", addr),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *userServiceClient) GetUser(req *GetUserRequest) (*GetUserResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetUser marshal: %w", err)
	}

	resp, err := c.client.Post(c.baseURL+"/UserService/GetUser", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("GetUser: %w", err)
	}
	defer resp.Body.Close()

	var result GetUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("GetUser decode: %w", err)
	}
	return &result, nil
}

func (c *userServiceClient) ListUsers(req *ListUsersRequest) (*ListUsersResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("ListUsers marshal: %w", err)
	}

	resp, err := c.client.Post(c.baseURL+"/UserService/ListUsers", "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("ListUsers: %w", err)
	}
	defer resp.Body.Close()

	var result ListUsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("ListUsers decode: %w", err)
	}
	return &result, nil
}

func main() {
	addr := ":50051"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	client := newUserServiceClient(addr)

	fmt.Printf("Menghubungi server di %s\n\n", addr)

	// Panggil GetUser via interface yang sama.
	userResp, err := client.GetUser(&GetUserRequest{ID: "42"})
	if err != nil {
		fmt.Printf("Error GetUser: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("GetUser response: ID=%s Name=%s Email=%s\n",
		userResp.User.ID, userResp.User.Name, userResp.User.Email)

	// Panggil ListUsers.
	listResp, err := client.ListUsers(&ListUsersRequest{Page: 1, Limit: 10})
	if err != nil {
		fmt.Printf("Error ListUsers: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("ListUsers response: %d user(s)\n", listResp.Total)
	for _, u := range listResp.Users {
		fmt.Printf("  - %s <%s>\n", u.Name, u.Email)
	}

	fmt.Println("\nSukses! Semua RPC call berhasil.")
}
