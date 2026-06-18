// atomic_vs_mutex.go — Perbandingan sync/atomic vs sync.Mutex
// Atomic cocok untuk counter sederhana (single variable).
// Mutex diperlukan saat update multiple variable secara atomik.
//go:build atomic_vs_mutex

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	const goroutines = 10
	const iterations = 50000

	// --- ATOMIC COUNTER ---
	var atomicCounter atomic.Int64

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				atomicCounter.Add(1) // Thread-safe, tanpa lock
			}
		}()
	}
	wg.Wait()
	atomicDur := time.Since(start)
	fmt.Printf("Atomic counter: %d (durasi: %v)\n", atomicCounter.Load(), atomicDur)

	// --- MUTEX COUNTER ---
	var mutexCounter int64
	var mu sync.Mutex

	start = time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				mu.Lock()
				mutexCounter++ // Critical section
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexDur := time.Since(start)
	fmt.Printf("Mutex counter: %d (durasi: %v)\n", mutexCounter, mutexDur)

	// --- Perbandingan ---
	fmt.Printf("\nAtomic %v vs Mutex %v\n", atomicDur, mutexDur)

	// --- KAPAN MUTEX DIPERLUKAN ---
	// Atomic hanya untuk operasi single variable.
	// Kalau perlu update multiple variable secara atomik (consistent), pakai Mutex.
	fmt.Println("\n=== Contoh: Mutex untuk multiple variable ===")

	type Account struct {
		mu      sync.Mutex
		balance int64
		txCount int64
	}

	acc := &Account{}
	var wg2 sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < iterations; j++ {
				acc.mu.Lock()
				acc.balance += 100
				acc.txCount++   // Dua variable harus konsisten
				acc.mu.Unlock() // Pakai Mutex karena dua variable
			}
		}()
	}
	wg2.Wait()
	fmt.Printf("Balance: %d, TxCount: %d\n", acc.balance, acc.txCount)
	fmt.Println("Konsisten karena Mutex mengunci kedua update secara atomik.")

	// --- Ringkasan ---
	fmt.Println("\nAturan pakai:")
	fmt.Println("- Atomic: counter, flag, status — single variable, performa tinggi")
	fmt.Println("- Mutex: multiple variable update, critical section kompleks")
	fmt.Println("- Tapi untuk sebagian besar kasus, Mutex lebih aman dan cukup cepat")
}
