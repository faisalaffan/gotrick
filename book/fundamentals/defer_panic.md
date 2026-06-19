# Defer, Panic, Recover

`defer` menjadwalkan eksekusi fungsi setelah fungsi induk selesai, dengan urutan LIFO (tumpukan). Argumen fungsi `defer` dievaluasi saat `defer` dipanggil, bukan saat dieksekusi. `panic` menghentikan alur normal dan menaik ke tumpukan panggilan. `recover` hanya berguna jika dipanggil di dalam fungsi `defer`, untuk menangkap panic dan mencegah program crash.

## Cara Menjalankan

```bash
go run ./fundamentals/defer_panic/ ./fundamentals/
```

## Source Code

```go
{{#include ../../fundamentals/defer_panic/main.go}}
```
