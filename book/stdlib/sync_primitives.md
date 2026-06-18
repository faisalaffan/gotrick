# Sync Primitives

`sync.Mutex`, `sync.RWMutex` (multiple reader), `sync.WaitGroup`, `sync.Once` (singleton), dan `sync.Map` vs map+mutex.

## Cara Menjalankan

```bash
go run -tags=sync_primitives ./stdlib/
```

## Source Code

```go
{{#include ../../stdlib/sync_primitives.go}}
```
