package database

import (
	"context"
	"fmt"
	"strings"
	"time"
)

const querySearch = `
select 
id, inserted_at, eventhub_timestamp, event_timestamp, topic_name, topic_offset, topic_partition, event_headers, event_key, event_value, last_indexed_at,
ARRAY [
	%s
] AS selection
from events
where %s
AND event_timestamp BETWEEN '%s' and '%s'
order by event_timestamp desc, id desc
OFFSET %d
LIMIT %d;
`

const queryCase = `
CASE WHEN %s = '%s' THEN true ELSE false END
`

type QueryParams struct {
	Column string
	Search string
}

type QueriedEvent struct {
	Selects []int
	Event   Event
}

func (q *Queries) QuearySearch(
	ctx context.Context,
	qps []QueryParams,
	start time.Time,
	end time.Time,
	offset int,
	limit int,
) ([]QueriedEvent, error) {
	cases := strings.Builder{}
	wheres := strings.Builder{}
	l := len(qps)
	for i, qp := range qps {
		cases.WriteString(fmt.Sprintf(queryCase, qp.Column, qp.Search))
		if l-1 > i {
			cases.WriteString(",")
		}
		wheres.WriteString(fmt.Sprintf(`%s = '%s'`, qp.Column, qp.Search))
		if l-1 > i {
			wheres.WriteString(" or ")
		}
	}

	format := "2006-01-02 15:04:05 -0700"
	query := fmt.Sprintf(querySearch, cases.String(), wheres.String(), start.Format(format), end.Format(format), offset, limit)
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []QueriedEvent
	for rows.Next() {
		var i Event
		var selections []bool
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
			&selections,
		); err != nil {
			return nil, err
		}
		items = append(items, QueriedEvent{
			Selects: boolsToIndices(selections),
			Event:   i,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func boolsToIndices(bools []bool) []int {
	var indices []int
	for i, v := range bools {
		if v {
			indices = append(indices, i)
		}
	}
	return indices
}
