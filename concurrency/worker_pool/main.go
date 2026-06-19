// worker_pool.go — Worker Pool Pattern
// N koki (worker) mengambil pesanan dari antrian, memasak, lalu mengirim hasil ke server.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// koki memproses pesanan dari channel pesanan dan mengirim hasil ke hasilPesanan.
// wg.Done() dipanggil saat koki selesai.
func koki(id int, pesanan <-chan int, hasilPesanan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for noPesanan := range pesanan {
		// Simulasi waktu memasak yang variatif
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		result := fmt.Sprintf("koki-%d selesai masak pesanan #%d → item: %d", id, noPesanan, noPesanan*noPesanan)
		hasilPesanan <- result
	}
}

func main() {
	const jumlahPesanan = 10
	const jumlahKoki = 3

	// Buffered channel agar kasir tidak blocking
	pesanan := make(chan int, jumlahPesanan)
	hasilPesanan := make(chan string, jumlahPesanan)

	var wg sync.WaitGroup

	// Start N koki goroutine
	for i := 1; i <= jumlahKoki; i++ {
		wg.Add(1)
		go koki(i, pesanan, hasilPesanan, &wg)
	}

	// Kirim pesanan ke channel
	for p := 1; p <= jumlahPesanan; p++ {
		pesanan <- p
	}
	close(pesanan) // Tidak ada pesanan lagi — koki akan exit dari range loop

	// Tunggu semua koki selesai di goroutine terpisah,
	// lalu close hasilPesanan channel agar consumer bisa exit.
	go func() {
		wg.Wait()
		close(hasilPesanan)
	}()

	// Collect dan print semua hasil
	for res := range hasilPesanan {
		fmt.Println(res)
	}

	fmt.Println("Semua koki selesai. Restoran tutup.")
}
