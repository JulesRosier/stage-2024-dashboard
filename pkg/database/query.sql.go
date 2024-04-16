// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
    eventhub_timestamp, topic_name, topic_offset,
    topic_partition, event_headers, event_key, event_value
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, inserted_at, eventhub_timestamp, event_timestamp, topic_name, topic_offset, topic_partition, event_headers, event_key, event_value, last_indexed_at
`

type CreateEventParams struct {
	EventhubTimestamp pgtype.Timestamptz
	TopicName         string
	TopicOffset       int64
	TopicPartition    int32
	EventHeaders      []byte
	EventKey          []byte
	EventValue        []byte
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, createEvent,
		arg.EventhubTimestamp,
		arg.TopicName,
		arg.TopicOffset,
		arg.TopicPartition,
		arg.EventHeaders,
		arg.EventKey,
		arg.EventValue,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.InsertedAt,
		&i.EventhubTimestamp,
		&i.EventTimestamp,
		&i.TopicName,
		&i.TopicOffset,
		&i.TopicPartition,
		&i.EventHeaders,
		&i.EventKey,
		&i.EventValue,
		&i.LastIndexedAt,
	)
	return i, err
}

const createEventIndexConfig = `-- name: CreateEventIndexConfig :one
INSERT INTO event_index_configs (
    topic_name, key_selector, index_column
) VALUES (
 $1, $2, $3
)
RETURNING id, inserted_at, topic_name, key_selector, index_column
`

type CreateEventIndexConfigParams struct {
	TopicName   string
	KeySelector []string
	IndexColumn string
}

func (q *Queries) CreateEventIndexConfig(ctx context.Context, arg CreateEventIndexConfigParams) (EventIndexConfig, error) {
	row := q.db.QueryRow(ctx, createEventIndexConfig, arg.TopicName, arg.KeySelector, arg.IndexColumn)
	var i EventIndexConfig
	err := row.Scan(
		&i.ID,
		&i.InsertedAt,
		&i.TopicName,
		&i.KeySelector,
		&i.IndexColumn,
	)
	return i, err
}

const createTimestampConfig = `-- name: CreateTimestampConfig :one

INSERT INTO timestamp_configs (
    topic_name, key_selector 
) VALUES (
 $1, $2 
)
RETURNING id, inserted_at, topic_name, key_selector
`

type CreateTimestampConfigParams struct {
	TopicName   string
	KeySelector []string
}

// -- name: tesmpppp :many
// select id, inserted_at, event_timestamp, topic_name, topic_offset, topic_partition, event_headers, event_key, event_value
// from events
// where index_bike = '133'
// order by event_timestamp desc;
func (q *Queries) CreateTimestampConfig(ctx context.Context, arg CreateTimestampConfigParams) (TimestampConfig, error) {
	row := q.db.QueryRow(ctx, createTimestampConfig, arg.TopicName, arg.KeySelector)
	var i TimestampConfig
	err := row.Scan(
		&i.ID,
		&i.InsertedAt,
		&i.TopicName,
		&i.KeySelector,
	)
	return i, err
}

const deleteEventIndexConfigs = `-- name: DeleteEventIndexConfigs :exec
DELETE FROM event_index_configs
WHERE id = $1
`

func (q *Queries) DeleteEventIndexConfigs(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteEventIndexConfigs, id)
	return err
}

const deleteTimestampConfigs = `-- name: DeleteTimestampConfigs :exec
DELETE FROM timestamp_configs
WHERE id = $1
`

func (q *Queries) DeleteTimestampConfigs(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteTimestampConfigs, id)
	return err
}

const getConfigStats = `-- name: GetConfigStats :many
with events as (
	SELECT DISTINCT ON (e.topic_name) e.topic_name
FROM events e
)
select text(min(e.topic_name)) as topic, count(ec.*) as config_count,
	case
		when count(tc.*) > 0 then 1
		else 0
	end as has_time_config
from event_index_configs ec
right join events e on e.topic_name = ec.topic_name
left join timestamp_configs tc on tc.topic_name = e.topic_name
group by ec.topic_name
order by min(ec.topic_name)
`

type GetConfigStatsRow struct {
	Topic         string
	ConfigCount   int64
	HasTimeConfig int32
}

func (q *Queries) GetConfigStats(ctx context.Context) ([]GetConfigStatsRow, error) {
	rows, err := q.db.Query(ctx, getConfigStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetConfigStatsRow
	for rows.Next() {
		var i GetConfigStatsRow
		if err := rows.Scan(&i.Topic, &i.ConfigCount, &i.HasTimeConfig); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEachEventTypeWithNoTimestampConfig = `-- name: GetEachEventTypeWithNoTimestampConfig :many
SELECT DISTINCT ON (e.topic_name) e.id, e.inserted_at, e.eventhub_timestamp, e.event_timestamp, e.topic_name, e.topic_offset, e.topic_partition, e.event_headers, e.event_key, e.event_value, e.last_indexed_at
FROM timestamp_configs tc
right join events e on tc.topic_name = e.topic_name
where key_selector is null
`

func (q *Queries) GetEachEventTypeWithNoTimestampConfig(ctx context.Context) ([]Event, error) {
	rows, err := q.db.Query(ctx, getEachEventTypeWithNoTimestampConfig)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.InsertedAt,
			&i.EventhubTimestamp,
			&i.EventTimestamp,
			&i.TopicName,
			&i.TopicOffset,
			&i.TopicPartition,
			&i.EventHeaders,
			&i.EventKey,
			&i.EventValue,
			&i.LastIndexedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEventIndexConfig = `-- name: GetEventIndexConfig :one
SELECT id, inserted_at, topic_name, key_selector, index_column
FROM event_index_configs
WHERE id = $1
`

func (q *Queries) GetEventIndexConfig(ctx context.Context, id int32) (EventIndexConfig, error) {
	row := q.db.QueryRow(ctx, getEventIndexConfig, id)
	var i EventIndexConfig
	err := row.Scan(
		&i.ID,
		&i.InsertedAt,
		&i.TopicName,
		&i.KeySelector,
		&i.IndexColumn,
	)
	return i, err
}

const getIndexColumns = `-- name: GetIndexColumns :many
SELECT column_name::text
FROM information_schema.columns
WHERE table_name   = 'events'
and column_name like 'index_%'
`

func (q *Queries) GetIndexColumns(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, getIndexColumns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var column_name string
		if err := rows.Scan(&column_name); err != nil {
			return nil, err
		}
		items = append(items, column_name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getIndexColumnsFromConfigs = `-- name: GetIndexColumnsFromConfigs :many
select distinct index_column
from event_index_configs
`

func (q *Queries) GetIndexColumnsFromConfigs(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, getIndexColumnsFromConfigs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var index_column string
		if err := rows.Scan(&index_column); err != nil {
			return nil, err
		}
		items = append(items, index_column)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTimestampConfig = `-- name: GetTimestampConfig :one
SELECT id, inserted_at, topic_name, key_selector
FROM timestamp_configs
WHERE id = $1
`

func (q *Queries) GetTimestampConfig(ctx context.Context, id int32) (TimestampConfig, error) {
	row := q.db.QueryRow(ctx, getTimestampConfig, id)
	var i TimestampConfig
	err := row.Scan(
		&i.ID,
		&i.InsertedAt,
		&i.TopicName,
		&i.KeySelector,
	)
	return i, err
}

const listAllTopicNames = `-- name: ListAllTopicNames :many
SELECT DISTINCT topic_name
FROM events
`

func (q *Queries) ListAllTopicNames(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, listAllTopicNames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var topic_name string
		if err := rows.Scan(&topic_name); err != nil {
			return nil, err
		}
		items = append(items, topic_name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllTopics = `-- name: ListAllTopics :many
SELECT DISTINCT ON (e.topic_name) e.id, e.inserted_at, e.eventhub_timestamp, e.event_timestamp, e.topic_name, e.topic_offset, e.topic_partition, e.event_headers, e.event_key, e.event_value, e.last_indexed_at
FROM events e
`

func (q *Queries) ListAllTopics(ctx context.Context) ([]Event, error) {
	rows, err := q.db.Query(ctx, listAllTopics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.InsertedAt,
			&i.EventhubTimestamp,
			&i.EventTimestamp,
			&i.TopicName,
			&i.TopicOffset,
			&i.TopicPartition,
			&i.EventHeaders,
			&i.EventKey,
			&i.EventValue,
			&i.LastIndexedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEventIndexConfigs = `-- name: ListEventIndexConfigs :many
SELECT id, inserted_at, topic_name, key_selector, index_column
FROM event_index_configs
ORDER BY inserted_at desc
`

func (q *Queries) ListEventIndexConfigs(ctx context.Context) ([]EventIndexConfig, error) {
	rows, err := q.db.Query(ctx, listEventIndexConfigs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EventIndexConfig
	for rows.Next() {
		var i EventIndexConfig
		if err := rows.Scan(
			&i.ID,
			&i.InsertedAt,
			&i.TopicName,
			&i.KeySelector,
			&i.IndexColumn,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTimestampConfigs = `-- name: ListTimestampConfigs :many
SELECT id, inserted_at, topic_name, key_selector
FROM timestamp_configs
ORDER BY inserted_at desc
`

func (q *Queries) ListTimestampConfigs(ctx context.Context) ([]TimestampConfig, error) {
	rows, err := q.db.Query(ctx, listTimestampConfigs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TimestampConfig
	for rows.Next() {
		var i TimestampConfig
		if err := rows.Scan(
			&i.ID,
			&i.InsertedAt,
			&i.TopicName,
			&i.KeySelector,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEventIndexConfig = `-- name: UpdateEventIndexConfig :one
UPDATE event_index_configs
SET topic_name = $2,
  index_column = $3,
  key_selector = $4
WHERE id = $1
RETURNING id, inserted_at, topic_name, key_selector, index_column
`

type UpdateEventIndexConfigParams struct {
	ID          int32
	TopicName   string
	IndexColumn string
	KeySelector []string
}

func (q *Queries) UpdateEventIndexConfig(ctx context.Context, arg UpdateEventIndexConfigParams) (EventIndexConfig, error) {
	row := q.db.QueryRow(ctx, updateEventIndexConfig,
		arg.ID,
		arg.TopicName,
		arg.IndexColumn,
		arg.KeySelector,
	)
	var i EventIndexConfig
	err := row.Scan(
		&i.ID,
		&i.InsertedAt,
		&i.TopicName,
		&i.KeySelector,
		&i.IndexColumn,
	)
	return i, err
}

const updateTimestampConfig = `-- name: UpdateTimestampConfig :one
UPDATE timestamp_configs
SET topic_name = $2,
  key_selector = $3
WHERE id = $1
RETURNING id, inserted_at, topic_name, key_selector
`

type UpdateTimestampConfigParams struct {
	ID          int32
	TopicName   string
	KeySelector []string
}

func (q *Queries) UpdateTimestampConfig(ctx context.Context, arg UpdateTimestampConfigParams) (TimestampConfig, error) {
	row := q.db.QueryRow(ctx, updateTimestampConfig, arg.ID, arg.TopicName, arg.KeySelector)
	var i TimestampConfig
	err := row.Scan(
		&i.ID,
		&i.InsertedAt,
		&i.TopicName,
		&i.KeySelector,
	)
	return i, err
}
