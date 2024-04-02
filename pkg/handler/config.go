package handler

import (
	"Stage-2024-dashboard/pkg/config"
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/view"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func EventIndexConfigHome(c echo.Context) error {
	q := database.GetQueries()
	topics, err := q.ListAllTopics(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.ConfigHome(topics))
}

func EventIndexConfigList(c echo.Context) error {
	q := database.GetQueries()
	configs, err := q.ListEventIndexConfigs(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.ListEventIndexConfigs(configs))
}

func EventIndexConfigCreate(c echo.Context) error {
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	column := strings.TrimSpace(c.FormValue("column"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	q := database.GetQueries()
	_, err := q.CreateEventIndexConfig(c.Request().Context(),
		database.CreateEventIndexConfigParams{
			TopicName:   topic,
			KeySelector: strings.Split(keys, ","),
			IndexColumn: column,
		},
	)
	if err != nil {
		return err
	}
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

// ================================
// TIMESTAMP
// ================================

func TimestampConfigList(c echo.Context) error {
	q := database.GetQueries()
	configs, err := q.ListTimestampConfigs(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.ListTimestampConfigs(configs))
}

func TimestampConfigCreate(c echo.Context) error {
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	q := database.GetQueries()
	_, err := q.CreateTimestampConfig(c.Request().Context(), database.CreateTimestampConfigParams{
		TopicName:   topic,
		KeySelector: strings.Split(keys, ","),
	})
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "newTimestampConfig")
	return render(c, view.TimestampConfigCreateForm())
}

func TimestampConfigDelete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	q := database.GetQueries()
	err = q.DeleteTimestampConfigs(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "newTimestampConfig")
	return c.NoContent(http.StatusNoContent)
}

func TimestampConfigEditForm(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	q := database.GetQueries()
	config, err := q.GetTimestampConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.TimestampConfigEditForm(config))
}

func TimestampConfigEdit(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	q := database.GetQueries()
	config, err := q.UpdateTimestampConfig(c.Request().Context(), database.UpdateTimestampConfigParams{
		ID:          int32(id),
		TopicName:   topic,
		KeySelector: strings.Split(keys, ","),
	})
	if err != nil {
		return err
	}
	return render(c, view.TimestampConfig(config))
}

func TimestampConfig(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	q := database.GetQueries()
	config, err := q.GetTimestampConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.TimestampConfig(config))
}

func TimestampConfigAuto(c echo.Context) error {
	err := config.AutoTimestampConfig(c.Request().Context())
	if err != nil {
		return err
	}
	c.Response().Header().Add("HX-Trigger", "newTimestampConfig")
	return c.NoContent(http.StatusNoContent)
}
