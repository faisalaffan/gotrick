
package main

import (
	"errors"
	"fmt"
)

// ============================================================
// Sentinel errors — pakai errors.New
// ============================================================

var (
	ErrNotFound = errors.New("data tidak ditemukan")
	ErrInvalid  = errors.New("input tidak valid")
)

// ============================================================
// Custom error type
// ============================================================

// ValidationError struct dengan field tambahan.
type ValidationError struct {
	Field string
	Value any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validasi error: field=%q value=%v", e.Field, e.Value)
}

// Fungsi yang mengembalikan error.
func findUser(id int) error {
	if id <= 0 {
		return fmt.Errorf("findUser: %w", ErrInvalid) // bungkus sentinel
	}
	if id > 100 {
		return &ValidationError{Field: "id", Value: id}
	}
	return nil
}

// ============================================================
// Type assertion untuk custom error
// ============================================================

func main() {
	// Sentinel errors
	fmt.Println("=== Sentinel Errors ===")
	err := findUser(0)
	if errors.Is(err, ErrInvalid) {
		fmt.Println("Ketemu sentinel ErrInvalid:", err)
	}
	fmt.Println()

	// errors.Is dengan wrapped error
	err2 := findUser(0)
	if errors.Is(err2, ErrInvalid) {
		fmt.Println("errors.Is bekerja walau dibungkus fmt.Errorf %w:", err2)
	}
	fmt.Println()

	// Custom error + type assertion
	fmt.Println("=== Custom Error Type ===")
	err3 := findUser(200)
	if err3 != nil {
		var ve *ValidationError
		if errors.As(err3, &ve) {
			fmt.Printf("Type assertion berhasil: field=%q value=%v\n", ve.Field, ve.Value)
		} else {
			fmt.Println("Error biasa:", err3)
		}
	}
	fmt.Println()

	// errors.New
	fmt.Println("=== errors.New / fmt.Errorf ===")
	errDasar := errors.New("kesalahan sederhana")
	errGabung := fmt.Errorf("gagal di %s: %w", "init", errDasar)
	fmt.Println("errGabung:", errGabung)
	fmt.Println("Unwrap:", errors.Unwrap(errGabung))
}
