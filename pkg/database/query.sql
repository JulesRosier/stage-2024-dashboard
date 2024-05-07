-- name: CreateEvent :one
INSERT INTO events (
    eventhub_timestamp, topic_name, topic_offset,
    topic_partition, event_headers, event_key, event_value, schema_format, schema_id, event_type
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: CreateEventIndexConfig :one
INSERT INTO event_index_configs (
    event_type, key_selector, index_column
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
SET event_type = $2,
  index_column = $3,
  key_selector = $4
WHERE id = $1
RETURNING *;

-- name: GetIndexColumns :many
SELECT column_name::text
FROM information_schema.columns
WHERE table_name   = 'events'
and column_name like 'index_%';

-- name: CreateTimestampConfig :one
INSERT INTO timestamp_configs (
    event_type, key_selector 
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
SET event_type = $2,
  key_selector = $3
WHERE id = $1
RETURNING *;

-- name: ListAllEventTypeNames :many
SELECT DISTINCT event_type
FROM events;

-- name: ListAllEventTypes :many
SELECT DISTINCT ON (e.event_type) e.*
FROM events e;

-- name: GetEachEventTypeWithNoTimestampConfig :many
SELECT DISTINCT ON (e.event_type) e.*
FROM timestamp_configs tc
right join events e on tc.event_type = e.event_type
where key_selector is null;

-- name: GetIndexColumnsFromConfigs :many
select distinct index_column
from event_index_configs;

-- name: GetConfigStats :many
with events as (
	SELECT DISTINCT ON (e.event_type) e.event_type
FROM events e
)
select text(min(e.event_type)) as topic, count(ec.*) as config_count,
	case
		when count(tc.*) > 0 then 1
		else 0
	end as has_time_config
from event_index_configs ec
right join events e on e.event_type = ec.event_type
left join timestamp_configs tc on tc.event_type = e.event_type
group by e.event_type
order by min(e.event_type);

-- name: GetRandomEvent :one
SELECT * FROM events
WHERE event_type = $1
ORDER BY random() ASC LIMIT 1;