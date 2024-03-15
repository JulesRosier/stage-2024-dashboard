package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

type Server struct {
	// port int
	e echo.Echo
}

func NewServer() *Server {
	NewServer := &Server{
		// port: port,
		e: *echo.New(),
	}

	return NewServer
}

func (s *Server) Start() {

	go func() {
		if err := s.e.Start("127.0.0.1:3000"); err != nil && err != http.ErrServerClosed {
			slog.Error("Shutting down the server", "error", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	}
}
