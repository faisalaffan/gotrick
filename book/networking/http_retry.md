# HTTP Retry dengan Exponential Backoff

HTTP client dengan retry logic menggunakan exponential backoff (1s, 2s, 4s) dan jitter acak plus-minus 20% untuk mencegah thundering herd. Retry hanya untuk error 5xx dan network error. Total timeout dikontrol via context, dan semua attempt akan berhenti jika context selesai.

## Cara Menjalankan

```bash
# Coba dengan URL yang mengembalikan 500
go run -tags=http_retry ./networking/ https://httpbin.org/status/500

# Coba dengan connection refused
go run -tags=http_retry ./networking/ http://localhost:9999
```

## Source Code

```go
{{#include ../../networking/http_retry.go}}
```
