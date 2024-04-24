package main

import (
	"Stage-2024-dashboard/pkg/broadcast"
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/demo"
	"Stage-2024-dashboard/pkg/handler"
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/kafka"
	"Stage-2024-dashboard/pkg/server"
	"Stage-2024-dashboard/pkg/settings"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
)

const banner = `
  ______               _    __      ___                        
 |  ____|             | |   \ \    / (_)                       
 | |____   _____ _ __ | |_   \ \  / / _  _____      _____ _ __ 
 |  __\ \ / / _ \ '_ \| __|   \ \/ / | |/ _ \ \ /\ / / _ \ '__|
 | |___\ V /  __/ | | | |_     \  /  | |  __/\ V  V /  __/ |   
 |______\_/ \___|_| |_|\__|     \/   |_|\___| \_/\_/ \___|_|
 > https://github.com/JulesRosier/stage-2024-dashboard
 =============================================================================
`

func main() {
	fmt.Print(banner)
	slog.SetDefault(slog.New(slog.Default().Handler()))

	set, err := settings.Load()
	helper.MaybeDie(err, "Failed to load configs")

	eventStream := make(chan database.Event, 10)

	ctx, cancel := context.WithCancel(context.Background())
	eventBr := broadcast.NewBroadcastServer(ctx, eventStream)

	q := database.NewQueries(set.Database)
	h := handler.NewHandler(q, &eventBr)

	server := server.NewServer(set.Server)
	server.RegisterRoutes(h)
	server.ApplyMiddleware()

	demo.Init(set.Kafka)

	server.Start()
	go kafka.EventImporter(q, eventStream, set.Kafka)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Received an interrupt signal, exiting...")

	cancel()
	server.Stop()
}
