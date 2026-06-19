
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
	nilai := 10
	defer fmt.Println("  nilai saat defer:", nilai) // tetap 10
	nilai = 20
	fmt.Println("  nilai setelah diubah:", nilai) // 20
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
	fileSementara, err := os.CreateTemp("", "contoh-*")
	if err != nil {
		panic(err)
	}
	// Cleanup: file dihapus saat fungsi selesai (apapun yang terjadi)
	defer os.Remove(fileSementara.Name())
	defer fileSementara.Close()

	fmt.Fprintln(fileSementara, "data penting")
	fmt.Println("  file temporary dibuat:", fileSementara.Name())
	// defer akan menjalankan fileSementara.Close() lalu os.Remove()
}

// init dipanggil otomatis sebelum main (tidak penting, hanya demo).
func init() {
	fmt.Println("=== Contoh Cleanup (simulasi) ===")
	cleanupExample()
	fmt.Println()
}
