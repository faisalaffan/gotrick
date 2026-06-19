# Context

`context.Background()`, `WithCancel`, `WithTimeout`, `WithDeadline`, `ctx.Done()`, `ctx.Err()`, dan propagasi cancellation ke goroutine child.

## Cara Menjalankan

```bash
go run ./stdlib/context_usage/
```

## Source Code

```go
{{#include ../../stdlib/context_usage/main.go}}
```
