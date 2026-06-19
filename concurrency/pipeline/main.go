// pipeline.go — Multi-Stage Pipeline
// Stage 1: generate panjang sisi tanah
// Stage 2: hitung luas tanah (sisi × sisi)
// Stage 3: cetak hasil luas
// Tiap stage jalan di goroutine terpisah. Channel upstream ditutup setelah selesai.

package main

import "fmt"

// genSisi — generator: mengirim panjang sisi 1..n ke channel
func genSisi(out chan<- int, n int) {
	for sisi := 1; sisi <= n; sisi++ {
		out <- sisi
	}
	close(out) // Selesai generate, tutup channel
}

// hitungLuas — membaca sisi dari in, menghitung luas (sisi×sisi), kirim ke out
func hitungLuas(in <-chan int, out chan<- int) {
	for sisi := range in {
		out <- sisi * sisi
	}
	close(out)
}

// cetakLuas — consumer: membaca luas dari in dan mencetak
func cetakLuas(in <-chan int) {
	for luas := range in {
		fmt.Printf("Luas tanah: %d m²\n", luas)
	}
}

func main() {
	const jumlahTanah = 10

	// Pipeline channels
	sisiOut := make(chan int)
	luasOut := make(chan int)

	// Jalankan stage di goroutine terpisah
	go genSisi(sisiOut, jumlahTanah)   // producer
	go hitungLuas(sisiOut, luasOut)     // middleware
	cetakLuas(luasOut)                  // consumer — jalan di main goroutine

	// Perhatikan urutan close:
	// genSisi close(sisiOut) → hitungLuas selesai range → hitungLuas close(luasOut) → cetakLuas selesai range
	fmt.Println("\nPipeline selesai.")
}
