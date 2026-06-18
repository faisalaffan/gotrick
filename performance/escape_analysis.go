// Demonstrasi escape analysis: variable yang ESCAPE ke heap vs yang tetap di stack.
//
// Cara menjalankan:
//   go run -tags=escape_analysis ./performance/
//
// Cek escape analysis:
//   go build -gcflags="-m" ./performance/
//   go build -gcflags="-m -m" ./performance/  # lebih detail
//
// Aturan umum:
//   - Stack:  lebih cepat, otomatis dibersihkan saat fungsi return
//   - Heap:   perlu GC, lebih lambat, perlu alokasi
//   - Escape: Go memindahkan variable ke heap jika compiler tidak bisa
//             buktikan bahwa referensinya tidak keluar dari fungsi

//go:build escape_analysis

package main

import "fmt"

// ==============================================================
// Case 1: Return pointer dari fungsi — ESCAPE ke heap
// ==============================================================
// KENAPA: pointer `v` dikembalikan ke caller. Compiler tidak bisa
// memastikan apakah caller masih menyimpan pointer setelah fungsi
// return. Demi safety, `v` dialokasikan di heap.
func newInt(n int) *int {
	v := n // v ESCAPE ke heap
	return &v
}

// ==============================================================
// Case 2: Simpan ke interface{} — ESCAPE ke heap
// ==============================================================
// KENAPA: interface{} adalah tipe abstrak yang menyimpan dynamic type.
// Go compiler tidak bisa menentukan ukuran konkret di stack, dan
// interface{} selalu menyimpan value sebagai pointer (dalam konteks
// tertentu). Nilai `42` dan struct point akan escape.
func printAny() {
	var x any = 42 // 42 ESCAPE — disimpan sebagai interface{}
	p := point{X: 1, Y: 2}
	var y any = p // p ESCAPE — disimpan sebagai interface{}
	fmt.Println(x, y)
}

type point struct {
	X, Y int
}

// ==============================================================
// Case 3: Closure capture variable — ESCAPE ke heap
// ==============================================================
// KENAPA: closure bisa dipanggil setelah fungsi induknya return.
// Karena `base` mungkin dirujuk lebih lama dari stack frame, compiler
// memindahkannya ke heap. Variable hasil penjumlahan juga escape
// karena closure dikembalikan sebagai function value (pointer).
func makeAdder(base int) func(int) int {
	// base ESCAPE — di-capture oleh closure
	return func(x int) int {
		return base + x
	}
}

// ==============================================================
// Case 4: Value type (tidak escape) — tetap di stack
// ==============================================================
// KENAPA: `a` adalah int, `b` hasil penjumlahan, keduanya tidak
// dirujuk di luar scope fungsi. Ukuran sudah diketahui saat compile.
// Tidak ada interface{}, pointer return, atau closure capture.
func noEscape(a, b int) int {
	// a dan b tetap di stack
	sum := a + b // sum juga di stack
	return sum
}

func main() {
	// Case 1
	p := newInt(42)
	fmt.Println("Case 1 - return pointer:", *p)

	// Case 2
	printAny()

	// Case 3
	add5 := makeAdder(5)
	fmt.Println("Case 3 - closure:", add5(10))

	// Case 4
	r := noEscape(3, 4)
	fmt.Println("Case 4 - no escape:", r)
}
