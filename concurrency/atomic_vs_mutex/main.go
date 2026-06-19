// atomic_vs_mutex.go — Perbandingan sync/atomic vs sync.Mutex
// Atomic cocok untuk counter sederhana (single variable).
// Mutex diperlukan saat update multiple variable secara atomik.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	const jumlahPelanggan = 10
	const transaksiPerPelanggan = 50000

	// --- ATOMIC COUNTER: hitung total pengunjung toko ---
	var pengunjungToko atomic.Int64

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < jumlahPelanggan; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < transaksiPerPelanggan; j++ {
				pengunjungToko.Add(1) // Thread-safe, tanpa lock
			}
		}()
	}
	wg.Wait()
	atomicDur := time.Since(start)
	fmt.Printf("Pengunjung toko: %d (durasi: %v)\n", pengunjungToko.Load(), atomicDur)

	// --- MUTEX COUNTER: hitung total transaksi toko ---
	var transaksiToko int64
	var mu sync.Mutex

	start = time.Now()
	for i := 0; i < jumlahPelanggan; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < transaksiPerPelanggan; j++ {
				mu.Lock()
				transaksiToko++ // Critical section
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexDur := time.Since(start)
	fmt.Printf("Transaksi toko: %d (durasi: %v)\n", transaksiToko, mutexDur)

	// --- Perbandingan ---
	fmt.Printf("\nAtomic (pengunjung) %v vs Mutex (transaksi) %v\n", atomicDur, mutexDur)

	// --- KAPAN MUTEX DIPERLUKAN ---
	// Atomic hanya untuk operasi single variable.
	// Kalau perlu update multiple variable secara atomik (consistent), pakai Mutex.
	fmt.Println("\n=== Contoh: Mutex untuk multiple variable ===")

	type Rekening struct {
		mu              sync.Mutex
		saldo           int64
		jumlahTransaksi int64
	}

	rek := &Rekening{}
	var wg2 sync.WaitGroup

	for i := 0; i < jumlahPelanggan; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for j := 0; j < transaksiPerPelanggan; j++ {
				rek.mu.Lock()
				rek.saldo += 100
				rek.jumlahTransaksi++   // Dua variable harus konsisten
				rek.mu.Unlock() // Pakai Mutex karena dua variable
			}
		}()
	}
	wg2.Wait()
	fmt.Printf("Saldo: Rp %d, Jumlah Transaksi: %d\n", rek.saldo, rek.jumlahTransaksi)
	fmt.Println("Konsisten karena Mutex mengunci kedua update secara atomik.")

	// --- Ringkasan ---
	fmt.Println("\nAturan pakai:")
	fmt.Println("- Atomic: counter pengunjung, flag status — single variable, performa tinggi")
	fmt.Println("- Mutex: update saldo + transaksi, critical section kompleks")
	fmt.Println("- Tapi untuk sebagian besar kasus, Mutex lebih aman dan cukup cepat")
}
