package server

import (
	"Stage-2024-dashboard/pkg/settings"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e        *echo.Echo
	settings settings.Server
}

func NewServer(set settings.Server) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	NewServer := &Server{
		e:        e,
		settings: set,
	}

	return NewServer
}

// Starts the server in a new routine
func (s *Server) Start() {
	slog.Info("Starting server")
	address := fmt.Sprintf("%s:%d", s.settings.Bind, s.settings.Port)
	go func() {
		if err := s.e.Start(address); err != nil && err != http.ErrServerClosed {
			slog.Error("Shutting down the server", "error", err.Error())
		}
	}()
	slog.Info("Server started", "address", address)
}

// Tries to the stops the server gracefully
func (s *Server) Stop() {
	slog.Info("Stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	}
}
