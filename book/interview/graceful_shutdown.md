# HTTP Server Graceful Shutdown

HTTP server dengan graceful shutdown menggunakan `signal.Notify` untuk SIGINT/SIGTERM. Endpoint `/slow` mensimulasikan long-running request (sleep 3 detik). Saat shutdown, `server.Shutdown(ctx)` menunggu request yang sedang berlangsung selesai dengan timeout maksimal 30 detik.

## Cara Menjalankan

```bash
go run ./interview/graceful_shutdown/ ./interview/
```

Kemudian di terminal lain, test dengan curl:

```bash
# Test long-running request, lalu tekan Ctrl+C di terminal server
curl http://localhost:8080/slow
```

## Source Code

```go
{{#include ../../interview/graceful_shutdown/main.go}}
```
