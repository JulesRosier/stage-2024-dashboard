package indexing

import (
	"Stage-2024-dashboard/pkg/database"
	"context"
	"log/slog"
)

const IndexPrefix = "index_"

func Index(ctx context.Context, q *database.Queries, full bool) (int64, error) {
	err := CreateIndexColumns(ctx, q)
	if err != nil {
		return 0, err
	}

	configs, err := q.ListEventIndexConfigs(ctx)
	if err != nil {
		return 0, err
	}

	byType := map[string][]database.EventIndexConfig{}
	for _, config := range configs {
		byType[config.EventType] = append(byType[config.EventType], config)
	}

	timestampConfigs, err := q.ListTimestampConfigs(ctx)
	if err != nil {
		return 0, err
	}
	timestampByTopic := map[string]database.TimestampConfig{}
	for _, config := range timestampConfigs {
		timestampByTopic[config.EventType] = config
	}

	var count int64
	for topic, configs := range byType {
		r, err := q.FullIndex(ctx, configs, timestampByTopic[topic], full)
		if err != nil {
			slog.Warn("Index failed", "topic", topic, "error", err)
		}
		count += r
	}

	return count, nil
}

func CreateIndexColumns(ctx context.Context, q *database.Queries) error {
	cs, err := q.GetIndexColumnsFromConfigs(ctx)
	if err != nil {
		return err
	}
	for _, c := range cs {
		err := q.AddColumn(ctx, "index_"+c)
		if err != nil {
			return err
		}
	}
	return nil
}
