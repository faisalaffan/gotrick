//go:build !types && !structs && !interfaces && !errors && !goroutines && !defer_panic

// Package fundamentals berisi contoh-contoh dasar Go: types, zero values, pointers,
// structs, interfaces, error handling, goroutines, channels, defer, panic, recover.
//
// Setiap file memiliki build tag sendiri. Jalankan dengan:
//
//	go run -tags=<tag> ./fundamentals/
//
// Tag yang tersedia:
//
//	types       - Types, zero values, pointers
//	structs     - Struct definition, methods, embedding
//	interfaces  - Interface, nil interface trap
//	errors      - Error interface, sentinel errors, custom error
//	goroutines  - Goroutine, channel, buffered/unbuffered
//	defer_panic - Defer, panic, recover
package main

import "fmt"

func main() {
	fmt.Println("=== GoTrick: Fundamentals ===")
	fmt.Println()
	fmt.Println("Gunakan build tag untuk menjalankan contoh spesifik:")
	fmt.Println()
	fmt.Println("  go run -tags=types       ./fundamentals/   # Types, zero values, pointers")
	fmt.Println("  go run -tags=structs     ./fundamentals/   # Struct & methods")
	fmt.Println("  go run -tags=interfaces  ./fundamentals/   # Interface & nil trap")
	fmt.Println("  go run -tags=errors      ./fundamentals/   # Error handling")
	fmt.Println("  go run -tags=goroutines  ./fundamentals/   # Goroutines & channels")
	fmt.Println("  go run -tags=defer_panic ./fundamentals/   # Defer, panic, recover")
}
