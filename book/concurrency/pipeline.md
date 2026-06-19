# Pipeline

Multi-stage pipeline dengan channels. Stage 1 generate numbers, Stage 2 mengkuadratkan, Stage 3 mencetak hasil. Setiap stage jalan di goroutine terpisah dengan channel sebagai penghubung -- channel upstream ditutup setelah selesai untuk menandakan akhir data ke stage berikutnya.

## Cara Menjalankan

```bash
go run ./concurrency/pipeline/ ./concurrency/
```

## Source Code

```go
{{#include ../../concurrency/pipeline/main.go}}
```
