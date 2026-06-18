// File: http_server.go
// HTTP server dengan custom mux, middleware logging, dan JSON response.

//go:build http_server

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// --- Middleware ---

// loggingMiddleware wrap http.Handler dengan logging request.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter untuk capture status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)

		log.Printf("[%s] %s %s | %d | %v",
			r.Method, r.URL.Path, r.RemoteAddr, lrw.statusCode, time.Since(start))
	})
}

// loggingResponseWriter mencatat status code response.
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// --- Handler ---

// apiResponse struktur standar JSON response.
type apiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

// jsonResponse helper kirim JSON.
func jsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("ERROR encode JSON: %v", err)
	}
}

func main() {
	// Buat custom mux
	mux := http.NewServeMux()

	// Route root
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, apiResponse{
			Status:  "ok",
			Message: "Selamat datang di Go HTTP server!",
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Route /hello
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Tamu"
		}
		jsonResponse(w, http.StatusOK, apiResponse{
			Status:  "ok",
			Message: fmt.Sprintf("Halo, %s!", name),
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Route /health
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, apiResponse{
			Status:  "healthy",
			Message: "Server berjalan normal.",
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Bungkus mux dengan logging middleware
	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server jalan di http://localhost:8080")
	fmt.Println("Coba test dengan curl:")
	fmt.Println("  curl http://localhost:8080/")
	fmt.Println("  curl http://localhost:8080/hello?name=Koding")
	fmt.Println("  curl http://localhost:8080/health")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server gagal: %v", err)
	}
}
