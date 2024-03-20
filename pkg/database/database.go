package database

import (
	"Stage-2024-dashboard/pkg/helper"
	"context"
	"fmt"
	"log/slog"
	"os"

	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var (
	//go:embed schema.sql
	ddl     string
	queries *Queries
)

func Init() *Queries {
	ctx := context.Background()

	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbDatabase := os.Getenv("DB_DATABASE")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")

	slog.Info("Starting database", "host", DbHost, "database", DbDatabase)
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DbUser, DbPassword, DbDatabase, DbHost, DbPort)

	// db, err := pgx.Connect(ctx, connStr)
	db, err := pgxpool.New(ctx, connStr)
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
