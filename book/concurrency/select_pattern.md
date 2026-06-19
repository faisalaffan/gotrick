# Select Pattern

`select` untuk menunggu multiple channel, `select` + `default` untuk non-blocking send/receive, `select` + `time.After` untuk timeout, `select` + `ctx.Done()` untuk cancellation, dan racing pattern untuk mengambil hasil tercepat dari beberapa goroutine.

## Cara Menjalankan

```bash
go run ./concurrency/select_pattern/ ./concurrency/
```

## Source Code

```go
{{#include ../../concurrency/select_pattern/main.go}}
```
