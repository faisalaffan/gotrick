
package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	activeRequests sync.WaitGroup
	shuttingDown   atomic.Bool
)

// shutdownGuard menolak request baru saat shutting down.
func shutdownGuard(c *fiber.Ctx) error {
	if shuttingDown.Load() {
		return c.Status(fiber.StatusServiceUnavailable).SendString("server shutting down — coba lagi nanti")
	}
	activeRequests.Add(1)
	defer activeRequests.Done()
	return c.Next()
}

// slowHandler mensimulasikan long-running request (sleep 3 detik).
func slowHandler(c *fiber.Ctx) error {
	log.Println("[slow] Menerima request — mulai proses 3 detik")

	// Simulasi kerja berat 3 detik
	time.Sleep(3 * time.Second)

	log.Println("[slow] Request selesai")
	return c.SendString("Selesai setelah 3 detik")
}

// healthHandler untuk readiness probe.
func healthHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "GoTrick Graceful Shutdown Demo",
	})

	// Pasang guard sebelum semua route
	app.Use(shutdownGuard)

	app.Get("/slow", slowHandler)
	app.Get("/health", healthHandler)

	// Channel untuk menangkap sinyal OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Jalankan server di goroutine terpisah
	go func() {
		log.Println("Server Fiber started on :8080")
		log.Println("Cara test:")
		log.Println("  1. curl http://localhost:8080/slow &")
		log.Println("  2. sleep 1  (tunggu request mulai)")
		log.Println("  3. Ctrl+C    (kirim sinyal ke server)")
		log.Println("  → Server tunggu /slow selesai (3 detik) baru shutdown")
		if err := app.Listen(":8080"); err != nil {
			log.Printf("Server listen error: %v", err)
		}
	}()

	// Blokir sampai sinyal diterima
	sig := <-quit
	log.Printf("Menerima sinyal: %v. Memulai graceful shutdown...", sig)

	// Set flag agar tidak terima request baru
	shuttingDown.Store(true)
	log.Println("Menolak request baru, menunggu request aktif selesai...")

	// Tunggu semua request yang sedang berjalan selesai
	activeRequests.Wait()
	log.Println("Semua request aktif selesai.")

	// Sekarang aman untuk shutdown Fiber
	if err := app.Shutdown(); err != nil {
		log.Printf("Shutdown error: %v", err)
	}

	log.Println("Server shutdown selesai. Bye.")
}
