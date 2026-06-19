package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Performance & Tooling ===")
	fmt.Println()
	fmt.Println("Benchmark:")
	fmt.Println("  go test -bench=. -benchmem ./performance/")
	fmt.Println()
	fmt.Println("Profiling (pprof):")
	fmt.Println("  go run ./performance/profiling/")
	fmt.Println()
	fmt.Println("Escape Analysis:")
	fmt.Println("  go run ./performance/escape_analysis/")
}
