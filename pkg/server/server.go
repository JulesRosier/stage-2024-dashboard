package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

type Server struct {
	// port int
	e *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	NewServer := &Server{
		// port: port,
		e: e,
	}
	e.HideBanner = true

	return NewServer
}

// Starts the server in a new routine
func (s *Server) Start() {
	slog.Info("Starting server")
	go func() {
		if err := s.e.Start("127.0.0.1:3000"); err != nil && err != http.ErrServerClosed {
			slog.Error("Shutting down the server", "error", err.Error())
		}
	}()
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
