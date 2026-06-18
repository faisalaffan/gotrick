//go:build !embedding && !functional_options && !error_wrapping

// Package idiomatic berisi contoh-contoh idiomatic Go: embedding (komposisi),
// functional options pattern, error wrapping (%w, errors.Is, errors.As),
// dan table-driven tests.
//
// Setiap file memiliki build tag sendiri. Jalankan dengan:
//
//	go run -tags=<tag> ./idiomatic/
//
// Tag yang tersedia:
//
//	embedding           - Struct & interface embedding
//	functional_options  - Functional options pattern
//	error_wrapping      - Error wrapping, errors.Is, errors.As
//
// Test:
//
//	go test -v ./idiomatic/
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Idiomatic Go ===")
	fmt.Println()
	fmt.Println("Gunakan build tag untuk menjalankan contoh spesifik:")
	fmt.Println()
	fmt.Println("  go run -tags=embedding           ./idiomatic/   # Embedding")
	fmt.Println("  go run -tags=functional_options  ./idiomatic/   # Functional options")
	fmt.Println("  go run -tags=error_wrapping      ./idiomatic/   # Error wrapping")
	fmt.Println()
	fmt.Println("Test:")
	fmt.Println("  go test -v ./idiomatic/")
}
