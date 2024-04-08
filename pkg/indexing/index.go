package indexing

import (
	"Stage-2024-dashboard/pkg/database"
	"context"
	"log/slog"
)

const IndexPrefix = "index_"

func FullIndex(ctx context.Context, q *database.Queries) error {
	err := CreateIndexColumns(ctx, q)
	if err != nil {
		return err
	}

	configs, err := q.ListEventIndexConfigs(ctx)
	if err != nil {
		return err
	}

	byTopic := map[string][]database.EventIndexConfig{}
	for _, config := range configs {
		byTopic[config.TopicName] = append(byTopic[config.TopicName], config)
	}

	timestampConfigs, err := q.ListTimestampConfigs(ctx)
	if err != nil {
		return err
	}
	timestampByTopic := map[string]database.TimestampConfig{}
	for _, config := range timestampConfigs {
		timestampByTopic[config.TopicName] = config
	}

	for topic, configs := range byTopic {
		err := q.FullIndex(ctx, configs, timestampByTopic[topic])
		if err != nil {
			slog.Warn("Index failed", "topic", topic, "error", err)
		}
	}

	return nil
}

func PartialIndex(ctx context.Context, q *database.Queries) error {
	configs, err := q.ListEventIndexConfigs(ctx)
	if err != nil {
		return err
	}
	for _, config := range configs {
		err = q.PartialIndex(ctx, config)
		if err != nil {
			return err
		}

	}
	return nil
}

func TimestampIndex(ctx context.Context, q *database.Queries) error {
	configs, err := q.ListTimestampConfigs(ctx)
	if err != nil {
		return err
	}
	for _, config := range configs {
		err := q.TimestampIndex(ctx, config)
		if err != nil {
			return err
		}
	}
	return nil
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
