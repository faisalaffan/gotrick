# Profiling dengan pprof

pprof adalah toolkit profiling bawaan Go. Dengan mengimpor `net/http/pprof`, server HTTP di port 6060 menyediakan endpoint untuk CPU profile, heap (memory), dan goroutine dump. File ini mensimulasikan tiga workload: SHA256 hashing (CPU-bound), alokasi memori periodik, dan memory leak ringan.

## Cara Menjalankan

```bash
go run ./performance/profiling/ ./performance/
```

## Source Code

```go
{{#include ../../performance/profiling/main.go}}
```
