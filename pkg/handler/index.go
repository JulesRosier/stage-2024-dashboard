package handler

import (
	"Stage-2024-dashboard/pkg/indexing"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func FullIndex(c echo.Context) error {
	start := time.Now()
	err := indexing.FullIndex(c.Request().Context())
	if err != nil {
		slog.Warn("Failed to full index", "error", err)
		return err
	}
	d := time.Since(start)
	return c.String(http.StatusOK, d.String())
}

func IndexNewEvents(c echo.Context) error {
	start := time.Now()
	err := indexing.PartialIndex(c.Request().Context())
	if err != nil {
		slog.Warn("Failed to partial index", "error", err)
		return err
	}
	d := time.Since(start)
	return c.String(http.StatusOK, d.String())
}

func IndexTimestamps(c echo.Context) error {
	start := time.Now()
	err := indexing.TimestampIndex(c.Request().Context())
	if err != nil {
		slog.Warn("Failed to index timestamps", "error", err)
		return err
	}
	d := time.Since(start)
	return c.String(http.StatusOK, d.String())
}
