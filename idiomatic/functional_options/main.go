
package main

import (
	"fmt"
	"time"
)

// ============================================================
// FUNCTIONAL OPTIONS PATTERN
//
// Kenapa lebih baik dari constructor dengan banyak parameter:
//   - Parameter opsional tidak perlu dikirim (zero value)
//   - Urutan parameter tidak penting
//   - Mudah ditambah opsi baru tanpa break API
//   - Default value bisa ditentukan di constructor
//   - Self-documenting: WithPort(8080) lebih jelas dari NewServer("", 8080, 0, false)
// ============================================================

// ServerConfig menyimpan konfigurasi server.
// Semua field lowercase (private) — hanya bisa di-set via options.
type ServerConfig struct {
	host    string
	port    int
	timeout time.Duration
	tls     bool
}

// String implementasi fmt.Stringer untuk debugging.
func (s ServerConfig) String() string {
	protocol := "http"
	if s.tls {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d (timeout: %v)", protocol, s.host, s.port, s.timeout)
}

// Option adalah fungsi yang memodifikasi ServerConfig.
// Ini adalah tipe functional option.
type Option func(*ServerConfig)

// defaultConfig mengembalikan ServerConfig dengan nilai default.
// Options nantinya akan meng-override nilai-nilai ini.
func defaultConfig() ServerConfig {
	return ServerConfig{
		host:    "localhost",
		port:    3000,
		timeout: 30 * time.Second,
		tls:     false,
	}
}

// --- Option constructors ---

// WithHost mengatur host server.
func WithHost(h string) Option {
	return func(c *ServerConfig) {
		c.host = h
	}
}

// WithPort mengatur port server.
func WithPort(p int) Option {
	return func(c *ServerConfig) {
		c.port = p
	}
}

// WithTimeout mengatur timeout server. Menerima time.Duration.
func WithTimeout(d time.Duration) Option {
	return func(c *ServerConfig) {
		c.timeout = d
	}
}

// WithTLS mengaktifkan TLS. Tanpa parameter karena default-nya false.
func WithTLS() Option {
	return func(c *ServerConfig) {
		c.tls = true
	}
}

// ============================================================
// SERVER — Produk akhir dari functional options
// ============================================================

// Server merepresentasikan server yang sudah dikonfigurasi.
type Server struct {
	config ServerConfig
}

// NewServer adalah constructor dengan functional options pattern.
// Parameter opsional dikirim sebagai variadic ...Option.
//
// Keuntungan dibandingkan constructor tradisional:
//
//	NewServer(host, port, timeout, tls)       — urutan kaku, parameter wajib
//	NewServer(opts...)                         — fleksibel, hanya set yang perlu
//
// Secara fundamental: dengan functional options, API tidak berubah saat
// kita menambah field baru ke config. Dengan positional parameters,
// setiap field baru adalah breaking change.
func NewServer(opts ...Option) *Server {
	// Mulai dari default config.
	cfg := defaultConfig()

	// Apply setiap option ke config.
	for _, opt := range opts {
		opt(&cfg)
	}

	return &Server{config: cfg}
}

// Start simulasi menjalankan server.
func (s *Server) Start() error {
	fmt.Printf("Starting server: %s\n", s.config.String())
	return nil
}

// ============================================================
// CONTOH LAIN — Database config dengan functional options
// ============================================================

// DBConfig konfigurasi koneksi database.
type DBConfig struct {
	dsn             string
	maxOpen         int
	maxIdle         int
	connMaxLifetime time.Duration
}

// DBOption functional option untuk DBConfig.
type DBOption func(*DBConfig)

func defaultDBConfig() DBConfig {
	return DBConfig{
		dsn:             "postgres://localhost:5432/toko_buku",
		maxOpen:         25,
		maxIdle:         5,
		connMaxLifetime: 5 * time.Minute,
	}
}

// WithDSN mengatur connection string database.
func WithDSN(dsn string) DBOption {
	return func(c *DBConfig) {
		c.dsn = dsn
	}
}

// WithMaxOpen mengatur max open connections.
func WithMaxOpen(n int) DBOption {
	return func(c *DBConfig) {
		c.maxOpen = n
	}
}

// NewDB membuat koneksi database (simulasi).
// Menggunakan functional options seperti NewServer.
func NewDB(opts ...DBOption) *DBConfig {
	cfg := defaultDBConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return &cfg
}

// ============================================================
// main — DEMONSTRASI
// ============================================================

func functionalOptionsDemo() {
	fmt.Println("=== FUNCTIONAL OPTIONS PATTERN ===")

	// Tanpa option — semua default.
	s1 := NewServer()
	s1.Start()

	fmt.Println()

	// Dengan beberapa option — hanya set yang perlu.
	s2 := NewServer(
		WithPort(8080),
		WithTLS(),
		WithTimeout(60*time.Second),
	)
	s2.Start()

	fmt.Println()

	// Opsi lain: override host juga.
	s3 := NewServer(
		WithHost("api.example.com"),
		WithPort(443),
		WithTLS(),
	)
	s3.Start()

	fmt.Println()

	// Contoh untuk database — skema toko online.
	db := NewDB(
		WithDSN("postgres://admin:rahasia@db.tokobuku.com:5432/inventory"),
		WithMaxOpen(100),
	)
	fmt.Printf("DB config: %+v\n", *db)
}

func main() {
	functionalOptionsDemo()
}
