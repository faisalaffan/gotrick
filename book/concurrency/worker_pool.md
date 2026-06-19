# Worker Pool

Bounded worker pool -- N worker goroutine mengambil job dari buffered channel, memproses, dan mengirim hasil ke results channel. Gunakan `sync.WaitGroup` untuk menunggu semua worker selesai, lalu close results channel agar consumer bisa exit dari range loop.

## Cara Menjalankan

```bash
go run ./concurrency/worker_pool/
```

## Source Code

```go
{{#include ../../concurrency/worker_pool/main.go}}
```
