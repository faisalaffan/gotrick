//go:build !profiling_demo && !escape_analysis

// Package performance berisi demo benchmark, profiling, dan escape analysis.
//
// Benchmark:     go test -bench=. -benchmem ./performance/
// Profiling:     go run -tags=profiling_demo ./performance/
// Escape:        go run -tags=escape_analysis ./performance/
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Performance & Tooling ===")
	fmt.Println()
	fmt.Println("Benchmark:")
	fmt.Println("  go test -bench=. -benchmem ./performance/")
	fmt.Println()
	fmt.Println("Profiling (pprof):")
	fmt.Println("  go run -tags=profiling_demo ./performance/")
	fmt.Println()
	fmt.Println("Escape Analysis:")
	fmt.Println("  go run -tags=escape_analysis ./performance/")
}
