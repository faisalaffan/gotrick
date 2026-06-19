
package main

import (
	"fmt"
	"os"
)

func main() {
	// ============================================================
	// Defer ordering — LIFO (stack)
	// ============================================================
	fmt.Println("=== Defer Ordering (LIFO) ===")
	defer fmt.Println("  deferred: pertama (dieksekusi terakhir)")
	defer fmt.Println("  deferred: kedua")
	defer fmt.Println("  deferred: ketiga (dieksekusi pertama)")
	fmt.Println("  fungsi main berjalan...")
	fmt.Println()

	// ============================================================
	// Defer — argument evaluation time
	// Argumen deferred function dievaluasi SAAT defer dipanggil,
	// bukan saat fungsi deferred dieksekusi.
	// ============================================================
	fmt.Println("=== Defer Argument Evaluation ===")
	x := 10
	defer fmt.Println("  nilai x saat defer:", x) // tetap 10
	x = 20
	fmt.Println("  nilai x setelah diubah:", x) // 20
	fmt.Println()

	// ============================================================
	// Panic + Recover
	// recover hanya berguna DI DALAM deferred function.
	// ============================================================
	fmt.Println("=== Panic + Recover ===")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  [recover] panic tertangkap:", r)
		}
	}()

	panicFunc()
	fmt.Println("  baris ini TIDAK akan dicetak") // tidak tercapai
}

func panicFunc() {
	fmt.Println("  akan panic...")
	panic("ada yang salah!")
}

// cleanupExample — contoh use case cleanup resource.
func cleanupExample() {
	f, err := os.CreateTemp("", "example-*")
	if err != nil {
		panic(err)
	}
	// Cleanup: file dihapus saat fungsi selesai (apapun yang terjadi)
	defer os.Remove(f.Name())
	defer f.Close()

	fmt.Fprintln(f, "data penting")
	fmt.Println("  file temporary dibuat:", f.Name())
	// defer akan menjalankan f.Close() lalu os.Remove()
}

// init dipanggil otomatis sebelum main (tidak penting, hanya demo).
func init() {
	fmt.Println("=== Contoh Cleanup (simulasi) ===")
	cleanupExample()
	fmt.Println()
}
