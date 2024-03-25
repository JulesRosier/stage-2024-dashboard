-- name: CreateEvent :one
INSERT INTO events (
    event_timestamp, topic_name, topic_offset,
    topic_partition, event_headers, event_key, event_value
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: CreateEventIndexConfig :one
INSERT INTO event_index_configs (
    topic_name, key_selector, index_column
) VALUES (
 $1, $2, $3
)
RETURNING *;

-- name: ListEventIndexConfigs :many
SELECT *
FROM event_index_configs
ORDER BY inserted_at desc;

-- name: DeleteEventIndexConfigs :exec
DELETE FROM event_index_configs
WHERE id = $1;

-- name: GetEventIndexConfig :one
SELECT *
FROM event_index_configs
WHERE id = $1;

-- name: UpdateEventIndexConfig :one
UPDATE event_index_configs
SET topic_name = $2,
  index_column = $3,
  key_selector = $4
WHERE id = $1
RETURNING *;

-- name: GetIndexColumns :many
SELECT column_name::text
FROM information_schema.columns
WHERE table_name   = 'events'
and column_name like 'index_%';

-- -- name: tesmpppp :many
-- select id, inserted_at, event_timestamp, topic_name, topic_offset, topic_partition, event_headers, event_key, event_value
-- from events
-- where index_bike = '133'
-- order by event_timestamp desc;