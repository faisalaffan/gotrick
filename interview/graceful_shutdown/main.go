// Graceful shutdown dengan Gin + net/http.Server.
// Gin pakai net/http di bawahnya — graceful shutdown via server.Shutdown().
// Pattern ini cocok untuk production fintech.

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	activeRequests sync.WaitGroup
	shuttingDown   atomic.Bool
)

// shutdownGuard menolak request baru saat shutting down.
func shutdownGuard(c *gin.Context) {
	if shuttingDown.Load() {
		c.String(http.StatusServiceUnavailable, "server shutting down — coba lagi nanti")
		c.Abort()
		return
	}
	activeRequests.Add(1)
	defer activeRequests.Done()
	c.Next()
}

// slowHandler mensimulasikan long-running request (sleep 3 detik).
func slowHandler(c *gin.Context) {
	log.Println("[slow] Menerima request — mulai proses 3 detik")

	// Simulasi kerja berat 3 detik — respect context cancellation
	select {
	case <-time.After(3 * time.Second):
		log.Println("[slow] Request selesai")
		c.String(200, "Selesai setelah 3 detik")
	case <-c.Request.Context().Done():
		log.Println("[slow] Request dibatalkan (client disconnect)")
		c.String(499, "request cancelled")
	}
}

// healthHandler untuk readiness probe.
func healthHandler(c *gin.Context) {
	c.String(200, "OK")
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	// Pasang guard sebelum semua route
	router.Use(shutdownGuard)

	router.GET("/slow", slowHandler)
	router.GET("/health", healthHandler)

	// Bungkus Gin router dalam http.Server agar bisa graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Channel untuk menangkap sinyal OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Jalankan server di goroutine
	go func() {
		log.Println("Server Gin started on :8080")
		log.Println("Cara test:")
		log.Println("  1. curl http://localhost:8080/slow &")
		log.Println("  2. sleep 1")
		log.Println("  3. Ctrl+C → server tunggu /slow selesai baru shutdown")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
		}
	}()

	// Blokir sampai sinyal diterima
	sig := <-quit
	log.Printf("Menerima sinyal: %v. Memulai graceful shutdown...", sig)

	// Set flag — tolak request baru
	shuttingDown.Store(true)
	log.Println("Menolak request baru, menunggu request aktif selesai...")

	// Tunggu semua request yang sedang berjalan
	activeRequests.Wait()
	log.Println("Semua request aktif selesai.")

	// Shutdown server dengan timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}

	log.Println("Server shutdown selesai. Bye.")
}
