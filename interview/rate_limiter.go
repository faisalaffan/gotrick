//go:build rate_limiter

package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter menerapkan algoritma Token Bucket.
// Token ditambahkan secara periodik berdasarkan rate (tokens/detik).
// Bucket memiliki kapasitas maksimum (burst).
type RateLimiter struct {
	rate      float64 // tokens per detik
	burst     float64 // kapasitas maksimum bucket
	tokens    float64 // token saat ini
	lastRefill time.Time
	mu        sync.Mutex
}

// NewRateLimiter membuat instance RateLimiter baru.
func NewRateLimiter(rate, burst float64) *RateLimiter {
	return &RateLimiter{
		rate:       rate,
		burst:      burst,
		tokens:     burst, // mulai penuh
		lastRefill: time.Now(),
	}
}

// Allow mengecek apakah request diizinkan.
// Satu token dikonsumsi per request. Jika token habis, request ditolak.
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Refill token berdasarkan waktu yang berlalu
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()
	rl.lastRefill = now

	// Tambah token, cap di burst
	rl.tokens += elapsed * rl.rate
	if rl.tokens > rl.burst {
		rl.tokens = rl.burst
	}

	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}
	return false
}

func main() {
	// Rate limiter: 5 request per detik, burst 5
	rl := NewRateLimiter(5, 5)

	fmt.Println("=== Rate Limiter Demo ===")
	fmt.Println("Rate: 5 request/detik, Burst: 5")
	fmt.Println("Mengirim 20 request berturut-turut...")
	fmt.Println()

	allowed := 0
	blocked := 0

	for i := range 20 {
		if rl.Allow() {
			allowed++
			fmt.Printf("[%02d] ALLOWED  ✅\n", i+1)
		} else {
			blocked++
			fmt.Printf("[%02d] BLOCKED  ❌\n", i+1)
		}
	}

	fmt.Printf("\nHasil: %d allowed, %d blocked (burst=5, sisanya kena rate limit)\n", allowed, blocked)
	fmt.Println()

	// Demonstrasi refill: tunggu 1 detik, token kembali
	fmt.Println("Tunggu 1 detik — token direfill...")
	time.Sleep(1 * time.Second)

	allowed2 := 0
	for i := range 5 {
		if rl.Allow() {
			allowed2++
			fmt.Printf("[refill %d] ALLOWED ✅\n", i+1)
		} else {
			fmt.Printf("[refill %d] BLOCKED ❌\n", i+1)
		}
	}
	fmt.Printf("\nSetelah refill 1 detik: %d allowed (seharusnya ~5)\n", allowed2)
}
