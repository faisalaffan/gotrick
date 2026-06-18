# HTTP Client

GET dengan `http.DefaultClient`, custom client dengan timeout, `http.NewRequestWithContext`, dan handle non-200 status.

## Cara Menjalankan

```bash
go run -tags=http_client ./stdlib/
```

## Source Code

```go
{{#include ../../stdlib/http_client.go}}
```
