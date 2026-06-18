//go:build embedding

package main

import (
	"fmt"
	"io"
	"strings"
)

// ============================================================
// 1. STRUCT EMBEDDING — Komposisi, bukan inheritance
// ============================================================

// User adalah struct dasar yang akan di-embed.
type User struct {
	Username string
	Email    string
}

// Login menampilkan pesan login untuk User.
func (u User) Login() {
	fmt.Printf("User %s login\n", u.Username)
}

// Greet mengembalikan sapaan ramah.
func (u User) Greet() string {
	return fmt.Sprintf("Halo, saya %s", u.Username)
}

// Permissions adalah jenis permission yang bisa dimiliki admin.
type Permissions []string

// Admin meng-embed User. Semua method User otomatis ter-promote ke Admin.
type Admin struct {
	User        // embedded — method promotion terjadi di sini
	Permissions Permissions
}

// Login override method Login milik User.
// Admin punya login sendiri yang berbeda dari User biasa.
func (a Admin) Login() {
	fmt.Printf("Admin %s login dengan %d permission\n", a.Username, len(a.Permissions))
}

// ============================================================
// 2. INTERFACE EMBEDDING — Gabung multiple interface
// ============================================================

// ReaderWriterCloser menggabungkan tiga interface standar jadi satu.
// Ini adalah interface embedding: io.Reader + io.Writer + io.Closer.
type ReaderWriterCloser interface {
	io.Reader
	io.Writer
	io.Closer
}

// LogReadCloser gabungan io.Reader dan io.Closer.
// Sama seperti fmt.Errorf("%w") di sisi interface.
type LogReadCloser interface {
	io.Reader
	io.Closer
}

// fileLog adalah implementasi sederhana dari LogReadCloser.
type fileLog struct {
	reader io.Reader
	closed bool
}

// Read delegasi ke reader internal.
func (f *fileLog) Read(p []byte) (int, error) {
	if f.closed {
		return 0, fmt.Errorf("fileLog sudah di-close")
	}
	return f.reader.Read(p)
}

// Close menandai fileLog sebagai tertutup.
func (f *fileLog) Close() error {
	f.closed = true
	return nil
}

// ============================================================
// 3. CONTOH LEBIH DALAM — Method promotion + override
// ============================================================

// Logger interface dengan satu method.
type Logger interface {
	Log(message string)
}

// ConsoleLogger implementasi Logger ke console.
type ConsoleLogger struct {
	Prefix string
}

func (c ConsoleLogger) Log(message string) {
	fmt.Printf("[%s] %s\n", c.Prefix, message)
}

// FileWriter bisa menulis ke string builder (simulasi file).
type FileWriter struct {
	buffer strings.Builder
}

func (f *FileWriter) Write(p []byte) (int, error) {
	return f.buffer.Write(p)
}

func (f *FileWriter) Content() string {
	return f.buffer.String()
}

// SmartLogger meng-embed ConsoleLogger dan *FileWriter.
// Method Log dari ConsoleLogger ter-promote ke SmartLogger.
// Write dari *FileWriter juga ter-promote.
type SmartLogger struct {
	ConsoleLogger           // embedded — method Log ter-promote
	*FileWriter             // embedded — method Write ter-promote
	TimestampEnabled bool
}

// Log override method Log dari ConsoleLogger.
// Sekarang SmartLogger punya Log sendiri, jadi method ConsoleLogger.Log
// tidak lagi ter-promote.
func (s SmartLogger) Log(message string) {
	if s.TimestampEnabled {
		fmt.Printf("[%s] [%d] %s\n", s.Prefix, len(message), message)
	} else {
		// Panggil method Log milik ConsoleLogger secara eksplisit
		s.ConsoleLogger.Log(message)
	}
}

// ============================================================
// 4. main — DEMONSTRASI
// ============================================================

func embeddingDemo() {
	fmt.Println("=== STRUCT EMBEDDING ===")

	user := User{Username: "budi", Email: "budi@example.com"}
	user.Login()
	fmt.Println(user.Greet())

	admin := Admin{
		User:        User{Username: "siti", Email: "siti@example.com"},
		Permissions: []string{"read", "write", "delete"},
	}

	// Method Login milik Admin dipanggil (override).
	admin.Login()

	// Method Greet milik User ter-promote ke Admin.
	// Ini adalah method promotion.
	fmt.Println(admin.Greet())

	// Field Username juga ter-promote dari User.
	fmt.Printf("Admin username: %s\n", admin.Username)

	fmt.Println("\n=== INTERFACE EMBEDDING ===")

	// fileLog mengimplementasikan LogReadCloser (io.Reader + io.Closer).
	fl := &fileLog{
		reader: strings.NewReader("hello world"),
	}

	buf := make([]byte, 5)
	n, err := fl.Read(buf)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(buf[:n]))
	}

	err = fl.Close()
	if err != nil {
		fmt.Printf("Error close: %v\n", err)
	}

	fmt.Println("\n=== METHOD PROMOTION + OVERRIDE ===")

	sl := SmartLogger{
		ConsoleLogger:    ConsoleLogger{Prefix: "APP"},
		FileWriter:       &FileWriter{},
		TimestampEnabled: true,
	}

	// Log milik SmartLogger dipanggil (override).
	sl.Log("ini pesan dengan timestamp")
	sl.Log("pesan lain")

	// Write ter-promote dari *FileWriter ke SmartLogger.
	sl.Write([]byte("data dari smart logger"))
	fmt.Printf("FileWriter content: %s\n", sl.FileWriter.Content())
}

func main() {
	embeddingDemo()
}
