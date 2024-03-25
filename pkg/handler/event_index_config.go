package handler

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/view"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func EventIndexConfigHome(c echo.Context) error {
	return render(c, view.EventIndexConfigHome())
}

func EventIndexConfigList(c echo.Context) error {
	q := database.GetQueries()
	configs, err := q.ListEventIndexConfigs(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.ListConfigs(configs))
}

func EventIndexConfigCreate(c echo.Context) error {
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	column := strings.TrimSpace(c.FormValue("column"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	q := database.GetQueries()
	q.CreateEventIndexConfig(c.Request().Context(),
		database.CreateEventIndexConfigParams{
			TopicName:   topic,
			KeySelector: strings.Split(keys, ","),
			IndexColumn: column,
		},
	)

	c.Response().Header().Add("HX-Trigger", "newConfig")
	return render(c, view.EventIndexConfigCreateForm())
}

func EventIndexConfigDelete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	q := database.GetQueries()
	err = q.DeleteEventIndexConfigs(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "newConfig")
	return c.NoContent(http.StatusNoContent)
}

func EventIndexConfigEditForm(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	q := database.GetQueries()
	config, err := q.GetEventIndexConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.EventIndexConfigEditForm(config))
}

func EventIndexConfigEdit(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	column := strings.TrimSpace(c.FormValue("column"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	q := database.GetQueries()
	config, err := q.UpdateEventIndexConfig(c.Request().Context(), database.UpdateEventIndexConfigParams{
		ID:          int32(id),
		TopicName:   topic,
		IndexColumn: column,
		KeySelector: strings.Split(keys, ","),
	})
	if err != nil {
		return err
	}

	fmt.Println(topic)
	fmt.Println(column)
	fmt.Println(keys)

	return render(c, view.EventIndexConfig(config))
}

func EventIndexConfig(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	q := database.GetQueries()
	config, err := q.GetEventIndexConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.EventIndexConfig(config))
}
