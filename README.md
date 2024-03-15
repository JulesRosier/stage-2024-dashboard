# Project Stage-2024-dashboard

## Env vars

example

```env
DB_USER=postgres
DB_PASSWORD=password
DB_DATABASE=testing
DB_HOST=localhost
DB_PORT=5432

SEED_BROKER=localhost:19092
REGISTRY=localhost:18081
```

## Dev Setup

### Dependencies

Latest version of Go and the following codegen tools.

```sh
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

### Code gen

```sh
sqlc generate
```
