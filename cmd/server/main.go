package main

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/kafka"
	"Stage-2024-dashboard/pkg/serde"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	slog.SetDefault(slog.New(slog.Default().Handler()))

	// server := server.NewServer()
	// server.RegisterRoutes()

	database.Init()
	temp()

	// server.Start()
}

func temp() {
	q := database.GetQueries()
	cl := kafka.GetClient()
	rcl := kafka.GetRepoClient()

	fmt.Println("aaa")

	ctx := context.Background()
	c := make(map[int]*serde.Serde)
	for {
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			panic(fmt.Sprint(errs))
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			j, err := serde.HandleDecode(ctx, rcl, record.Value, c)
			helper.MaybeDieErr(err)
			fmt.Println(string(j))
			q.CreateEvent(ctx, database.CreateEventParams{
				Data:      j,
				TopicName: pgtype.Text{String: record.Topic, Valid: true},
			})
		}

	}
}
