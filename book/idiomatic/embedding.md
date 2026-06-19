# Struct Embedding

Struct embedding adalah mekanisme komposisi di Go, bukan inheritance. Method dan field dari struct yang di-embed otomatis ter-promote ke struct induk. Embedding juga bisa dilakukan pada interface untuk menggabungkan beberapa contract menjadi satu.

## Cara Menjalankan

```bash
go run ./idiomatic/embedding/ ./idiomatic/
```

## Source Code

```go
{{#include ../../idiomatic/embedding/main.go}}
```
