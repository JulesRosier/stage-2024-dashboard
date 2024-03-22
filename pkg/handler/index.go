package handler

import (
	"Stage-2024-dashboard/pkg/indexing"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func FullIndex(c echo.Context) error {
	start := time.Now()
	indexing.FullIndex(c.Request().Context())
	d := time.Since(start)
	return c.String(http.StatusOK, d.String())
}

func IndexNewEvents(c echo.Context) error {
	start := time.Now()
	indexing.PartialIndex(c.Request().Context())
	d := time.Since(start)
	return c.String(http.StatusOK, d.String())
}
