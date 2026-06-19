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

// panggilAPI mensimulasikan request API dengan delay acak.
func panggilAPI(ctx context.Context, nama string, delay time.Duration) string {
	select {
	case <-time.After(delay):
		return fmt.Sprintf("%s selesai dalam %v", nama, delay)
	case <-ctx.Done():
		return fmt.Sprintf("%s dibatalkan: %v", nama, ctx.Err())
	}
}

func main() {
	// --- Pattern 1: select tunggu multiple channel ---
	fmt.Println("=== Pattern 1: Multiple channel ===")
	cuaca := make(chan string, 1)
	kurs := make(chan string, 1)

	go func() { time.Sleep(50 * time.Millisecond); cuaca <- "CekCuaca: 32°C, cerah" }()
	go func() { time.Sleep(100 * time.Millisecond); kurs <- "CekKurs: 1 USD = Rp 16.200" }()

	select {
	case msg := <-cuaca:
		fmt.Println("Cuaca menang:", msg)
	case msg := <-kurs:
		fmt.Println("Kurs menang:", msg)
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
	apiLambat := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		apiLambat <- "data cuaca akhirnya sampai"
	}()

	select {
	case msg := <-apiLambat:
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

	hasil := make(chan string, 3)
	daftarAPI := []struct {
		name  string
		delay time.Duration
	}{
		{"CekCuaca", time.Duration(rand.Intn(300)+50) * time.Millisecond},
		{"CekKurs", time.Duration(rand.Intn(300)+50) * time.Millisecond},
		{"CekTraffic", time.Duration(rand.Intn(300)+50) * time.Millisecond},
	}

	for _, svc := range daftarAPI {
		go func(s struct {
			name  string
			delay time.Duration
		}) {
			hasil <- panggilAPI(ctx2, s.name, s.delay)
		}(svc)
	}

	// Ambil hasil pertama yang datang
	select {
	case pemenang := <-hasil:
		fmt.Println("Pemenang:", pemenang)
		cancel2() // Batalkan worker lain
	case <-time.After(1 * time.Second):
		fmt.Println("Semua timeout")
	}

	time.Sleep(50 * time.Millisecond) // Biarkan goroutine lain selesai log
	fmt.Println("Select pattern demo selesai.")
}
