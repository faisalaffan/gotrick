// race_demo.go — Race Condition Demo + 3 Fix
// Menunjukkan data race, lalu memperbaikinya dengan:
//   1. sync.Mutex
//   2. sync/atomic
//   3. Channel
//
// Deteksi race: go run -race race_demo.go

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	const jumlahKasir = 100
	const transaksiPerKasir = 1000

	// --- DEMO RACE CONDITION (TANPA SYNC) ---
	// Jalankan dengan: go run -race race_demo.go
	// Akan keluar: WARNING: DATA RACE
	fmt.Println("=== RACE CONDITION (tanpa sinkronisasi) ===")
	{
		var totalPengunjung int
		var wg sync.WaitGroup

		for i := 0; i < jumlahKasir; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < transaksiPerKasir; j++ {
					totalPengunjung++ // DATA RACE! Multiple goroutine write tanpa sync
				}
			}()
		}
		wg.Wait()
		fmt.Printf("Total pengunjung (race): %d (seharusnya %d)\n", totalPengunjung, jumlahKasir*transaksiPerKasir)
	}
	// Hasilnya tidak prediktif dan biasanya salah.

	// --- FIX 1: sync.Mutex ---
	fmt.Println("\n=== FIX 1: sync.Mutex ===")
	{
		var totalPengunjung int
		var mu sync.Mutex
		var wg sync.WaitGroup

		for i := 0; i < jumlahKasir; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < transaksiPerKasir; j++ {
					mu.Lock()
					totalPengunjung++ // Hanya satu goroutine boleh akses dalam satu waktu
					mu.Unlock()
				}
			}()
		}
		wg.Wait()
		fmt.Printf("Total pengunjung (mutex): %d\n", totalPengunjung)
	}

	// --- FIX 2: sync/atomic ---
	fmt.Println("\n=== FIX 2: sync/atomic ===")
	{
		var totalPengunjung atomic.Int64
		var wg sync.WaitGroup

		for i := 0; i < jumlahKasir; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < transaksiPerKasir; j++ {
					totalPengunjung.Add(1) // Operasi atomik, tanpa explicit lock
				}
			}()
		}
		wg.Wait()
		fmt.Printf("Total pengunjung (atomic): %d\n", totalPengunjung.Load())
	}

	// --- FIX 3: Channel ---
	fmt.Println("\n=== FIX 3: Channel ===")
	{
		const totalIncrements = jumlahKasir * transaksiPerKasir
		totalPengunjung := 0
		inc := make(chan int)
		done := make(chan struct{})

		// Satu goroutine dedicated untuk akses counter (mutual exclusion via channel)
		go func() {
			for delta := range inc {
				totalPengunjung += delta
			}
			done <- struct{}{}
		}()

		var wg sync.WaitGroup
		for i := 0; i < jumlahKasir; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < transaksiPerKasir; j++ {
					inc <- 1 // Kirim increment via channel
				}
			}()
		}

		wg.Wait()
		close(inc) // Tutup channel — goroutine receiver akan exit
		<-done     // Tunggu receiver selesai
		fmt.Printf("Total pengunjung (channel): %d\n", totalPengunjung)
	}

	// --- CARA DETEKSI ---
	fmt.Println("\n=== Cara Deteksi Race Condition ===")
	fmt.Println("1. go run -race race_demo.go")
	fmt.Println("2. go test -race ./...")
	fmt.Println("3. go build -race ./...")
	fmt.Println("\nRace detector akan report:")
	fmt.Println("- WARNING: DATA RACE")
	fmt.Println("- Read/Write by goroutine X")
	fmt.Println("- Previous write by goroutine Y")
}
