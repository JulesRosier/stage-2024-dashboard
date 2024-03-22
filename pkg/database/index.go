package database

import (
	"context"
	"fmt"
	"log/slog"
)

const Index = `
UPDATE events
SET %s = (e.event_value%s)::VARCHAR
FROM events e
WHERE 
    events.id = e.id 
    AND events.topic_name = '%s';
`

const IndexNew = `
UPDATE events
SET %s = (e.event_value%s)::VARCHAR
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
