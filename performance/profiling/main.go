// Demo profiling dengan pprof.
// Server berjalan di port 6060.
//
// Cara menjalankan:
//   go run -tags=profiling_demo ./performance/
//
// Cara profiling:
//   CPU:        go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
//   Memory:     go tool pprof http://localhost:6060/debug/pprof/heap
//   Goroutine:  go tool pprof http://localhost:6060/debug/pprof/goroutine
//
// Setelah masuk shell interaktif pprof:
//   top10         — lihat fungsi paling berat
//   list <func>   — lihat source baris per baris
//   web           — visualisasi flamegraph (butuh graphviz)
//   pdf           — simpan sebagai PDF


package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // mendaftarkan handler pprof ke DefaultServeMux
	"runtime"
	"time"
)

// cpuWork melakukan hash SHA256 berulang — simulasi beban CPU
func cpuWork(iter int) {
	for range iter {
		d := sha256.Sum256([]byte("gotrick-benchmark"))
		_ = d
	}
}

// allocWork membuat slice besar berulang — simulasi alokasi memori
func allocWork(count, size int) [][]byte {
	blocks := make([][]byte, 0, count)
	for range count {
		b := make([]byte, size)
		for i := range b {
			b[i] = byte(i % 256)
		}
		blocks = append(blocks, b)
	}
	return blocks
}

// leakWork menyimpan referensi agar GC tidak bisa collect — simulasi memory leak ringan
var leakSink [][]byte

func leakWork(count, size int) {
	for range count {
		b := make([]byte, size)
		for i := range b {
			b[i] = byte(i % 256)
		}
		leakSink = append(leakSink, b)
	}
}

func main() {
	// Jalankan background workload
	go func() {
		for {
			cpuWork(1_000_000)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			_ = allocWork(100, 10_240) // 100 x 10KB = ~1MB per siklus
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			leakWork(10, 102_400) // 10 x 100KB = ~1MB, tidak di-release
			time.Sleep(5 * time.Second)
		}
	}()

	// Print memory stats periodik
	go func() {
		var m runtime.MemStats
		for {
			runtime.ReadMemStats(&m)
			fmt.Printf("[mem] Alloc=%v MiB\tTotalAlloc=%v Mi\tSys=%v Mi\tGoroutines=%d\n",
				m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, runtime.NumGoroutine())
			time.Sleep(3 * time.Second)
		}
	}()

	log.Println("pprof server running at http://localhost:6060/debug/pprof/")
	log.Fatal(http.ListenAndServe(":6060", nil))
}
