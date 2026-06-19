
package main

import (
	"fmt"
	"time"
)

func main() {
	// ============================================================
	// Goroutine dasar dengan anonymous function
	// ============================================================
	fmt.Println("=== Goroutine Dasar ===")

	selesai := make(chan bool)

	go func() {
		fmt.Println("  [goroutine] Halo dari goroutine!")
		selesai <- true
	}()

	<-selesai // blocking sampai goroutine selesai
	fmt.Println()

	// ============================================================
	// Unbuffered channel — send/receive blocking
	// ============================================================
	fmt.Println("=== Unbuffered Channel ===")

	kanal := make(chan int)

	go func() {
		kanal <- 42 // blocking sampai main menerima
		fmt.Println("  [goroutine] 42 terkirim")
	}()

	angka := <-kanal // blocking sampai ada data
	fmt.Printf("  [main]   menerima %d\n\n", angka)

	// ============================================================
	// Buffered channel — tidak blocking sampai penuh
	// ============================================================
	fmt.Println("=== Buffered Channel ===")

	bufCh := make(chan string, 2)
	bufCh <- "satu"
	bufCh <- "dua"
	// bufCh <- "tiga" // ini akan blocking — buffer penuh!

	fmt.Println("  dari channel:", <-bufCh)
	fmt.Println("  dari channel:", <-bufCh)
	fmt.Println()

	// ============================================================
	// Close channel + range over channel
	// ============================================================
	fmt.Println("=== Close + Range ===")

	nums := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		nums <- i * 10
	}
	close(nums) // wajib close sebelum range

	for bilangan := range nums {
		fmt.Printf("  %d\n", bilangan)
	}
	fmt.Println()

	// ============================================================
	// Directional channel
	// ============================================================
	fmt.Println("=== Directional Channel ===")

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// send-only channel (chan<- string)
	ping := func(msg string, out chan<- string) {
		out <- msg
	}

	// receive-only channel (<-chan string)
	pong := func(in <-chan string, out chan<- string) {
		msg := <-in
		out <- msg + "!"
	}

	ping("halo", pings)
	pong(pings, pongs)
	fmt.Println("  hasil:", <-pongs)

	// Beri waktu goroutine lain selesai (hanya demo)
	time.Sleep(10 * time.Millisecond)
}
