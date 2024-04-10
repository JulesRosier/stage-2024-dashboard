package database

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"strings"
)

const Index = `
UPDATE events
SET last_indexed_at = current_timestamp, %s
FROM events e
WHERE 
    events.id = e.id 
    AND events.topic_name = '%s';
`

const setFromJson = `
index_%s = (e.event_value%s)::VARCHAR
`

const setFromJsonTimestamp = `
event_timestamp = TO_TIMESTAMP(SUBSTRING(e.event_value%s FROM 1 FOR 26), 'YYYY-MM-DD"T"HH24:MI:SS.USZ')
`

const IndexNew = `
UPDATE events
SET index_%s = (e.event_value%s)::VARCHAR
FROM events e
WHERE 
    events.id = e.id 
    AND events.topic_name = '%s'
	AND events.index_%s IS NULL;
`

// TimestampConfig is optional
func (q *Queries) FullIndex(ctx context.Context, configs []EventIndexConfig, timestampConfig TimestampConfig) (int64, error) {
	sets := strings.Builder{}
	l := len(configs)
	for i, config := range configs {
		sets.WriteString(fmt.Sprintf(setFromJson, config.IndexColumn, createJsonSelector(config.KeySelector)))
		if l-1 > i {
			sets.WriteString(",")
		}
	}
	indexTimestamp := !reflect.ValueOf(timestampConfig).IsZero()
	if indexTimestamp {
		sets.WriteString(",")
		sets.WriteString(fmt.Sprintf(setFromJsonTimestamp, createJsonSelector(timestampConfig.KeySelector)))
	}

	query := fmt.Sprintf(Index, sets.String(), configs[0].TopicName)
	slog.Info("Indexing", "topic", configs[0].TopicName, "timestamp?", indexTimestamp)
	r, err := q.db.Exec(ctx, query)
	return r.RowsAffected(), err
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
