package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"testing"
)

// ============================================================
// KALKULATOR — fungsi yang akan di-test
// ============================================================

// CalculatorError untuk error spesifik kalkulator.
type CalculatorError struct {
	Op  string
	Err error
}

func (c CalculatorError) Error() string {
	return fmt.Sprintf("kalkulator: %s: %v", c.Op, c.Err)
}

func (c CalculatorError) Unwrap() error {
	return c.Err
}

var ErrDivisionByZero = errors.New("pembagian dengan nol")

// Add menjumlahkan dua bilangan.
func Add(a, b int) int { return a + b }

// Subtract mengurangkan b dari a.
func Subtract(a, b int) int { return a - b }

// Multiply mengalikan dua bilangan.
func Multiply(a, b int) int { return a * b }

// Divide membagi a dengan b. Error jika b == 0.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, CalculatorError{Op: "Divide", Err: ErrDivisionByZero}
	}
	return a / b, nil
}

// Power menghitung a pangkat b.
func Power(a, b int) float64 {
	return math.Pow(float64(a), float64(b))
}

// ============================================================
// TABLE-DRIVEN TEST — Calculator
//
// Cara run:
//
//	go test -v ./idiomatic/ -run TestCalculator
//	go test -v ./idiomatic/ -run "TestCalculator/Add"
//	go test -v ./idiomatic/ -run "TestCalculator/Divide"
// ============================================================

func TestCalculator(t *testing.T) {
	// Setup: bisa dipakai untuk inisialisasi bersama.
	// Dalam kasus sederhana seperti ini, setup minimal.
	_ = t // suppress unused warning jika setup kosong

	// --- Subtest: Add ---
	t.Run("Add", func(t *testing.T) {
		tests := []struct {
			name string
			a, b int
			want int
		}{
			{"positive", 2, 3, 5},
			{"negative", -1, -2, -3},
			{"mixed", -5, 10, 5},
			{"zero", 0, 0, 0},
			{"large", 1000000, 2000000, 3000000},
		}

		for _, tt := range tests {
			tt := tt // capture range variable (Go < 1.22), idiom untuk parallel test
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel() // subtests jalan paralel
				got := Add(tt.a, tt.b)
				if got != tt.want {
					t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
				}
			})
		}
	})

	// --- Subtest: Subtract ---
	t.Run("Subtract", func(t *testing.T) {
		tests := []struct {
			name string
			a, b int
			want int
		}{
			{"positive", 10, 3, 7},
			{"negative result", 3, 10, -7},
			{"zero", 5, 5, 0},
			{"negative numbers", -5, -3, -2},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Subtract(tt.a, tt.b)
				if got != tt.want {
					t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
				}
			})
		}
	})

	// --- Subtest: Multiply ---
	t.Run("Multiply", func(t *testing.T) {
		tests := []struct {
			name string
			a, b int
			want int
		}{
			{"positive", 4, 5, 20},
			{"negative", -3, 4, -12},
			{"double negative", -2, -6, 12},
			{"zero", 7, 0, 0},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				got := Multiply(tt.a, tt.b)
				if got != tt.want {
					t.Errorf("Multiply(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
				}
			})
		}
	})

	// --- Subtest: Divide ---
	t.Run("Divide", func(t *testing.T) {
		// Subtests untuk Divide tidak paralel karena beberapa test
		// memodifikasi state yang sama (tidak ada di sini, tapi untuk variasi).

		t.Run("success", func(t *testing.T) {
			t.Parallel()
			tests := []struct {
				name     string
				a, b     int
				want     int
				wantErr  bool
				wantErrIs error
			}{
				{"simple", 10, 2, 5, false, nil},
				{"negative", -10, 2, -5, false, nil},
				{"zero numerator", 0, 5, 0, false, nil},
				{"not exact", 7, 3, 2, false, nil}, // integer division
			}

			for _, tt := range tests {
				tt := tt
				t.Run(tt.name, func(t *testing.T) {
					t.Parallel()
					got, err := Divide(tt.a, tt.b)

					if tt.wantErr {
						t.Errorf("Divide(%d, %d) seharusnya error", tt.a, tt.b)
					}
					if err != nil {
						t.Errorf("Divide(%d, %d) unexpected error: %v", tt.a, tt.b, err)
					}
					if got != tt.want {
						t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
					}
				})
			}
		})

		// --- Subtest: Division by zero ---
		t.Run("divide by zero", func(t *testing.T) {
			t.Parallel()

			_, err := Divide(10, 0)
			if err == nil {
				t.Fatal("Divide(10, 0) seharusnya error, tapi nil")
			}

			// Cek error chain dengan errors.Is.
			if !errors.Is(err, ErrDivisionByZero) {
				t.Errorf("errors.Is(err, ErrDivisionByZero) = false; want true. err=%v", err)
			}

			// Cek custom error type dengan errors.As.
			var calcErr CalculatorError
			if !errors.As(err, &calcErr) {
				t.Errorf("errors.As(err, &calcErr) = false; want true. err=%v", err)
			}
			if calcErr.Op != "Divide" {
				t.Errorf("calcErr.Op = %s; want Divide", calcErr.Op)
			}
		})
	})

	// Teardown: cleanup setelah semua test selesai.
	// Dalam kasus nyata: close DB, close file, dll.
	t.Log("Setup dan Teardown: otomatis via t.Run scope")
}

// ============================================================
// TABLE-DRIVEN TEST — String Utilities (contoh tambahan)
// ============================================================

func TestStringUtils(t *testing.T) {
	// Test ToUpper via strings package.
	t.Run("ToUpper", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			input string
			want  string
		}{
			{"hello", "HELLO"},
			{"golang", "GOLANG"},
			{"already UPPER", "ALREADY UPPER"},
			{"mixed CASE", "MIXED CASE"},
			{"", ""},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.input, func(t *testing.T) {
				t.Parallel()
				got := strings.ToUpper(tt.input)
				if got != tt.want {
					t.Errorf("ToUpper(%q) = %q; want %q", tt.input, got, tt.want)
				}
			})
		}
	})

	// Test TrimSpace.
	t.Run("TrimSpace", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			input string
			want  string
		}{
			{"  hello  ", "hello"},
			{"\tspaces\n", "spaces"},
			{"no-trim", "no-trim"},
			{"", ""},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(fmt.Sprintf("%q", tt.input), func(t *testing.T) {
				t.Parallel()
				got := strings.TrimSpace(tt.input)
				if got != tt.want {
					t.Errorf("TrimSpace(%q) = %q; want %q", tt.input, got, tt.want)
				}
			})
		}
	})
}

// ============================================================
// BENCHMARK — contoh sederhana (opsional)
// ============================================================

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(100, 3) //nolint:errcheck
	}
}
