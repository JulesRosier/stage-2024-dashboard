package database

import (
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/settings"
	"context"
	"fmt"
	"log/slog"

	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	//go:embed schema.sql
	ddl                string
	checkDatabaseQuery = `SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = $1;`
)

func NewQueries(set settings.Database) *Queries {
	ctx := context.Background()

	DbUser := set.User
	DbPassword := set.Paddword
	DbDatabase := set.Database
	DbHost := set.Host
	DbPort := set.Port

	slog.Info("Starting database", "host", DbHost)
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d database=postgres sslmode=disable",
		DbUser, DbPassword, DbHost, DbPort)

	dbTemp, err := pgxpool.New(ctx, connStr)
	if err != nil {
		helper.DieMsg(err, "Database err")
	}

	slog.Info("Checking if database exists", "database", DbDatabase)
	r := dbTemp.QueryRow(ctx, checkDatabaseQuery, DbDatabase)
	var dbName string
	r.Scan(&dbName)
	if dbName != DbDatabase {
		slog.Info("Database not found, creating", "database", DbDatabase)
		_, err = dbTemp.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", DbDatabase))
		helper.MaybeDie(err, "Failed to create database")
	} else {
		slog.Info("Database found", "database", DbDatabase)
	}

	slog.Info("Starting database", "host", DbHost, "database", DbDatabase)
	connStr = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		DbUser, DbPassword, DbDatabase, DbHost, DbPort)

	db, err := pgxpool.New(ctx, connStr)
	if err != nil {
		helper.DieMsg(err, "Database err")
	}

	// create tables
	if _, err := db.Exec(ctx, ddl); err != nil {
		panic(err)
	}

	queries := New(db)

	return queries
}
