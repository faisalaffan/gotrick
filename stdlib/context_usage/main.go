// File: context_usage.go
// Context: Background, TODO, WithCancel, WithTimeout, WithDeadline,
// Done(), Err(), dan cancel propagation.


package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 1. context.Background() — root context
	fmt.Println("=== 1. context.Background() ===")
	ctxBg := context.Background()
	fmt.Printf("Type: %T, Err: %v, Deadline: ", ctxBg, ctxBg.Err())
	if d, ok := ctxBg.Deadline(); ok {
		fmt.Println(d)
	} else {
		fmt.Println("tidak ada deadline")
	}

	// 2. context.TODO() — placeholder
	fmt.Println("\n=== 2. context.TODO() ===")
	ctxTODO := context.TODO()
	fmt.Printf("Type: %T, Err: %v\n", ctxTODO, ctxTODO.Err())

	// 3. WithCancel — cancel manual, goroutine child mendeteksi
	fmt.Println("\n=== 3. context.WithCancel — cancel propagation ===")
	ctxCancel, cancelAll := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Jalankan 3 goroutine worker
	for i := range 3 {
		wg.Add(1)
		go pekerja(ctxCancel, i+1, &wg)
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println("[main] Membatalkan semua worker...")
	cancelAll() // semua worker akan berhenti via ctx.Done()
	wg.Wait()
	fmt.Println("[main] Semua worker selesai.")

	// 4. WithTimeout — auto cancel setelah durasi
	fmt.Println("\n=== 4. context.WithTimeout — auto cancel ===")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelTimeout() // tetap panggil defer meski auto cancel

	start := time.Now()
	<-ctxTimeout.Done()
	fmt.Printf("Timeout setelah %v, Err: %v\n", time.Since(start), ctxTimeout.Err())
	fmt.Printf("  ctx.Err() == DeadlineExceeded? %v\n", ctxTimeout.Err() == context.DeadlineExceeded)

	// 5. WithDeadline — cancel di waktu spesifik
	fmt.Println("\n=== 5. context.WithDeadline — deadline spesifik ===")
	deadline := time.Now().Add(1 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	defer cancelDeadline()

	<-ctxDeadline.Done()
	fmt.Printf("Deadline lewat: Err=%v\n", ctxDeadline.Err())
	fmt.Printf("  ctx.Err() == DeadlineExceeded? %v\n", ctxDeadline.Err() == context.DeadlineExceeded)

	// 6. ctx.Done() + ctx.Err() — deteksi cancellation reason
	fmt.Println("\n=== 6. ctx.Done() + ctx.Err() ===")
	ctxTest, cancelTest := context.WithCancelCause(context.Background())
	go func() {
		time.Sleep(300 * time.Millisecond)
		cancelTest(fmt.Errorf("koneksi database terputus: kode %d", rand.Intn(100)))
	}()

	<-ctxTest.Done()
	fmt.Printf("Channel Done terbuka, Err: %v\n", ctxTest.Err())
	fmt.Printf("Cause: %v\n", context.Cause(ctxTest))

	fmt.Println("\nSelesai. Ringkasan:")
	fmt.Println("  - Background: root context")
	fmt.Println("  - TODO: placeholder saat belum yakin context apa")
	fmt.Println("  - WithCancel: cancel manual")
	fmt.Println("  - WithTimeout: cancel otomatis setelah N durasi")
	fmt.Println("  - WithDeadline: cancel di waktu spesifik")
	fmt.Println("  - ctx.Done(): channel yang ditutup saat cancel")
	fmt.Println("  - ctx.Err(): alasan cancellation")
}

// pekerja goroutine yang mendengarkan ctx.Done().
func pekerja(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  [pekerja %d] Mendeteksi cancel! Err: %v\n", id, ctx.Err())
			return
		default:
			// Simulasi kerja
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("  [pekerja %d] Bekerja...\n", id)
		}
	}
}
