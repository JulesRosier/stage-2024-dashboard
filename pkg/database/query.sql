-- name: CreateEvent :one
INSERT INTO events (
    eventhub_timestamp, topic_name, topic_offset,
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

-- name: CreateTimestampConfig :one
INSERT INTO timestamp_configs (
    topic_name, key_selector 
) VALUES (
 $1, $2 
)
RETURNING *;

-- name: ListTimestampConfigs :many
SELECT *
FROM timestamp_configs
ORDER BY inserted_at desc;

-- name: DeleteTimestampConfigs :exec
DELETE FROM timestamp_configs
WHERE id = $1;

-- name: GetTimestampConfig :one
SELECT *
FROM timestamp_configs
WHERE id = $1;

-- name: UpdateTimestampConfig :one
UPDATE timestamp_configs
SET topic_name = $2,
  key_selector = $3
WHERE id = $1
RETURNING *;

-- name: ListAllTopics :many
SELECT DISTINCT topic_name
FROM events;

-- name: GetEachEventTypeWithNoConfig :many
SELECT DISTINCT ON (e.topic_name) e.*
FROM timestamp_configs tc
right join events e on tc.topic_name = e.topic_name
where key_selector is null;