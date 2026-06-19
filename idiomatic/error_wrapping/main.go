
package main

import (
	"errors"
	"fmt"
)

// ============================================================
// SENTINEL ERRORS — error fixed sebagai nilai package-level
// ============================================================

// Sentinel errors didefinisikan dengan fmt.Errorf atau errors.New.
// Ini adalah error final yang akan dicek dengan errors.Is.
var (
	ErrNotFound   = errors.New("data tidak ditemukan")
	ErrForbidden  = errors.New("akses ditolak")
	ErrDuplicate  = errors.New("data sudah ada")
	ErrRepository = errors.New("kesalahan repository")
)

// ============================================================
// CUSTOM ERROR TYPE — dengan Unwrap() untuk chain
// ============================================================

// ValidationError adalah custom error type yang membawa field informasi.
// errors.As bisa mengekstrak tipe ini dari error chain.
type ValidationError struct {
	Field string
	Value any
	Err   error // wrapped error (bisa nil)
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal untuk field '%s' (nilai: %v): %v", v.Field, v.Value, v.Err)
}

// Unwrap mengembalikan error yang dibungkus.
// Dengan Unwrap, ValidationError jadi bagian dari error chain.
func (v ValidationError) Unwrap() error {
	return v.Err
}

// NewValidationError membuat ValidationError. err bisa nil.
func NewValidationError(field string, value any, err error) ValidationError {
	return ValidationError{Field: field, Value: value, Err: err}
}

// ============================================================
// LAYER REPOSITORY — sumber error "bawah"
// ============================================================

// Repository adalah layer data store.
// Di sini error "mentah" berasal.
type Repository struct {
	// Simulasi data store.
	data map[string]string
}

// NewRepository membuat repository baru.
func NewRepository() *Repository {
	return &Repository{
		data: make(map[string]string),
	}
}

// Get mengambil data dari repository.
// Mengembalikan ErrNotFound jika key tidak ada.
func (r *Repository) Get(key string) (string, error) {
	val, ok := r.data[key]
	if !ok {
		return "", fmt.Errorf("%w: key=%s", ErrNotFound, key)
	}
	return val, nil
}

// Save menyimpan data ke repository.
// Mengembalikan ErrDuplicate jika key sudah ada.
func (r *Repository) Save(key, value string) error {
	if _, ok := r.data[key]; ok {
		return fmt.Errorf("repository.Save: %w", ErrDuplicate)
	}
	r.data[key] = value
	return nil
}

// ============================================================
// LAYER SERVICE — wrapper dengan konteks tambahan
// ============================================================

// Service adalah layer business logic.
// Di sini error dari repository di-wrap dengan konteks tambahan.
type Service struct {
	repo *Repository
}

// NewService membuat service baru.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// GetUser mengambil data user.
// Error dari repository di-wrap dengan konteks tambahan.
func (s *Service) GetUser(username string) (string, error) {
	val, err := s.repo.Get("user:" + username)
	if err != nil {
		// Wrap error repository dengan konteks service.
		// %w membuat error ini bisa di-unwrap dan dicek dengan errors.Is/As.
		return "", fmt.Errorf("service.GetUser: %w", err)
	}
	return val, nil
}

// CreateUser membuat user baru.
func (s *Service) CreateUser(username, data string) error {
	err := s.repo.Save("user:"+username, data)
	if err != nil {
		// Wrap dengan konteks tambahan.
		return fmt.Errorf("service.CreateUser: %w", err)
	}
	return nil
}

// ProcessData simulasi validasi + penyimpanan.
// Menunjukkan errors.As untuk custom error type.
func (s *Service) ProcessData(key string, value string) error {
	// Validasi: value tidak boleh kosong.
	if value == "" {
		return NewValidationError(key, value, errors.New("value tidak boleh kosong"))
	}

	return s.CreateUser(key, value)
}

// ============================================================
// ERROR HANDLER — consumer error chain
// ============================================================

// ErrorHandler mendemonstrasikan errors.Is, errors.As, errors.Unwrap.
type ErrorHandler struct{}

func (h ErrorHandler) Handle(err error) {
	if err == nil {
		fmt.Println("Tidak ada error")
		return
	}

	fmt.Printf("\n=== Error: %v ===\n", err)

	// errors.Is — cek apakah ada sentinel error di chain.
	// errors.Is berjalan sepanjang error chain (dengan Unwrap).
	if errors.Is(err, ErrNotFound) {
		fmt.Println("  -> errors.Is: ErrNotFound ditemukan di chain")
	}
	if errors.Is(err, ErrDuplicate) {
		fmt.Println("  -> errors.Is: ErrDuplicate ditemukan di chain")
	}
	if errors.Is(err, ErrForbidden) {
		fmt.Println("  -> errors.Is: ErrForbidden ditemukan di chain")
	}

	// errors.As — extract custom error type dari chain.
	var valErr ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("  -> errors.As: ValidationError ter-ekstrak (field=%s, value=%v)\n",
			valErr.Field, valErr.Value)
	}

	// errors.Unwrap — dapatkan error langsung di bawahnya.
	// Ini hanya satu level. errors.Is dan errors.As sudah traverse semua level.
	unwrapped := errors.Unwrap(err)
	if unwrapped != nil {
		fmt.Printf("  -> errors.Unwrap: %v\n", unwrapped)
	} else {
		fmt.Println("  -> errors.Unwrap: nil (error asli/sentinel)")
	}
}

// ============================================================
// main — DEMONSTRASI
// ============================================================

func errorWrappingDemo() {
	repo := NewRepository()
	svc := NewService(repo)
	handler := ErrorHandler{}

	// Simpan data.
	err := svc.CreateUser("alice", "data-alice")
	if err != nil {
		fmt.Printf("Gagal create: %v\n", err)
	}

	// 1. Coba ambil user yang tidak ada.
	fmt.Println("\n--- Test 1: Get user tidak ada ---")
	_, err = svc.GetUser("budi")
	handler.Handle(err)

	// 2. Coba duplikasi user.
	fmt.Println("\n--- Test 2: Duplikasi user ---")
	err = svc.CreateUser("alice", "data-alice-lagi")
	// Oops, kita sudah create alice — coba create lagi via ProcessData.
	err = svc.CreateUser("alice", "data-alice-lagi")
	handler.Handle(err)

	// 3. Validasi error dengan custom type.
	fmt.Println("\n--- Test 3: Validasi gagal ---")
	err = svc.ProcessData("email", "")
	handler.Handle(err)

	// 4. Error success — tidak ada error.
	fmt.Println("\n--- Test 4: Tidak ada error ---")
	handler.Handle(nil)

	// 5. Cek langsung dengan errors.Is pada error yang di-wrap.
	fmt.Println("\n--- Test 5: errors.Is langsung ---")
	baseErr := fmt.Errorf("lapisan luar: %w", fmt.Errorf("lapisan tengah: %w", ErrForbidden))
	if errors.Is(baseErr, ErrForbidden) {
		fmt.Println("  ErrForbidden ditemukan walau dibungkus 2 layer")
	}
}

func main() {
	errorWrappingDemo()
}
