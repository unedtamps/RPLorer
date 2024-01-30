# Go-Backend

Golang backend todo-app with gin-gonic framework

## Dependencies
__install this dependencies before run this project__
1.  [__Migrate__](https://github.com/golang-migrate/migrate)
- make migration  ``make create-migrate``
- migrate up ``make migrate-up``
- migrate down ``make migrate-down``
2. [__Sqlc__](https://github.com/sqlc-dev/sqlc)

    Sqlc generates type-safe code from SQL, see code in [__internal/repository__](./internal/repository) ``make sqlc``

3. [__Air__](https://github.com/cosmtrek/air)

    Air is yet another live-reloading command line utility for developing Go applications. Run air in your project root directory, leave it alone, and focus on your code. Edit config in [__air.toml__](./air.toml)

## Run
1. Install Golang Dependencies ``make install``
2. Run Project in development mode ``make dev``
3. Run Project in production mode ``make prod``

## Project
```
├── config
│   ├── database.go
│   └── server.go
├── internal
│   ├── db
│   │   ├── migration
│   │   │   ├── migrate_schema.down.sql
│   │   │   └── migrate_schema.up.sql
│   │   └── query
│   │       ├── todo.sql
│   │       └── users.sql
│   └── repository
│       ├── main_test.go
│       ├── store.go
│       ├── store_test.go
│       └── users_test.go
├── src
│   ├── handler
│   │   ├── main.go
│   │   ├── todo.go
│   │   └── user.go
│   ├── service
│   │   ├── main.go
│   │   ├── todo.go
│   │   └── user.go
│   └─ main.go
├── util
│   ├── error.go
│   ├── random_data.go
│   ├── response.go
│   └── users_util.go
├──.env.example
├──.gitignore
├── LICENSE
├── air.toml
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── sqlc.yaml
```
