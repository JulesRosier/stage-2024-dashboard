package handler

import (
	"Stage-2024-dashboard/pkg/config"
	"Stage-2024-dashboard/pkg/database"
	renderer "Stage-2024-dashboard/pkg/render"
	"Stage-2024-dashboard/pkg/view"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EventIndexConfigHome(c echo.Context) error {
	topics, err := h.Q.ListAllTopicNames(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.ConfigHome(topics))
}

func (h *Handler) EventIndexConfigList(c echo.Context) error {
	configs, err := h.Q.ListEventIndexConfigs(c.Request().Context())
	if err != nil {
		return err
	}

	byTopic := map[string][]database.EventIndexConfig{}
	for _, config := range configs {
		byTopic[config.TopicName] = append(byTopic[config.TopicName], config)
	}
	keys := []string{}
	for k, _ := range byTopic {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return render(c, view.ListEventIndexConfigs(byTopic, keys))
}

func (h *Handler) EventIndexConfigCreate(c echo.Context) error {
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	column := strings.TrimSpace(c.FormValue("column"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	_, err := h.Q.CreateEventIndexConfig(c.Request().Context(),
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

func (h *Handler) EventIndexConfigDelete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	err = h.Q.DeleteEventIndexConfigs(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "newConfig")
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) EventIndexConfigEditForm(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	config, err := h.Q.GetEventIndexConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.EventIndexConfigEditForm(config))
}

func (h *Handler) EventIndexConfigEdit(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	column := strings.TrimSpace(c.FormValue("column"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	config, err := h.Q.UpdateEventIndexConfig(c.Request().Context(), database.UpdateEventIndexConfigParams{
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

func (h *Handler) EventIndexConfig(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	config, err := h.Q.GetEventIndexConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.EventIndexConfig(config))
}

// ================================
// TIMESTAMP
// ================================

func (h *Handler) TimestampConfigList(c echo.Context) error {
	configs, err := h.Q.ListTimestampConfigs(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.ListTimestampConfigs(configs))
}

func (h *Handler) TimestampConfigCreate(c echo.Context) error {
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	_, err := h.Q.CreateTimestampConfig(c.Request().Context(), database.CreateTimestampConfigParams{
		TopicName:   topic,
		KeySelector: strings.Split(keys, ","),
	})
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "newTimestampConfig")
	return render(c, view.TimestampConfigCreateForm())
}

func (h *Handler) TimestampConfigDelete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	err = h.Q.DeleteTimestampConfigs(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "newTimestampConfig")
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) TimestampConfigEditForm(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	config, err := h.Q.GetTimestampConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.TimestampConfigEditForm(config))
}

func (h *Handler) TimestampConfigEdit(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}
	// TODO: field validation
	topic := strings.TrimSpace(c.FormValue("topic"))
	keys := strings.TrimSpace(c.FormValue("keys"))

	config, err := h.Q.UpdateTimestampConfig(c.Request().Context(), database.UpdateTimestampConfigParams{
		ID:          int32(id),
		TopicName:   topic,
		KeySelector: strings.Split(keys, ","),
	})
	if err != nil {
		return err
	}
	return render(c, view.TimestampConfig(config))
}

func (h *Handler) TimestampConfig(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return echo.NotFoundHandler(c)
	}

	config, err := h.Q.GetTimestampConfig(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return render(c, view.TimestampConfig(config))
}

func (h *Handler) TimestampConfigAuto(c echo.Context) error {
	err := config.AutoTimestampConfig(c.Request().Context(), h.Q)
	if err != nil {
		return err
	}
	c.Response().Header().Add("HX-Trigger", "newTimestampConfig")
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) EventIndexConfigAuto(c echo.Context) error {
	err := config.AutoEventIndexConfig(c.Request().Context(), h.Q)
	if err != nil {
		return err
	}
	c.Response().Header().Add("HX-Trigger", "newConfig")
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) EventExample(c echo.Context) error {
	topic := c.QueryParam("topic")
	if topic == "" {
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return nil
	}
	event, err := h.Q.GetRandomEvent(c.Request().Context(), topic)
	if err != nil {
		return err
	}

	return c.String(200, renderer.RenderJson(event.EventValue))
}
