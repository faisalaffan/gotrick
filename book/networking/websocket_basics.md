# Simulasi WebSocket

WebSocket menyediakan komunikasi full-duplex melalui satu koneksi TCP. Simulasi ini mendemonstrasikan pola WebSocket tanpa library eksternal: HTTP upgrade handshake (101 Switching Protocols), read/write loop dengan protokol line-based, dan connection hijacking via `http.Hijacker`. Di production, gunakan `gorilla/websocket`.

## Cara Menjalankan

```bash
# Terminal 1: jalankan server
go run -tags=server ./networking/websocket_basics/

# Terminal 2: jalankan client
go run -tags=client ./networking/websocket_basics/
```

## Source Code

### server.go

```go
{{#include ../../networking/websocket_basics/server.go}}
```

### client.go

```go
{{#include ../../networking/websocket_basics/client.go}}
```
