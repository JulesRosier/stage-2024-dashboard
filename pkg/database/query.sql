-- name: CreateEvent :one
INSERT INTO events (
    event_timestamp, topic_name, topic_offset,
    topic_partition, event_headers, event_key, event_value
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: ListTopicEvents : many
select

-- name: CreateEventKeyConfig :one
INSERT INTO event_key_configs (
    topic_name, key_selector, index_column
) VALUES (
 $1, $2, $3
)
RETURNING *;