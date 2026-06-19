# Functional Options Pattern

Functional options adalah pattern untuk constructor dengan parameter opsional yang fleksibel. Daripada positional argument yang kaku dan mudah salah, pattern ini menggunakan variadic `...Option` dengan tipe `type Option func(*Config)` sehingga pemanggil hanya perlu menyetel parameter yang relevan.

## Cara Menjalankan

```bash
go run ./idiomatic/functional_options/ ./idiomatic/
```

## Source Code

```go
{{#include ../../idiomatic/functional_options/main.go}}
```
