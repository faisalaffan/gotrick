<p align="center">
  <a href="README.md">English</a> ·
  <a href="README.id.md">Bahasa Indonesia</a>
</p>

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="assets/02_BANNER_DARK.png">
    <source media="(prefers-color-scheme: light)" srcset="assets/03_BANNER_LIGHT.png">
    <img alt="GoTrick Banner" src="assets/02_BANNER_DARK.png" width="100%">
  </picture>
</p>

<p align="center">
  <img src="assets/01_LOGO.png" alt="GoTrick Logo" width="120">
</p>

<h1 align="center">GoTrick</h1>

<p align="center">
  <strong>Structured Go exercises — from fundamentals to fintech interview prep.</strong>
</p>

<p align="center">
  <a href="https://github.com/faisalaffan/gotrick/actions/workflows/mdbook.yml"><img src="https://github.com/faisalaffan/gotrick/actions/workflows/mdbook.yml/badge.svg" alt="GitHub Actions"></a>
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go" alt="Go version">
  <img src="https://img.shields.io/badge/dependencies-zero-00ADD8" alt="Zero dependencies">
</p>

---

## About

GoTrick is a curated collection of Go code examples designed for backend engineers preparing for fintech roles. Covers everything from basic types to production-grade concurrency patterns — all with zero external dependencies.

## Documentation

**[Read the full docs →](https://faisalaffan.github.io/gotrick/)**

## Topics

| Section | Topics | Files |
|---------|--------|-------|
| **Fundamentals** | Types, zero values, pointers, structs, interfaces, errors, goroutines, channels, defer, panic, recover | 6 |
| **Standard Library** | Fiber HTTP, JSON, context, sync primitives, I/O, bufio, time | 7 |
| **Concurrency** | Worker pool, fan-out/fan-in, pipeline, select, atomic vs mutex, race condition | 6 |
| **Idiomatic Go** | Embedding, functional options, error wrapping, table-driven tests | 4 |
| **Performance** | Benchmark, pprof profiling, escape analysis | 3 |
| **Networking** | HTTP retry, gRPC, WebSocket | 3 |
| **Interview Prep** | Concurrent transfer, rate limiter, graceful shutdown, retry backoff, idempotency | 5 |

## Quick Start

```bash
# Clone
git clone git@github.com:faisalaffan/gotrick.git
cd gotrick

# Run any example (pick the tag)
go run -tags=transfer     ./interview/   # Concurrent transfer + deadlock avoidance
go run -tags=worker_pool  ./concurrency/  # Bounded worker pool
go run -tags=interfaces   ./fundamentals/ # Interface basics + nil trap

# Default view (shows available tags)
go run ./fundamentals/

# Run tests
go test ./... -v
go test -bench=. -benchmem ./performance/
```

## How It Works

Every file uses **build tags** (`//go:build <tag>`) so you can run examples independently:

```bash
# file: transfer.go → //go:build transfer
go run -tags=transfer ./interview/
```

See [full docs](https://faisalaffan.github.io/gotrick/) for all available tags and explanations.

## Prerequisites

- Go 1.21 or later
- Everything uses **standard library only** — no external dependencies required

## Design System

<p align="center">
  <img src="assets/00_DESIGN_SYSTEM.png" alt="GoTrick Design System" width="600">
</p>

## License

MIT — see [LICENSE](LICENSE).
