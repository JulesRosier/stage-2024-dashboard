package handler

import (
	"Stage-2024-dashboard/pkg/indexing"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FullIndex(c echo.Context) error {
	start := time.Now()
	r, err := indexing.Index(c.Request().Context(), h.Q, true)
	if err != nil {
		slog.Warn("Failed to full index", "error", err)
		return err
	}
	d := time.Since(start)
	return c.String(http.StatusOK, fmt.Sprintf("%v (%d rows effected)", d, r))
}

func (h *Handler) IndexNewEvents(c echo.Context) error {
	start := time.Now()
	r, err := indexing.Index(c.Request().Context(), h.Q, false)
	if err != nil {
		slog.Warn("Failed to full index", "error", err)
		return err
	}
	d := time.Since(start)
	return c.String(http.StatusOK, fmt.Sprintf("%v (%d rows effected)", d, r))
}
