# Structs & Methods

Struct adalah tipe data bentukan yang mengelompokkan field-field terkait. Method dapat didefinisikan dengan value receiver (tidak mengubah struct asli) atau pointer receiver (dapat mengubah struct asli). Go menggunakan komposisi lewat embedded struct sebagai alternatif pewarisan.

## Cara Menjalankan

```bash
go run -tags=structs ./fundamentals/
```

## Source Code

```go
{{#include ../../fundamentals/structs.go}}
```
