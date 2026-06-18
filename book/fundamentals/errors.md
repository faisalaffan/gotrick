# Error Handling

Error di Go direpresentasikan oleh interface `error` bawaan. `errors.New()` dan `fmt.Errorf()` membuat error sederhana, sementara custom error type memungkinkan data tambahan. Sentinel errors adalah error bernilai tetap yang bisa dibandingkan dengan `==` untuk menentukan jenis kesalahan.

## Cara Menjalankan

```bash
go run -tags=errors ./fundamentals/
```

## Source Code

```go
{{#include ../../fundamentals/errors.go}}
```
