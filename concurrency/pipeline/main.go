// pipeline.go — Multi-Stage Pipeline
// Stage 1: generate numbers
// Stage 2: square numbers
// Stage 3: print results
// Tiap stage jalan di goroutine terpisah. Channel upstream ditutup setelah selesai.

package main

import "fmt"

// stage1 — generator: mengirim angka 1..n ke channel
func stage1(out chan<- int, n int) {
	for i := 1; i <= n; i++ {
		out <- i
	}
	close(out) // Selesai generate, tutup channel
}

// stage2 — square: membaca dari in, mengkuadratkan, kirim ke out
func stage2(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * val
	}
	close(out)
}

// stage3 — consumer: membaca dari in dan mencetak
func stage3(in <-chan int) {
	for val := range in {
		fmt.Printf("Hasil: %d\n", val)
	}
}

func main() {
	const n = 10

	// Pipeline channels
	genOut := make(chan int)
	sqOut := make(chan int)

	// Jalankan stage di goroutine terpisah
	go stage1(genOut, n) // producer
	go stage2(genOut, sqOut) // middleware
	stage3(sqOut) // consumer — jalan di main goroutine

	// Perhatikan urutan close:
	// stage1 close(genOut) → stage2 selesai range → stage2 close(sqOut) → stage3 selesai range
	fmt.Println("Pipeline selesai.")
}
