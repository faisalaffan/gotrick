package main

import "fmt"

// main mendemonstrasikan tipe dasar, zero values, dan pointer di Go.
func main() {
	// ============================================================
	// Tipe dasar & zero values
	// ============================================================
	var angka int
	var pecahan float64
	var teks string
	var bendera bool
	fmt.Println("=== Zero Values ===")
	fmt.Printf("int:     %d\n", angka)     // 0
	fmt.Printf("float64: %.1f\n", pecahan) // 0.0
	fmt.Printf("string:  %q\n", teks)      // "" (string kosong)
	fmt.Printf("bool:    %t\n\n", bendera) // false

	// ============================================================
	// Array vs Slice
	// ============================================================
	var arr [3]int                // array — fixed size
	hewan := make([]string, 0, 3) // slice — dynamic

	arr[0] = 10
	hewan = append(hewan, "kucing", "anjing", "burung")

	fmt.Println("=== Array & Slice ===")
	fmt.Printf("array:  %v (len=%d)\n", arr, len(arr))
	fmt.Printf("slice:  %v (len=%d, cap=%d)\n\n", hewan, len(hewan), cap(hewan))

	// ============================================================
	// Map
	// ============================================================
	hargaBuah := map[string]int{"mangga": 5000, "pisang": 3000}
	harga, ketemu := hargaBuah["apel"] // zero value jika key tidak ada
	fmt.Println("=== Map ===")
	fmt.Printf("hargaBuah: %v\n", hargaBuah)
	fmt.Printf("hargaBuah[\"apel\"] = %d (exists=%t)\n\n", harga, ketemu)

	// ============================================================
	// Pointer
	// ============================================================
	nilai := 42
	pointer := &nilai // & = alamat
	*pointer = 21     // * = dereference
	fmt.Println("=== Pointer ===")
	fmt.Printf("nilai=%d, *pointer=%d\n\n", nilai, *pointer)

	// new() — alokasi zero value, balik pointer
	angkaBaru := new(int)
	*angkaBaru = 99
	fmt.Printf("new(int) → *angkaBaru=%d\n\n", *angkaBaru)

	// Pointer ke struct
	type Koordinat struct {
		X, Y int
	}
	titik := &Koordinat{X: 3, Y: 5}
	titik.Y = 7 // otomatis dereference (*titik).Y
	fmt.Printf("titik.X=%d, titik.Y=%d\n", titik.X, titik.Y)
}
