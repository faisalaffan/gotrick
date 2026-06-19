# Types & Zero Values

Tipe data dasar Go terdiri dari integer, float, string, bool, serta koleksi seperti array, slice, dan map. Setiap variabel yang dideklarasi tanpa inisialisasi akan memiliki zero value sesuai tipenya (misal `0` untuk int, `""` untuk string). Pointer (`&`, `*`, `new()`) memungkinkan akses langsung ke alamat memori.

## Cara Menjalankan

```bash
go run ./fundamentals/types/ ./fundamentals/
```

## Source Code

```go
{{#include ../../fundamentals/types/main.go}}
```
