# Project Stage-2024-dashboard

## TODO

- [ ] Red buttons stil have a blue border

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
docker build . -t  ghcr.io/julesrosier/stage-2024-dashboard:latest
docker push ghcr.io/julesrosier/stage-2024-dashboard:latest
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
sqlc generate
templ generate
```

## Queries

```sql
UPDATE events
SET index_bikestation = (e.event_value->>'stationId')::VARCHAR
FROM events e
WHERE
    events.id = 441304
    AND events.topic_name = 'donkey-locations';
```
