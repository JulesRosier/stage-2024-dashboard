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
	%s
    events.id = e.id 
    AND events.event_type = '%s';
`

const onlyNewQuery = `
events.last_indexed_at IS NULL AND
`

const setFromJson = `
index_%s = (e.event_value%s)::VARCHAR
`

const setFromJsonTimestamp = `
event_timestamp = TO_TIMESTAMP(SUBSTRING(e.event_value%s FROM 1 FOR 26), 'YYYY-MM-DD"T"HH24:MI:SS.USZ')
`

// TimestampConfig is optional
func (q *Queries) FullIndex(ctx context.Context, configs []EventIndexConfig, timestampConfig TimestampConfig, full bool) (int64, error) {
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

	var new string
	if full {
		new = ""
	} else {
		new = onlyNewQuery
	}
	query := fmt.Sprintf(Index, sets.String(), new, configs[0].EventType)
	slog.Info("Indexing", "event_type", configs[0].EventType, "timestamp?", indexTimestamp)
	r, err := q.db.Exec(ctx, query)
	return r.RowsAffected(), err
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
