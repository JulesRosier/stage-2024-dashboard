package main

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/kafka"
	"Stage-2024-dashboard/pkg/server"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	slog.SetDefault(slog.New(slog.Default().Handler()))

	server := server.NewServer()
	server.RegisterRoutes()
	server.ApplyMiddleware()

	database.Init()

	server.Start()
	go kafka.EventImporter()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Received an interrupt signal, exiting...")

	server.Stop()
}
