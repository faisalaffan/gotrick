# Idempotency Pattern

Idempotency memastikan operasi dengan key yang sama hanya dieksekusi sekali. Pattern ini kritis untuk sistem fintech. `sync.Map` sebagai storage (di production: Redis + SETNX + TTL), lock per-key via `sync.Once`, goroutine pertama mengeksekusi fn dan sisanya menunggu hasil identik melalui channel.

## Cara Menjalankan

```bash
go run ./interview/idempotency/ ./interview/
```

## Source Code

```go
{{#include ../../interview/idempotency/main.go}}
```
