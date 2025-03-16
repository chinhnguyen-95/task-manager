# 📩 Task Manager API

A task management backend built in **Go** using **Clean Architecture**, **GORM**, **PostgreSQL**, and **Keycloak** for authentication via **JWT**.

> 💡 Designed to demonstrate domain-driven design, dependency injection with `wire`, Swagger API documentation, automated testing with `mockery`, and containerized infrastructure with Docker Compose.

---

## 📦 Features

- ✅ JWT-based authentication using **Keycloak**
- ✅ User registration & login
- ✅ Task CRUD with project/user association
- ✅ Task assignment and comment system
- ✅ RESTful API with **Swagger** docs
- ✅ Modular Clean Architecture
- ✅ Dependency injection via `wire`
- ✅ Unit tests with `testify` and mocks from `mockery`
- ✅ Database migrations
- ✅ Dockerized with health checks

---

## ✨ Getting Started

### 🧱 Prerequisites

- [Go](https://go.dev/) 1.23+
- [Docker](https://www.docker.com/)
- [swag](https://github.com/swaggo/swag) (`go install github.com/swaggo/swag/cmd/swag@v1.16.4`)
- [mockery](https://github.com/vektra/mockery) (`go install github.com/vektra/mockery/v2@latest`)

---

### 🥪 Run the full stack

```bash
# Build and run everything
docker compose up --build
```

📌 The API will be available at:  
`http://localhost:8080/api/v1`

---

## 📚 Swagger API Docs

Once running, visit:

> 📄 `http://localhost:8080/swagger/index.html`

To test endpoints like:

- `POST /api/v1/register`
- `POST /api/v1/login`
- `GET /api/v1/tasks/:id`
- `PUT /api/v1/tasks/:id/assign`

🔐 Use the `Authorize` button with your Bearer token from `/login`.

---

## 🔐 Authentication

- Identity Provider: [Keycloak](https://www.keycloak.org/)
- Realm: `task-manager`
- Client: `task-client`
- Grant Type: **Direct Access (Password Grant)**
- Token format: **JWT**

---

## 🗃️ Example `.env`

```env
DATABASE_DSN=host=db user=demo password=demo dbname=taskdb port=5432 sslmode=disable
KEYCLOAK_BASE_URL=http://keycloak:8080
KEYCLOAK_REALM=task-manager
KEYCLOAK_CLIENT_ID=task-client
KEYCLOAK_ADMIN_USERNAME=admin
KEYCLOAK_ADMIN_PASSWORD=admin
PORT=8080
```

---

## 🔧 Tech Stack

| Layer        | Tech                                             |
|--------------|--------------------------------------------------|
| Language     | Go                                               |
| Framework    | Gin                                              |
| Auth         | Keycloak (JWT)                                   |
| DB           | PostgreSQL + GORM                                |
| DI           | Google `wire`                                    |
| Docs         | Swaggo (`swag`) + Swagger UI                     |
| Mocks        | `mockery` + `testify`                            |
| Migration    | `goose`                                          |
| Infra        | Docker + Docker Compose                          |
| Architecture | Clean Architecture (inspired by `go-clean-arch`) |

---

## ✨ Author

Built with 💻 and ☕ by **Hucci (Chính Nguyễn)**
> Contact me on [LinkedIn](https://www.linkedin.com/in/chinhnguyen-dev) or [GitHub](https://github.com/chinhnguyen-95)
