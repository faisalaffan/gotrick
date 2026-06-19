
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Account merepresentasikan rekening bank.
// sync.Mutex melindungi Balance dari race condition concurrent.
type Account struct {
	ID      int
	Balance int
	mu      sync.Mutex
}

// Transfer memindahkan dana antar dua account.
// Lock ordering berdasarkan ID lebih kecil dulu untuk menghindari deadlock.
// Mengecek balance cukup sebelum transfer.
func Transfer(from, to *Account, amount int) error {
	// Lock ordering: kunci account dengan ID lebih kecil dulu
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
		return fmt.Errorf("transfer ke akun sendiri: from=%d == to=%d", from.ID, to.ID)
	}

	// Kritikal section dimulai
	if from.Balance < amount {
		from.mu.Unlock()
		to.mu.Unlock()
		return fmt.Errorf("saldo tidak cukup: account %d punya %d, perlu %d", from.ID, from.Balance, amount)
	}

	from.Balance -= amount
	to.Balance += amount
	// Kritikal section selesai

	from.mu.Unlock()
	to.mu.Unlock()
	return nil
}

// BalanceWithLock aman dibaca concurrent karena pakai lock.
func BalanceWithLock(a *Account) int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}

func main() {
	// Inisialisasi 3 account dengan saldo awal
	accounts := []*Account{
		{ID: 1, Balance: 100_000},
		{ID: 2, Balance: 50_000},
		{ID: 3, Balance: 25_000},
	}

	// Hitung total awal
	totalAwal := 0
	for _, a := range accounts {
		totalAwal += BalanceWithLock(a)
	}
	fmt.Printf("Total saldo awal: Rp %d\n\n", totalAwal)

	var wg sync.WaitGroup
	numGoroutines := 10

	for i := range numGoroutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Pilih pengirim dan penerima acak (pastikan berbeda)
			fromIdx := rand.Intn(len(accounts))
			toIdx := rand.Intn(len(accounts))
			for toIdx == fromIdx {
				toIdx = rand.Intn(len(accounts))
			}

			amount := (rand.Intn(10) + 1) * 1000
			err := Transfer(accounts[fromIdx], accounts[toIdx], amount)
			if err != nil {
				fmt.Printf("[goroutine %02d] GAGAL: %s\n", id, err)
			} else {
				fmt.Printf("[goroutine %02d] OK: Rp %d dari Account %d → Account %d\n",
					id, amount, accounts[fromIdx].ID, accounts[toIdx].ID)
			}
		}(i)
	}

	wg.Wait()

	// Verifikasi total balance preserved
	totalAkhir := 0
	for _, a := range accounts {
		totalAkhir += BalanceWithLock(a)
	}
	fmt.Printf("\n=== HASIL AKHIR ===\n")
	for _, a := range accounts {
		fmt.Printf("Account %d: Rp %d\n", a.ID, BalanceWithLock(a))
	}
	fmt.Printf("Total saldo akhir: Rp %d\n", totalAkhir)
	fmt.Printf("Preserved? %v (tidak ada dana hilang/bertambah)\n", totalAwal == totalAkhir)
}
