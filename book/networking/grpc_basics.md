# Simulasi Pola gRPC

gRPC menggunakan protobuf untuk code generation dari file `.proto`, HTTP/2 untuk transport, dan streaming bidireksional. Simulasi ini mereplikasi arsitektur gRPC tanpa protobuf: service interface sebagai contract, implementasi server, dan client stub yang mengimplementasi interface yang sama. Transport menggunakan HTTP/1.1 + JSON, dengan pola routing `/Package/Method` seperti gRPC.

## Cara Menjalankan

```bash
# Terminal 1: jalankan server
go run -tags=server ./networking/grpc_basics/

# Terminal 2: jalankan client
go run -tags=client ./networking/grpc_basics/
```

## Source Code

### service.go (contract)

```go
{{#include ../../networking/grpc_basics/service.go}}
```

### server.go (implementasi)

```go
{{#include ../../networking/grpc_basics/server.go}}
```

### client.go (stub)

```go
{{#include ../../networking/grpc_basics/client.go}}
```
