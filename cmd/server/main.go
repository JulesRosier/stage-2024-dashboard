package main

import (
	"Stage-2024-dashboard/pkg/alert"
	"Stage-2024-dashboard/pkg/broadcast"
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/demo"
	"Stage-2024-dashboard/pkg/handler"
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/indexing"
	"Stage-2024-dashboard/pkg/kafka"
	"Stage-2024-dashboard/pkg/logger"
	"Stage-2024-dashboard/pkg/scheduler"
	"Stage-2024-dashboard/pkg/server"
	"Stage-2024-dashboard/pkg/settings"
	"Stage-2024-dashboard/pkg/view"
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

	set, err := settings.Load()
	helper.MaybeDie(err, "Failed to load configs")

	slog.SetDefault(logger.NewLogger(set.Logger))

	eventStream := make(chan database.Event, 10)

	server := server.NewServer(set.Server)
	defer server.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	eventBr := broadcast.NewBroadcastServer(ctx, eventStream)
	defer cancel()

	q := database.NewQueries(set.Database)
	h := handler.NewHandler(q, &eventBr)
	rr := view.HashPublicFS()

	server.ApplyMiddleware(rr)
	server.RegisterRoutes(h)

	demo.Init(set.Kafka)

	server.Start()
	go kafka.EventImporter(q, eventStream, set.Kafka)

	s := scheduler.NewScheduler()
	defer s.Stop()
	s.Schedule(set.Indexing.Interval, func() {
		_, err := indexing.Index(context.Background(), q, false)
		if err != nil {
			slog.Warn("Failed to complete scheduled index")
		}
	})
	s.Schedule(set.Alert.Interval, func() { alert.CheckDeltas(set.Alert, q) })

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Received an interrupt signal, exiting...")

}
