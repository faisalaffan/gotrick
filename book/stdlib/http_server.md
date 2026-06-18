# HTTP Server & Middleware

Server HTTP dengan `http.HandleFunc`, custom mux via `http.NewServeMux`, middleware pattern (logging wrapper), dan JSON response.

## Cara Menjalankan

```bash
go run -tags=http_server ./stdlib/
```

## Source Code

```go
{{#include ../../stdlib/http_server.go}}
```
