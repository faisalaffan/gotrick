# Atomic vs Mutex

`sync/atomic` (`AddInt64`, `LoadInt64`) vs `sync.Mutex` -- Atomic cukup untuk counter single variable dengan performa ~2x lebih cepat. Mutex diperlukan saat update multiple variable secara atomik agar konsistensi data terjamin.

## Cara Menjalankan

```bash
go run -tags=atomic_vs_mutex ./concurrency/
```

## Source Code

```go
{{#include ../../concurrency/atomic_vs_mutex.go}}
```
