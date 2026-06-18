# Benchmark

Benchmark mengukur performa kode secara kuantitatif. File ini mendemonstrasikan tiga perbandingan penting: string concatenation (`+` vs `strings.Builder` vs `fmt.Sprintf`), slice pre-allocation dengan kapasitas tetap vs append dinamis, dan map lookup vs slice linear lookup untuk dataset kecil dan besar.

## Cara Menjalankan

```bash
go test -bench=. -benchmem ./performance/
```

## Source Code

```go
{{#include ../../performance/benchmark_test.go}}
```
