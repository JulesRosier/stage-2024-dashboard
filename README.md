# Project Stage-2024-dashboard

## Config

```yml
# Default configs
# config/config.yaml
server:
  port: 3000
  bind: 127.0.0.1
  debug: false
database:
  user: postgres
  password: postgres
  database: event-viewer
  host: 127.0.0.1
  port: 5432
kafka:
  brokers: []
  consumeGroup: ""
  schemaRegistry:
    urls: []
  auth:
    user: ""
    password: ""
logger:
  level: INFO
indexing:
  interval: 1h0m0s
alert:
  slack:
    webhookURL: ""
  eventDeltas: []
  interval: 1h0m0
```

```yml
# Example configs
# config/config.yaml
server:
  port: 4000
  debug: true
database:
  user: postgres
  password: password
  database: dashboard
kafka:
  consumeGroup: dashboard
  brokers:
    - localhost:19092
  schemaRegistry:
    urls:
      - localhost:18081
indexing:
  interval: 5m
alert:
  interval: 200h
  slack:
    webhookURL: "https://hooks.slack.com/services/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"
  eventDeltas:
    - topicA: bike_reserved
      topicB: bike_picked_up
      index: index_bike_id
      maxDelta: 25m0s
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

### Code gen

```sh
make codegen
```

### Running and building local

see `Makefile`
