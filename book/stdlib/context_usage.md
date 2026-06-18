# Context

`context.Background()`, `WithCancel`, `WithTimeout`, `WithDeadline`, `ctx.Done()`, `ctx.Err()`, dan propagasi cancellation ke goroutine child.

## Cara Menjalankan

```bash
go run -tags=context_usage ./stdlib/
```

## Source Code

```go
{{#include ../../stdlib/context_usage.go}}
```
