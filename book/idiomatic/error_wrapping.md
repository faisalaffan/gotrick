# Error Wrapping

Error wrapping di Go menggunakan `fmt.Errorf` dengan verb `%w` untuk membungkus error dalam chain. `errors.Is` menelusuri chain untuk mencari sentinel error, `errors.As` mengekstrak tipe error spesifik, dan `errors.Unwrap` mengambil satu level di bawahnya.

## Cara Menjalankan

```bash
go run -tags=error_wrapping ./idiomatic/
```

## Source Code

```go
{{#include ../../idiomatic/error_wrapping.go}}
```
