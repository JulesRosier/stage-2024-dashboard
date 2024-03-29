package indexing

import (
	"Stage-2024-dashboard/pkg/database"
	"context"
)

const IndexPrefix = "index_"

func FullIndex(ctx context.Context) error {
	q := database.GetQueries()
	configs, err := q.ListEventIndexConfigs(ctx)
	if err != nil {
		return err
	}
	for _, config := range configs {
		err := q.AddColumn(ctx, IndexPrefix+config.IndexColumn)
		if err != nil {
			return err
		}
		err = q.FullIndex(ctx, config)
		if err != nil {
			return err
		}

	}
	return nil
}

func PartialIndex(ctx context.Context) error {
	q := database.GetQueries()
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
