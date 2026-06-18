// File: io_bufio.go
// Custom io.Reader, io.Copy, bufio.Scanner, bufio.Writer, io.TeeReader.

//go:build io_bufio

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// --- Custom Reader: uppercaseReader ---

// uppercaseReader mengubah semua byte menjadi uppercase.
type uppercaseReader struct {
	src io.Reader
}

func (r *uppercaseReader) Read(p []byte) (int, error) {
	n, err := r.src.Read(p)
	if n > 0 {
		// Ubah byte ke uppercase (hanya ASCII)
		for i := range n {
			if p[i] >= 'a' && p[i] <= 'z' {
				p[i] -= 32
			}
		}
	}
	return n, err
}

func main() {
	// 1. Custom io.Reader — uppercase
	fmt.Println("=== 1. Custom io.Reader (uppercase) ===")
	input := "hello, world! ini adalah test reader.\n"
	upperReader := &uppercaseReader{src: strings.NewReader(input)}
	data, _ := io.ReadAll(upperReader)
	fmt.Print(string(data))

	// 2. io.Copy — copy dari reader ke writer
	fmt.Println("\n=== 2. io.Copy ===")
	src := strings.NewReader("Data dari source reader\n")
	dst := &strings.Builder{}
	written, err := io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Copied %d bytes: %s", written, dst.String())

	// 3. bufio.Scanner — baca baris per baris
	fmt.Println("=== 3. bufio.Scanner — baris per baris ===")
	multiLine := `Baris pertama
Baris kedua
Baris ketiga dengan angka 123
Baris keempat`
	scanner := bufio.NewScanner(strings.NewReader(multiLine))
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		fmt.Printf("  Line %d: %s\n", lineNum, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("ERROR scanner: %v\n", err)
	}

	// 4. bufio.Scanner — baca file (jika ada argumen)
	fmt.Println("\n=== 4. bufio.Scanner — baca file ===")
	// Buat file temp dulu
	tmpFile := "/tmp/gotrick_io_test.txt"
	content := "apple\nbanana\ncherry\ndurian\n"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile)

	f, err := os.Open(tmpFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		fmt.Printf("  Buah: %s\n", fileScanner.Text())
	}

	// 5. bufio.Writer — buffered write
	fmt.Println("\n=== 5. bufio.Writer — buffered write ===")
	bufWriter := bufio.NewWriter(os.Stdout)
	bufWriter.WriteString("  Ini ")
	bufWriter.WriteString("ditulis ")
	bufWriter.WriteString("dengan ")
	bufWriter.WriteString("bufio.Writer!\n")
	bufWriter.Flush() // flush penting! tanpa ini data mungkin tidak tertulis
	fmt.Println("  (Flush memastikan semua data ditulis ke underlying writer)")

	// 6. io.TeeReader — baca sambil copy
	fmt.Println("\n=== 6. io.TeeReader — baca + copy simultan ===")
	var teeBuf strings.Builder
	teeSrc := strings.NewReader("Data untuk TeeReader — akan muncul dua kali.\n")
	teeReader := io.TeeReader(teeSrc, &teeBuf)

	// Baca TeeReader — hasilnya juga otomatis masuk ke teeBuf
	teeData, _ := io.ReadAll(teeReader)
	fmt.Printf("  Hasil ReadAll: %s", string(teeData))
	fmt.Printf("  Hasil Tee (tercopy): %s", teeBuf.String())

	// 7. io.Copy dengan batasan (LimitReader)
	fmt.Println("\n=== 7. io.LimitReader ===")
	longSrc := strings.NewReader("Ini adalah teks panjang yang hanya akan dibaca sebagian oleh LimitReader.")
	limited := io.LimitReader(longSrc, 20)
	limitedData, _ := io.ReadAll(limited)
	fmt.Printf("  Terbaca (limited 20 byte): %s\n", string(limitedData))
}
