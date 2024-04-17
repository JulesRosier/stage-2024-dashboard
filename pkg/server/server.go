package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
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
	e.HideBanner = true
	e.HidePort = true
	NewServer := &Server{
		// port: port,
		e: e,
	}

	return NewServer
}

// Starts the server in a new routine
func (s *Server) Start() {
	port := ":3000"
	slog.Info("Starting server")
	bind := os.Getenv("HOST")
	if bind == "" {
		bind = "127.0.0.1"
	}
	go func() {
		if err := s.e.Start(bind + port); err != nil && err != http.ErrServerClosed {
			slog.Error("Shutting down the server", "error", err.Error())
		}
	}()
	slog.Info("Server started", "bind", bind, "port", port)
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
