# Fan-out Fan-in

Fan-out: satu producer mendistribusikan data ke banyak worker channel. Fan-in: multiple worker mengirim hasil ke satu merged channel. Gunakan `sync.WaitGroup` untuk menunggu semua worker dan goroutine fanIn selesai, lalu close merged channel.

## Cara Menjalankan

```bash
go run ./concurrency/fan_out_fan_in/ ./concurrency/
```

## Source Code

```go
{{#include ../../concurrency/fan_out_fan_in/main.go}}
```
