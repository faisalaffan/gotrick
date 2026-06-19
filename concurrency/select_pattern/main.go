// select_pattern.go — Select Statement Patterns
// 1. select tunggu multiple channel
// 2. select dengan default — non-blocking
// 3. select dengan time.After — timeout
// 4. select dengan ctx.Done() — cancellation
// 5. Racing: ambil hasil tercepat dari beberapa goroutine

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// slowAPICall mensimulasikan request API dengan delay acak.
func slowAPICall(ctx context.Context, name string, delay time.Duration) string {
	select {
	case <-time.After(delay):
		return fmt.Sprintf("%s selesai dalam %v", name, delay)
	case <-ctx.Done():
		return fmt.Sprintf("%s dibatalkan: %v", name, ctx.Err())
	}
}

func main() {
	// --- Pattern 1: select tunggu multiple channel ---
	fmt.Println("=== Pattern 1: Multiple channel ===")
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() { time.Sleep(50 * time.Millisecond); ch1 <- "dari ch1" }()
	go func() { time.Sleep(100 * time.Millisecond); ch2 <- "dari ch2" }()

	select {
	case msg := <-ch1:
		fmt.Println("ch1 menang:", msg)
	case msg := <-ch2:
		fmt.Println("ch2 menang:", msg)
	}

	// --- Pattern 2: non-blocking send/receive ---
	fmt.Println("\n=== Pattern 2: Non-blocking ===")
	ch := make(chan int, 1)
	ch <- 42

	select {
	case val := <-ch:
		fmt.Println("Terima:", val)
	default:
		fmt.Println("Channel kosong — tidak blocking")
	}

	select {
	case ch <- 99:
		fmt.Println("Berhasil kirim 99")
	default:
		fmt.Println("Channel penuh — tidak blocking")
	}

	// --- Pattern 3: timeout ---
	fmt.Println("\n=== Pattern 3: Timeout ===")
	slow := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slow <- "data akhirnya sampai"
	}()

	select {
	case msg := <-slow:
		fmt.Println(msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout! Operasi terlalu lambat.")
	}

	// --- Pattern 4: cancellation dengan context ---
	fmt.Println("\n=== Pattern 4: Cancellation ===")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel() // batalkan setelah 50ms
	}()

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Selesai normal")
	case <-ctx.Done():
		fmt.Println("Dibatalkan:", ctx.Err())
	}

	// --- Pattern 5: Racing — ambil hasil tercepat ---
	fmt.Println("\n=== Pattern 5: Racing ===")
	ctx2, cancel2 := context.WithCancel(context.Background())
	defer cancel2()

	results := make(chan string, 3)
	services := []struct {
		name  string
		delay time.Duration
	}{
		{"server-A", time.Duration(rand.Intn(300)+50) * time.Millisecond},
		{"server-B", time.Duration(rand.Intn(300)+50) * time.Millisecond},
		{"server-C", time.Duration(rand.Intn(300)+50) * time.Millisecond},
	}

	for _, svc := range services {
		go func(s struct {
			name  string
			delay time.Duration
		}) {
			results <- slowAPICall(ctx2, s.name, s.delay)
		}(svc)
	}

	// Ambil hasil pertama yang datang
	select {
	case winner := <-results:
		fmt.Println("Pemenang:", winner)
		cancel2() // Batalkan worker lain
	case <-time.After(1 * time.Second):
		fmt.Println("Semua timeout")
	}

	time.Sleep(50 * time.Millisecond) // Biarkan goroutine lain selesai log
	fmt.Println("Select pattern demo selesai.")
}
