package database

import (
	"context"
	"fmt"
	"log/slog"
)

const AddColumn = `ALTER TABLE events ADD COLUMN %s VARCHAR(255)`

func (q *Queries) AddColumn(ctx context.Context, columnName string) error {
	query := fmt.Sprintf(AddColumn, columnName)
	slog.Info("Adding Column", "column_name", columnName)
	_, err := q.db.Exec(ctx, query)
	return err
}
