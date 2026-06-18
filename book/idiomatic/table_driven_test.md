# Table-Driven Test

Table-driven test adalah idiom utama testing di Go. Test case didefinisikan sebagai slice of struct dengan input dan expected output, lalu diiterasi dalam subtest `t.Run()`. Subtest bisa dijalankan paralel dengan `t.Parallel()` untuk mempercepat eksekusi.

## Cara Menjalankan

```bash
go test -v ./idiomatic/
```

## Source Code

```go
{{#include ../../idiomatic/table_driven_test.go}}
```
