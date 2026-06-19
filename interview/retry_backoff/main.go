
package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const maxBackoff = 30 * time.Second

// Retry menjalankan fn hingga maksimal maxAttempts kali.
// Exponential backoff: 1s, 2s, 4s, 8s, ... (capped di maxBackoff)
// Jitter 0-50% acak per attempt untuk hindari thundering herd.
// Context cancellation dihormati — berhenti jika ctx.Done().
func Retry(ctx context.Context, maxAttempts int, fn func() error) error {
	var lastErr error

	for attempt := range maxAttempts {
		// Cek context sebelum eksekusi
		select {
		case <-ctx.Done():
			return fmt.Errorf("retry dibatalkan: %w (attempt %d, last error: %v)",
				ctx.Err(), attempt+1, lastErr)
		default:
		}

		if err := fn(); err != nil {
			lastErr = err
			if attempt == maxAttempts-1 {
				// Ini attempt terakhir — return error
				return fmt.Errorf("retry gagal setelah %d attempt: %w", maxAttempts, err)
			}

			// Hitung backoff: 2^attempt detik, capped, + jitter
			backoff := time.Duration(math.Pow(2, float64(attempt))) * time.Second
			if backoff > maxBackoff {
				backoff = maxBackoff
			}

			// Jitter: tambah 0-50% acak
			jitter := time.Duration(rand.Float64() * 0.5 * float64(backoff))
			totalWait := backoff + jitter

			fmt.Printf("[attempt %d/%d] Gagal: %v → retry dalam %v\n",
				attempt+1, maxAttempts, err, totalWait)

			// Tunggu dengan respect terhadap context cancellation
			select {
			case <-time.After(totalWait):
			case <-ctx.Done():
				return fmt.Errorf("retry dibatalkan saat backoff: %w (last error: %v)",
					ctx.Err(), lastErr)
			}
		} else {
			fmt.Printf("[attempt %d/%d] Berhasil!\n", attempt+1, maxAttempts)
			return nil
		}
	}

	return fmt.Errorf("retry gagal setelah %d attempt: %w", maxAttempts, lastErr)
}

func main() {
	attemptCount := 0

	// Simulasi fungsi yang gagal 3x lalu sukses
	fn := func() error {
		attemptCount++
		if attemptCount <= 3 {
			return fmt.Errorf("simulasi error ke-%d", attemptCount)
		}
		return nil
	}

	ctx := context.Background()
	fmt.Println("=== Retry dengan Exponential Backoff + Jitter ===")
	fmt.Println("Fungsi akan gagal 3x lalu sukses di attempt ke-4")
	fmt.Println()

	err := Retry(ctx, 5, fn)
	if err != nil {
		fmt.Printf("\nFinal error: %v\n", err)
	} else {
		fmt.Printf("\nSukses! Total eksekusi fungsi: %d\n", attemptCount)
	}
}
