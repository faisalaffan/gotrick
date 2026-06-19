# Atomic vs Mutex

`sync/atomic` (`AddInt64`, `LoadInt64`) vs `sync.Mutex` -- Atomic cukup untuk counter single variable dengan performa ~2x lebih cepat. Mutex diperlukan saat update multiple variable secara atomik agar konsistensi data terjamin.

## Cara Menjalankan

```bash
go run ./concurrency/atomic_vs_mutex/
```

## Source Code

```go
{{#include ../../concurrency/atomic_vs_mutex/main.go}}
```
