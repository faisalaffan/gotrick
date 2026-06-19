# Token Bucket Rate Limiter

Rate limiter dengan algoritma Token Bucket. Token ditambahkan secara periodik berdasarkan rate (token per detik), dengan kapasitas maksimum (burst). Method `Allow()` mengkonsumsi satu token per request. Demo mengirim 20 request berturut-turut — hanya burst pertama yang di-allow, sisanya di-rate limit.

## Cara Menjalankan

```bash
go run ./interview/rate_limiter/
```

## Source Code

```go
{{#include ../../interview/rate_limiter/main.go}}
```
