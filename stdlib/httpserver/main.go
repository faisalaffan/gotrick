// Server HTTP dengan Gin — routing, middleware, JSON response, dan route grouping.
// Gin menggunakan net/http di bawahnya, cocok untuk production fintech.

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// --- Middleware ---

// loggingMiddleware mencatat setiap request: method, path, status, durasi.
func loggingMiddleware(c *gin.Context) {
	start := time.Now()

	// Lanjutkan ke handler berikutnya
	c.Next()

	// Setelah handler selesai, log
	log.Printf("[%s] %s | %d | %v",
		c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
}

// --- Struct Response ---

type apiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

func main() {
	// Buat Gin router
	router := gin.Default()

	// Pasang middleware logging global
	router.Use(loggingMiddleware)

	// Route GET /
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, apiResponse{
			Status:  "ok",
			Message: "Selamat datang di Go Gin server!",
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Route GET /hello?name=...
	router.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Tamu")
		c.JSON(200, apiResponse{
			Status:  "ok",
			Message: fmt.Sprintf("Halo, %s!", name),
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Route GET /health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, apiResponse{
			Status:  "healthy",
			Message: "Server berjalan normal.",
			Time:    time.Now().Format(time.RFC3339),
		})
	})

	// Group route dengan prefix /api
	api := router.Group("/api")
	api.GET("/karyawan", func(c *gin.Context) {
		karyawan := []gin.H{
			{"id": "1", "nama": "Budi"},
			{"id": "2", "nama": "Siti"},
		}
		c.JSON(200, gin.H{
			"status": "ok",
			"data":   karyawan,
		})
	})

	fmt.Println("Server Gin jalan di http://localhost:8080")
	fmt.Println("Coba test dengan curl:")
	fmt.Println("  curl http://localhost:8080/")
	fmt.Println("  curl http://localhost:8080/hello?name=Koding")
	fmt.Println("  curl http://localhost:8080/health")
	fmt.Println("  curl http://localhost:8080/api/karyawan")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server gagal: %v", err)
	}
}
