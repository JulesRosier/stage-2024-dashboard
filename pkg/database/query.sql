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
FROM event_index_configs;

-- name: DeleteEventIndexConfigs :exec
DELETE FROM event_index_configs
WHERE id = $1;