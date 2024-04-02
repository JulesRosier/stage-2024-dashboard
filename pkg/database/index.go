package database

import (
	"context"
	"fmt"
	"log/slog"
)

const Index = `
UPDATE events
SET index_%s = (e.event_value%s)::VARCHAR
FROM events e
WHERE 
    events.id = e.id 
    AND events.topic_name = '%s';
`

const IndexNew = `
UPDATE events
SET index_%s = (e.event_value%s)::VARCHAR
FROM events e
WHERE 
    events.id = e.id 
    AND events.topic_name = '%s'
	AND events.%s IS NULL;
`

func (q *Queries) FullIndex(ctx context.Context, config EventIndexConfig) error {
	query := fmt.Sprintf(Index, config.IndexColumn, createJsonSelector(config.KeySelector), config.TopicName)
	slog.Info("Indexing", "config_id", config.ID)
	_, err := q.db.Exec(ctx, query)
	return err
}

// only index those that not has been indexed jet
func (q *Queries) PartialIndex(ctx context.Context, config EventIndexConfig) error {
	query := fmt.Sprintf(IndexNew, config.IndexColumn, createJsonSelector(config.KeySelector), config.TopicName, config.IndexColumn)

	slog.Info("Indexing", "config_id", config.ID)
	_, err := q.db.Exec(ctx, query)
	return err
}

const TimestampIndex = `
UPDATE events
SET event_timestamp = TO_TIMESTAMP(SUBSTRING(e.event_value%s FROM 1 FOR 26), 'YYYY-MM-DD"T"HH24:MI:SS.USZ')
FROM events e
WHERE
	events.id = e.id
	and events.topic_name = '%s';
`

func (q *Queries) TimestampIndex(ctx context.Context, config TimestampConfig) error {
	query := fmt.Sprintf(TimestampIndex, createJsonSelector(config.KeySelector), config.TopicName)
	slog.Info("Indexing timestamps", "config_id", config.ID)
	_, err := q.db.Exec(ctx, query)
	return err
}
func createJsonSelector(qs []string) string {
	jsonSelect := ""
	l := len(qs) - 1
	for i, s := range qs {
		if l == i {
			jsonSelect += "->>"
		} else {
			jsonSelect += "->"
		}
		jsonSelect += "'" + s + "'"
	}
	return jsonSelect
}
