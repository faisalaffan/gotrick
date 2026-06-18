# Goroutines & Channels

Goroutine adalah eksekusi ringan yang berjalan concurrent (`go f()`). Channel menjadi media komunikasi antar goroutine: unbuffered channel sinkron (saling menunggu), buffered channel asinkron hingga kapasitas penuh. Channel bisa ditutup (`close()`) dan diiterasi dengan `range`. Directional channel (`chan<-`, `<-chan`) membatasi arah pengiriman.

## Cara Menjalankan

```bash
go run -tags=goroutines ./fundamentals/
```

## Source Code

```go
{{#include ../../fundamentals/goroutines.go}}
```
