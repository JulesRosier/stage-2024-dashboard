package server

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func (s *Server) ApplyMiddleware(reRoutes map[string]string) {
	slog.Info("In middleware func", "routes", reRoutes)
	s.e.Pre(echoMw.Rewrite(reRoutes))
	s.e.Use(echoMw.RequestLoggerWithConfig(echoMw.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v echoMw.RequestLoggerValues) error {
			slog.Info("request",
				"method", v.Method,
				"status", v.Status,
				"latency", v.Latency,
				"path", v.URI,
			)
			return nil

		},
	}))
	s.e.Use(echoMw.GzipWithConfig(echoMw.GzipConfig{
		Level: 5,
	}))
}
