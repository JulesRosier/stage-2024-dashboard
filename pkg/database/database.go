package database

import (
	"Stage-2024-dashboard/pkg/helper"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	_ "embed"
)

var (
	//go:embed schema.sql
	ddl     string
	queries *Queries
)

func Init() *Queries {
	ctx := context.Background()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		"postgres", "password", "testing", "localhost", 5432)

	db, err := pgx.Connect(ctx, connStr)
	if err != nil {
		helper.DieMsg("Database err", err)
	}

	// create tables
	if _, err := db.Exec(ctx, ddl); err != nil {
		panic(err)
	}
	queries = New(db)
	return queries
}

func GetQueries() *Queries {
	return queries
}
