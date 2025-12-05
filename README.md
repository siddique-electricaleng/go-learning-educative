# Backend Engineering with Go
> This project was built following the YouTube video [Building a Production API in Golang from Scratch (Ecommerce project)](https://youtu.be/s3XItrqfccw?si=koUVU72SetlnyluS) by Tiago, with additional help from ChatGPT.

## Overview
A learning project for backend REST API development using Go.

## Features
- RESTful API design
- Database integration
- Error handling
- Clean Layered Architecture
- DB Migrations
- Dependency Injections

---

## Getting Started

### Prerequisites
- Go 1.25.4 or higher
- Git
- Dependencies:
    - github.com/go-chi/chi/v5 v5.2.3
    - github.com/jackc/pgx/v5 v5.7.6
    - golang.org/x/crypto v0.37.0
- Tools:
    - [goose](https://github.com/pressly/goose) - Database migrations
    - [sqlc](https://sqlc.dev/) - SQL code generation
    - [Air](https://github.com/cosmtrek/air) - Hot reloading for development

### Installation
1. Clone the repository
2. Navigate to the project directory
3. Run `go mod download`

---

## Project Structure
- `./cmd/main.go` - Entry point
- `./internal/products/handlers.go` - Sample HTTP request handlers for a service (products)
- `./internal/products/services.go` - Sample Service handlers for a service (product)
- `./internal/adapters/postgresql/sqlc/models.go` - Data structures
- `./internal/adapters/postgresql/sqlc/db.go` - Database operations

---

## Contributing
Pull requests are welcome.
