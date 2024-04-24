# Project Stage-2024-dashboard

## Config

```yml
# config/config.yaml
server:
  port: 3000
database:
  user: ps_user
  password: SecurePassword
  database: dashboard
  host: postgres
  port: 5432
kafka:
  consumeGroup: eventviewer
  brokers:
    - redpanda-0.redpanda.redpanda.svc.cluster.local:9093
    - redpanda-1.redpanda.redpanda.svc.cluster.local:9093
    - redpanda-2.redpanda.redpanda.svc.cluster.local:9093
  schemaRegistry:
    urls:
      - redpanda-0.redpanda.redpanda.svc.cluster.local:8081
      - redpanda-1.redpanda.redpanda.svc.cluster.local:8081
      - redpanda-2.redpanda.redpanda.svc.cluster.local:8081
  auth:
    user: superuser
    password: secretpassword
```

```yml
server:
  port: 4000
database:
  user: postgres
  password: password
  database: demo
kafka:
  consumeGroup: testing
  brokers:
    - localhost:19092
  schemaRegistry:
    urls:
      - localhost:18081
```

## Dev Setup

### Docker

#### Setup

Order matters!!

```sh
docker compose -f .\docker-compose-dev.yaml up -d
```

Make sure your `.env` is configured

```sh
docker compose up -d
```

#### Building

```sh
make docker-build
```

### Dependencies

Latest version of Go and the following codegen tools.

```sh
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

Air is optional but strongly recommended.

```sh
go install github.com/cosmtrek/air@latest
```

### Code gen

```sh
make codegen
```
