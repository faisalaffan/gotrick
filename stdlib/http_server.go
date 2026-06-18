// File: http_server.go
// HTTP server dengan Fiber — routing, middleware, dan JSON response.

//go:build http_server

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// --- Middleware ---

// loggingMiddleware mencatat setiap request: method, path, status, durasi.
func loggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Lanjutkan ke handler berikutnya
	err := c.Next()

	// Setelah handler selesai, log
	log.Printf("[%s] %s | %d | %v",
		c.Method(), c.Path(), c.Response().StatusCode(), time.Since(start))

	return err
}

// --- Struct Response ---

type apiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	// Buat Fiber app
	app := fiber.New(fiber.Config{
		AppName: "GoTrick HTTP Server",
	})

	// Pasang middleware logging global
	app.Use(loggingMiddleware)

	// Route GET /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(apiResponse{
			Status:  "ok",
			Message: "Selamat datang di Go Fiber server!",
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Route GET /hello?name=...
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Tamu")
		return c.JSON(apiResponse{
			Status:  "ok",
			Message: fmt.Sprintf("Halo, %s!", name),
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Route GET /health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(apiResponse{
			Status:  "healthy",
			Message: "Server berjalan normal.",
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Group route dengan prefix /api
	api := app.Group("/api")
	api.Get("/users", func(c *fiber.Ctx) error {
		users := []map[string]string{
			{"id": "1", "name": "Budi"},
			{"id": "2", "name": "Siti"},
		}
		return c.JSON(fiber.Map{
			"status": "ok",
			"data":   users,
		})
	})

	fmt.Println("Server Fiber jalan di http://localhost:8080")
	fmt.Println("Coba test dengan curl:")
	fmt.Println("  curl http://localhost:8080/")
	fmt.Println("  curl http://localhost:8080/hello?name=Koding")
	fmt.Println("  curl http://localhost:8080/health")
	fmt.Println("  curl http://localhost:8080/api/users")

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Server gagal: %v", err)
	}
}
