## 🔌 gRPC Support

This project also exposes a **gRPC API** alongside the REST API, using the same service and domain logic.

### 🧠 Why gRPC?

- High-performance communication for internal services
- Type-safe contracts using Protocol Buffers
- Supports streaming and advanced scenarios
- Easily integrates with microservices, mobile, or other Go apps

---

### 🚀 Start gRPC Server

The gRPC server is started **automatically** along with the REST server when running:

```bash
docker compose up --build
```

It listens on:
```
localhost:50051
```

---

### 📄 Proto Definition

Located at:
```
proto/task_manager.proto
```

Generated Go code lives in:
```
pkg/pb/taskmanager/
```

To regenerate manually:

```bash
protoc \
  --go_out=. \
  --go-grpc_out=. \
  --proto_path=proto \
  proto/task_manager.proto
```

---

### 🧪 Test gRPC with grpcui

Install [grpcui](https://github.com/fullstorydev/grpcui):

```bash
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
grpcui -plaintext localhost:50051
```

> Opens a web UI for testing gRPC calls interactively.

---

### 🔐 Authentication for gRPC

- gRPC supports JWT tokens via the `Authorization` header (as metadata).
- Example:  
  `Authorization: Bearer <your-token>`

- Applied via a custom **JWT Unary Interceptor** (see `internal/grpc/middleware`).
- Authentication is **skipped automatically** for `/Login` and `/Register` via method-based filtering.

---

### 🧬 gRPC Architecture Overview

| Layer        | Component                          |
|--------------|-------------------------------------|
| Proto        | `proto/task_manager.proto`          |
| Services     | `internal/grpc/*.go`          |
| Contracts    | `pkg/pb/taskmanager/*.pb.go`        |
| Interceptor  | `JWTUnaryInterceptor`               |
| Integration  | Shared with domain + REST services  |

