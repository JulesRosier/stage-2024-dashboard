package handler

import (
	"Stage-2024-dashboard/pkg/database"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func FullIndex(c echo.Context) error {
	q := database.GetQueries()
	ctx := context.Background()
	start := time.Now()
	configs, err := q.ListEventIndexConfigs(c.Request().Context())
	if err != nil {
		return err
	}
	for _, config := range configs {
		err := q.AddColumn(ctx, config.IndexColumn)
		if err != nil {
			return err
		}
		err = q.FullIndex(ctx, config)
		if err != nil {
			return err
		}

	}
	d := time.Now().Sub(start)
	return c.String(http.StatusOK, d.String())
}

func IndexNewEvents(c echo.Context) error {
	q := database.GetQueries()
	ctx := context.Background()
	start := time.Now()
	configs, err := q.ListEventIndexConfigs(c.Request().Context())
	if err != nil {
		return err
	}
	for _, config := range configs {
		// err := q.AddColumn(ctx, config.IndexColumn)
		// if err != nil {
		// 	return err
		// }
		err = q.IndexNew(ctx, config)
		if err != nil {
			return err
		}

	}
	d := time.Now().Sub(start)
	return c.String(http.StatusOK, d.String())
}
