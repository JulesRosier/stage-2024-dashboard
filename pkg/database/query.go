package database

import (
	"context"
	"fmt"
)

const querySearch = `
select id, inserted_at, eventhub_timestamp, event_timestamp, topic_name, topic_offset, topic_partition, event_headers, event_key, event_value
from events
where %s = $1
order by event_timestamp desc
limit %d;
`

func (q *Queries) QuearySearch(ctx context.Context, column string, key string) ([]Event, error) {
	query := fmt.Sprintf(querySearch, column, 10)
	rows, err := q.db.Query(ctx, query, key)
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
