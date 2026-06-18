//go:build graceful_shutdown

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// slowHandler mensimulasikan long-running request (sleep 3 detik).
func slowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[slow] Menerima request — mulai proses 3 detik")
	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "Selesai setelah 3 detik")
		log.Println("[slow] Request selesai")
	case <-r.Context().Done():
		// Client disconnect — lebih cepat cleanup
		log.Println("[slow] Request dibatalkan (client disconnect)")
		http.Error(w, "request cancelled", http.StatusRequestTimeout)
	}
}

// healthHandler untuk readiness probe.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/slow", slowHandler)
	mux.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Channel untuk menangkap sinyal OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Jalankan server di goroutine terpisah
	go func() {
		log.Println("Server started on :8080")
		log.Println("Cara test: curl http://localhost:8080/slow lalu Ctrl+C")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Blokir sampai sinyal diterima
	sig := <-quit
	log.Printf("Menerima sinyal: %v. Memulai shutdown...", sig)

	// Beri waktu 30 detik untuk request yang sedang berlangsung
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v. Force close...", err)
		// Jika timeout, panggil Close() untuk force stop
		if closeErr := server.Close(); closeErr != nil {
			log.Printf("Close error: %v", closeErr)
		}
	}

	log.Println("Server shutdown selesai. Bye.")
}
