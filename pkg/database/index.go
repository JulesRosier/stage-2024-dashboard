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

// AND events.index_bikestation IS NULL;

func (q *Queries) FullIndex(ctx context.Context, config EventKeyConfig) error {
	jsonSelect := ""
	l := len(config.KeySelector) - 1
	for i, s := range config.KeySelector {
		if l == i {
			jsonSelect += "->>"
		} else {
			jsonSelect += "->"
		}
		jsonSelect += "'" + s + "'"
	}
	query := fmt.Sprintf(Index, config.IndexColumn, jsonSelect, config.TopicName)
	slog.Info("Indexing", "config_id", config.ID)
	_, err := q.db.Exec(ctx, query)
	return err
}
