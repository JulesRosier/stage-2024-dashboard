package database

import (
	"Stage-2024-dashboard/pkg/settings"
	"context"
	"fmt"
	"time"
)

const checkDeltasQuery = `
with B as (
    select *
    from events
    where event_type = '%s'
), A as (
    select *
    from events
    where event_type = '%s'
)
SELECT A.event_timestamp - B.event_timestamp as delta, A.%s, A.event_timestamp
FROM A
LEFT JOIN (
    SELECT *,
           ROW_NUMBER() OVER (PARTITION BY %s ORDER BY event_timestamp desc) AS rn
    FROM B
) AS B
ON A.%s = B.%s AND B.rn = 1
where A.event_timestamp - B.event_timestamp > '%fh'::INTERVAL AND A.event_timestamp >= NOW() - '%fh'::INTERVAL
order by delta;
`

type EventDelta struct {
	Delta     time.Duration
	Id        string
	Timestamp time.Time
}

func (q *Queries) CheckDeltas(ctx context.Context, ed settings.EventDelta, inter time.Duration) ([]EventDelta, error) {
	query := fmt.Sprintf(checkDeltasQuery, ed.TopicA, ed.TopicB, ed.Index, ed.Index, ed.Index, ed.Index, ed.MaxDelta.Hours(), ed.MaxDelta.Hours()+inter.Hours())
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EventDelta
	for rows.Next() {
		var i EventDelta
		if err := rows.Scan(
			&i.Delta,
			&i.Id,
			&i.Timestamp,
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
