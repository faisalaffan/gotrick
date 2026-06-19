# GoTrick — Latihan Go untuk Backend & Fintech

Kumpulan contoh kode Go yang mencakup **fundamentals**, **standard library**, **concurrency patterns**, **idiomatic Go**, **performance & tooling**, **networking**, dan **interview prep** untuk sistem finansial.

## Kenapa?

Go adalah bahasa utama di banyak fintech Indonesia (Superbank, BTPN, IAG). Repo ini dirancang sebagai latihan terstruktur dari basic sampai topik interview.

## Struktur

| Bab | Isi | File |
|-----|-----|------|
| **Fundamentals** | Types, structs, interfaces, errors, goroutines, defer/panic | 6 |
| **Standard Library** | HTTP server/client, JSON, context, sync, I/O, time | 7 |
| **Concurrency** | Worker pool, fan-in/out, pipeline, select, atomic, race | 6 |
| **Idiomatic Go** | Embedding, functional options, error wrapping, table test | 4 |
| **Performance** | Benchmark, profiling, escape analysis | 3 |
| **Networking** | HTTP retry, gRPC simulation, WebSocket | 3 |
| **Interview** | Transfer, rate limiter, graceful shutdown, retry, idempotency | 5 |

## Cara Menjalankan

Setiap contoh bisa dijalankan dengan build tag:

```bash
# Default (bantuan)
go run ./fundamentals/

# Contoh spesifik
go run ./interview/transfer/
go run ./concurrency/worker_pool/

# Test
go test ./idiomatic/ -v
go test -bench=. -benchmem ./performance/
```

## Prasyarat

- Go 1.21+
- Semua contoh menggunakan **standard library** (zero external dependencies), kecuali networking yang disimulasikan.
