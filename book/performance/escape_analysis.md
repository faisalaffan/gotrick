# Escape Analysis

Escape analysis adalah optimisasi compiler Go yang menentukan apakah sebuah variable dialokasikan di stack (cepat, otomatis dibersihkan) atau heap (lebih lambat, perlu GC). Variable escape ke heap ketika: direturn sebagai pointer, disimpan ke `interface{}`, atau di-capture oleh closure. Value type yang tidak dirujuk di luar fungsi tetap di stack.

## Cara Menjalankan

```bash
go run ./performance/escape_analysis/ ./performance/

# Cek analisa escape:
go build -gcflags="-m" ./performance/
```

## Source Code

```go
{{#include ../../performance/escape_analysis/main.go}}
```
