# HTTP Server & Middleware

Server HTTP dengan **Gin** (`github.com/gin-gonic/gin`) — routing, middleware logging, JSON response, dan route grouping. Gin menggunakan `net/http` di bawahnya, cocok untuk production fintech.

## Cara Menjalankan

```bash
go run ./stdlib/httpserver/
```

## Source Code

```go
{{#include ../../stdlib/httpserver/main.go}}
```
