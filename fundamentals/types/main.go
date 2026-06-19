
package main

import "fmt"

// main mendemonstrasikan tipe dasar, zero values, dan pointer di Go.
func main() {
	// ============================================================
	// Tipe dasar & zero values
	// ============================================================
	var i int
	var f float64
	var s string
	var b bool
	fmt.Println("=== Zero Values ===")
	fmt.Printf("int:     %d\n", i)   // 0
	fmt.Printf("float64: %.1f\n", f) // 0.0
	fmt.Printf("string:  %q\n", s)   // "" (string kosong)
	fmt.Printf("bool:    %t\n\n", b) // false

	// ============================================================
	// Array vs Slice
	// ============================================================
	var arr [3]int             // array — fixed size
	sl := make([]string, 0, 3) // slice — dynamic

	arr[0] = 10
	sl = append(sl, "a", "b", "c")

	fmt.Println("=== Array & Slice ===")
	fmt.Printf("array:  %v (len=%d)\n", arr, len(arr))
	fmt.Printf("slice:  %v (len=%d, cap=%d)\n\n", sl, len(sl), cap(sl))

	// ============================================================
	// Map
	// ============================================================
	m := map[string]int{"x": 1, "y": 2}
	val, ok := m["z"] // zero value jika key tidak ada
	fmt.Println("=== Map ===")
	fmt.Printf("map:    %v\n", m)
	fmt.Printf("m[\"z\"] = %d (exists=%t)\n\n", val, ok)

	// ============================================================
	// Pointer
	// ============================================================
	n := 42
	p := &n // & = alamat
	*p = 21 // * = dereference
	fmt.Println("=== Pointer ===")
	fmt.Printf("n=%d, *p=%d\n\n", n, *p)

	// new() — alokasi zero value, balik pointer
	px := new(int)
	*px = 99
	fmt.Printf("new(int) → *px=%d\n\n", *px)

	// Pointer ke struct
	type Point struct {
		X, Y int
	}
	pt := &Point{X: 3, Y: 5}
	pt.Y = 7 // otomatis dereference (*pt).Y
	fmt.Printf("pt.X=%d, pt.Y=%d\n", pt.X, pt.Y)
}
