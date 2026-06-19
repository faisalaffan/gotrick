// File: sync_primitives.go
// Mutex, RWMutex, WaitGroup, Once, dan Map.


package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 1. sync.Mutex — guard shared visitor count
	fmt.Println("=== 1. sync.Mutex — visitor counter ===")
	var (
		mu         sync.Mutex
		pengunjung int
		muWg       sync.WaitGroup
	)
	for range 10 {
		muWg.Add(1)
		go func() {
			defer muWg.Done()
			for range 1000 {
				mu.Lock()
				pengunjung++
				mu.Unlock()
			}
		}()
	}
	muWg.Wait()
	fmt.Printf("Total pengunjung (10 goroutine x 1000): %d (expected: 10000)\n", pengunjung)

	// 2. sync.RWMutex — multiple reader, single writer
	fmt.Println("\n=== 2. sync.RWMutex — reader/writer ===")
	var (
		rwmu  sync.RWMutex
		konfig string = "default"
		rwWg  sync.WaitGroup
	)
	// 3 reader goroutine
	for i := range 3 {
		rwWg.Add(1)
		go func(id int) {
			defer rwWg.Done()
			for range 3 {
				rwmu.RLock()
				fmt.Printf("  [pembaca %d] Membaca: %s\n", id, konfig)
				time.Sleep(50 * time.Millisecond)
				rwmu.RUnlock()
			}
		}(i + 1)
	}
	// 1 writer goroutine
	rwWg.Add(1)
	go func() {
		defer rwWg.Done()
		for i := range 2 {
			time.Sleep(30 * time.Millisecond)
			rwmu.Lock()
			konfig = fmt.Sprintf("versi-%d", i+1)
			fmt.Printf("  [penulis] Menulis: %s\n", konfig)
			time.Sleep(20 * time.Millisecond)
			rwmu.Unlock()
		}
	}()
	rwWg.Wait()

	// 3. sync.WaitGroup — tunggu goroutine selesai
	fmt.Println("\n=== 3. sync.WaitGroup ===")
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
			fmt.Printf("  Antrean %d selesai\n", id)
		}(i + 1)
	}
	wg.Wait()
	fmt.Println("Semua antrean selesai.")

	// 4. sync.Once — eksekusi sekali (singleton pattern)
	fmt.Println("\n=== 4. sync.Once — singleton ===")
	var once sync.Once
	for range 5 {
		go func() {
			once.Do(func() {
				fmt.Println("  [ONCE] Initialisasi hanya sekali!")
			})
		}()
	}
	time.Sleep(100 * time.Millisecond)

	// 5. sync.Map vs map + Mutex
	fmt.Println("\n=== 5. sync.Map vs map+Mutex ===")
	syncMapUsage()
	mutexMapUsage()
}

// syncMapUsage demo sync.Map concurrent.
func syncMapUsage() {
	var sm sync.Map
	var wg sync.WaitGroup

	// Writer
	for i := range 10 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("produk-%d", n), n*100)
		}(i)
	}
	wg.Wait()

	// Reader
	sm.Range(func(key, value any) bool {
		fmt.Printf("  sync.Map[%v] = %v\n", key, value)
		return true
	})

	// Load/Delete
	val, ok := sm.Load("produk-5")
	fmt.Printf("  Load produk-5: %v (exists: %t)\n", val, ok)
	sm.Delete("produk-5")
	_, ok = sm.Load("produk-5")
	fmt.Printf("  Setelah delete, produk-5 exists: %t\n", ok)

	// LoadOrStore
	actual, loaded := sm.LoadOrStore("produk-99", 9999)
	fmt.Printf("  LoadOrStore produk-99: %v (was loaded: %t)\n", actual, loaded)
	actual, loaded = sm.LoadOrStore("produk-99", 8888)
	fmt.Printf("  LoadOrStore produk-99 lagi: %v (was loaded: %t)\n", actual, loaded)
}

// mutexMapUsage demo map biasa + Mutex.
func mutexMapUsage() {
	var (
		mu sync.Mutex
		m  = make(map[string]int)
		wg sync.WaitGroup
	)
	for i := range 10 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			m[fmt.Sprintf("produk-%d", n)] = n * 100
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	mu.Lock()
	for k, v := range m {
		fmt.Printf("  mutexMap[%s] = %d\n", k, v)
	}
	mu.Unlock()
}
