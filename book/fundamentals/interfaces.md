# Interfaces

Interface di Go dideklarasikan sebagai kumpulan method signature. Satisfaction bersifat implisit -- sebuah tipe secara otomatis memenuhi interface jika memiliki method-method yang dibutuhkan, tanpa keyword `implements`. Interface kosong (`any`) dapat menampung nilai tipe apa pun. Perhatikan nil interface trap: interface nil hanya terjadi jika kedua tipe dan nilainya nil.

## Cara Menjalankan

```bash
go run ./fundamentals/interfaces/
```

## Source Code

```go
{{#include ../../fundamentals/interfaces/main.go}}
```
