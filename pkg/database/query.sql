-- name: CreateEvent :one
INSERT INTO events (
    data, topic_name
) VALUES (
  $1, $2
)
RETURNING *;