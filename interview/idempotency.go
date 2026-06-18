//go:build idempotency

package main

import (
	"fmt"
	"sync"
	"time"
)

// IdempotencyStore memastikan operasi dengan key yang sama hanya dieksekusi sekali.
// Di production: ganti sync.Map dengan Redis + SETNX + TTL.
type IdempotencyStore struct {
	results sync.Map   // key → *idempotencyEntry
}

// idempotencyEntry menyimpan result + lock per-key.
type idempotencyEntry struct {
	once     sync.Once
	done     chan struct{}
	result   interface{}
	err      error
}

// Execute menjalankan fn hanya jika key belum pernah diproses.
// Concurrent request dengan key sama akan menunggu hasil yang sama.
func (s *IdempotencyStore) Execute(key string, fn func() (interface{}, error)) (interface{}, error) {
	// Dapatkan atau buat entry untuk key ini
	entry := &idempotencyEntry{
		done: make(chan struct{}),
	}

	// sync.Map.LoadOrStore atomic — hanya 1 goroutine yang menyimpan entry baru
	actual, loaded := s.results.LoadOrStore(key, entry)
	if loaded {
		entry = actual.(*idempotencyEntry)
	}

	// Eksekusi fn hanya sekali (via sync.Once) — goroutine pertama yang sampai
	entry.once.Do(func() {
		entry.result, entry.err = fn()
		close(entry.done) // beri tahu goroutine lain yang menunggu
	})

	// Jika bukan goroutine pertama, tunggu hasil dari goroutine pertama
	if !loaded {
		// Ini goroutine pertama — result sudah ada, langsung return
		return entry.result, entry.err
	}

	// Goroutine lain: tunggu sampai goroutine pertama selesai
	<-entry.done
	return entry.result, entry.err
}

func main() {
	store := &IdempotencyStore{}
	var wg sync.WaitGroup

	idempotencyKey := "payment:txn:TRX20250618001"
	numGoroutines := 5

	fmt.Println("=== Idempotency Demo ===")
	fmt.Printf("5 goroutine concurrent dengan key: %s\n", idempotencyKey)
	fmt.Println("Hanya 1 goroutine yang akan mengeksekusi fn — sisanya menunggu hasil yang sama")
	fmt.Println()

	for i := range numGoroutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			result, err := store.Execute(idempotencyKey, func() (interface{}, error) {
				// Simulasi operasi berat (e.g., proses payment)
				fmt.Printf("  → Goroutine %02d: EKSEKUSI fn (hanya terjadi sekali!)\n", id)
				time.Sleep(2 * time.Second) // simulasi proses
				return map[string]interface{}{
					"status":   "success",
					"txn_id":   idempotencyKey,
					"amount":   250_000,
					"eksekutor": fmt.Sprintf("goroutine-%02d", id),
				}, nil
			})

			if err != nil {
				fmt.Printf("  ← Goroutine %02d: ERROR %v\n", id, err)
			} else {
				r := result.(map[string]interface{})
				fmt.Printf("  ← Goroutine %02d: OK — status=%v amount=Rp %v\n",
					id, r["status"], r["amount"])
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("\nSemua goroutine mendapat result yang identik — idempotency terjamin.")
}
