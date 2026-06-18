//go:build !worker_pool && !fan_out_fan_in && !pipeline && !select_pattern && !atomic_vs_mutex && !race_demo

// Package concurrency berisi contoh-contoh concurrency patterns: worker pool,
// fan-out/fan-in, pipeline, select, atomic vs mutex, race condition.
//
// Setiap file memiliki build tag sendiri. Jalankan dengan:
//
//	go run -tags=<tag> ./concurrency/
//
// Tag yang tersedia:
//
//	worker_pool     - Bounded worker pool dengan WaitGroup
//	fan_out_fan_in  - Fan-out distribusi + fan-in aggregasi
//	pipeline        - Multi-stage pipeline dengan channels
//	select_pattern  - Select: multi-channel, default, timeout, cancel
//	atomic_vs_mutex - sync/atomic vs sync.Mutex comparison
//	race_demo       - Race condition demo + 3 fix
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Concurrency Patterns ===")
	fmt.Println()
	fmt.Println("Gunakan build tag untuk menjalankan contoh spesifik:")
	fmt.Println()
	fmt.Println("  go run -tags=worker_pool     ./concurrency/   # Worker pool")
	fmt.Println("  go run -tags=fan_out_fan_in  ./concurrency/   # Fan-out / fan-in")
	fmt.Println("  go run -tags=pipeline        ./concurrency/   # Pipeline")
	fmt.Println("  go run -tags=select_pattern  ./concurrency/   # Select patterns")
	fmt.Println("  go run -tags=atomic_vs_mutex ./concurrency/   # Atomic vs mutex")
	fmt.Println("  go run -tags=race_demo       ./concurrency/   # Race condition")
}
