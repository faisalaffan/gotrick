# Retry dengan Exponential Backoff + Jitter

Retry pattern dengan exponential backoff menggunakan `math.Pow(2, attempt)` untuk base delay (1s, 2s, 4s, 8s) dan jitter acak 0-50% per attempt untuk mencegah thundering herd. Context cancellation dihormati di setiap tahap. Return error final dengan informasi attempt terakhir.

## Cara Menjalankan

```bash
go run ./interview/retry_backoff/ ./interview/
```

## Source Code

```go
{{#include ../../interview/retry_backoff/main.go}}
```
