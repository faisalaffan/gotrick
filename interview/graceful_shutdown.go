//go:build graceful_shutdown

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

// slowHandler mensimulasikan long-running request (sleep 3 detik).
func slowHandler(c *fiber.Ctx) error {
	log.Println("[slow] Menerima request — mulai proses 3 detik")

	// Simulasi kerja 3 detik
	select {
	case <-time.After(3 * time.Second):
		log.Println("[slow] Request selesai")
		return c.SendString("Selesai setelah 3 detik")
	case <-c.Context().Done():
		log.Println("[slow] Request dibatalkan (client disconnect)")
		return c.Status(fiber.StatusRequestTimeout).SendString("request cancelled")
	}
}

// healthHandler untuk readiness probe.
func healthHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "GoTrick Graceful Shutdown Demo",
	})

	app.Get("/slow", slowHandler)
	app.Get("/health", healthHandler)

	// Channel untuk menangkap sinyal OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Jalankan server di goroutine terpisah
	go func() {
		log.Println("Server Fiber started on :8080")
		log.Println("Cara test:")
		log.Println("  1. curl http://localhost:8080/slow")
		log.Println("  2. Langsung Ctrl+C — server tunggu /slow selesai")
		if err := app.Listen(":8080"); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	// Blokir sampai sinyal diterima
	sig := <-quit
	log.Printf("Menerima sinyal: %v. Memulai graceful shutdown...", sig)

	// Fiber ShutdownWithTimeout: tunggu request ongoing selesai, maksimal 30 detik
	if err := app.ShutdownWithTimeout(30 * time.Second); err != nil {
		log.Printf("Shutdown error: %v", err)
	}

	log.Println("Server shutdown selesai. Bye.")
}
