# Transfer Concurrent Antar Rekening

Simulasi transfer antar rekening dengan 10 goroutine concurrent. Lock ordering konsisten (kunci account dengan ID lebih kecil dulu) untuk mencegah deadlock. Setiap transfer mengecek saldo, melakukan debit/kredit, dan total balance diverifikasi tetap preserved.

## Cara Menjalankan

```bash
go run -tags=transfer ./interview/
```

## Source Code

```go
{{#include ../../interview/transfer.go}}
```
