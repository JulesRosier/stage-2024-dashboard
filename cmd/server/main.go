package main

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/handler"
	"Stage-2024-dashboard/pkg/kafka"
	"Stage-2024-dashboard/pkg/server"
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
