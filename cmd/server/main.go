package main

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/handler"
	"Stage-2024-dashboard/pkg/kafka"
	"Stage-2024-dashboard/pkg/server"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	slog.SetDefault(slog.New(slog.Default().Handler()))

	q := database.NewQueries()
	h := handler.NewHandler(q)

	server := server.NewServer()
	server.RegisterRoutes(h)
	server.ApplyMiddleware()

	server.Start()
	go kafka.EventImporter(q)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Received an interrupt signal, exiting...")

	server.Stop()
}
