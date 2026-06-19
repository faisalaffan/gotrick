
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Rekening merepresentasikan rekening bank.
// sync.Mutex melindungi Saldo dari race condition concurrent.
type Rekening struct {
	ID      int
	Saldo   int
	mu      sync.Mutex
}

// Transfer memindahkan dana antar dua rekening.
// Lock ordering berdasarkan ID lebih kecil dulu untuk menghindari deadlock.
// Mengecek saldo cukup sebelum transfer.
func Transfer(from, to *Rekening, amount int) error {
	// Lock ordering: kunci rekening dengan ID lebih kecil dulu
	// Ini mencegah circular wait → deadlock
	if from.ID < to.ID {
		from.mu.Lock()
		to.mu.Lock()
	} else if from.ID > to.ID {
		to.mu.Lock()
		from.mu.Lock()
	} else {
		// Transfer ke diri sendiri — hanya perlu 1 lock
		from.mu.Lock()
		defer from.mu.Unlock()
		return fmt.Errorf("transfer ke rekening sendiri: from=%d == to=%d", from.ID, to.ID)
	}

	// Kritikal section dimulai
	if from.Saldo < amount {
		from.mu.Unlock()
		to.mu.Unlock()
		return fmt.Errorf("saldo tidak cukup: rekening %d punya %d, perlu %d", from.ID, from.Saldo, amount)
	}

	from.Saldo -= amount
	to.Saldo += amount
	// Kritikal section selesai

	from.mu.Unlock()
	to.mu.Unlock()
	return nil
}

// SaldoWithLock aman dibaca concurrent karena pakai lock.
func SaldoWithLock(a *Rekening) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Saldo
}

func main() {
	// Inisialisasi 3 rekening dengan saldo awal
	rekening := []*Rekening{
		{ID: 1, Saldo: 100_000},
		{ID: 2, Saldo: 50_000},
		{ID: 3, Saldo: 25_000},
	}

	// Hitung total awal
	totalAwal := 0
	for _, r := range rekening {
		totalAwal += SaldoWithLock(r)
	}
	fmt.Printf("Total saldo awal: Rp %d\n\n", totalAwal)

	var wg sync.WaitGroup
	jumlahGoroutine := 10

	for i := range jumlahGoroutine {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Pilih pengirim dan penerima acak (pastikan berbeda)
			dariIdx := rand.Intn(len(rekening))
			keIdx := rand.Intn(len(rekening))
			for keIdx == dariIdx {
				keIdx = rand.Intn(len(rekening))
			}

			jumlah := (rand.Intn(10) + 1) * 1000
			err := Transfer(rekening[dariIdx], rekening[keIdx], jumlah)
			if err != nil {
				fmt.Printf("[goroutine %02d] GAGAL: %s\n", id, err)
			} else {
				fmt.Printf("[goroutine %02d] OK: Rp %d dari Rekening %d → Rekening %d\n",
					id, jumlah, rekening[dariIdx].ID, rekening[keIdx].ID)
			}
		}(i)
	}

	wg.Wait()

	// Verifikasi total saldo preserved
	totalAkhir := 0
	for _, r := range rekening {
		totalAkhir += SaldoWithLock(r)
	}
	fmt.Printf("\n=== HASIL AKHIR ===\n")
	for _, r := range rekening {
		fmt.Printf("Rekening %d: Rp %d\n", r.ID, SaldoWithLock(r))
	}
	fmt.Printf("Total saldo akhir: Rp %d\n", totalAkhir)
	fmt.Printf("Preserved? %v (tidak ada dana hilang/bertambah)\n", totalAwal == totalAkhir)
}
