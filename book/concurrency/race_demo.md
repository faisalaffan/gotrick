# Race Condition

Demonstrasi race condition -- multiple goroutine increment shared counter tanpa sinkronisasi menghasilkan hasil yang tidak prediktif. Tiga fix: `sync.Mutex`, `sync/atomic`, dan channel. Deteksi dengan `go run -race`.

## Cara Menjalankan

```bash
go run ./concurrency/race_demo/ ./concurrency/
```

## Source Code

```go
{{#include ../../concurrency/race_demo/main.go}}
```
