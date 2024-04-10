package handler

import (
	"Stage-2024-dashboard/pkg/database"
	renderer "Stage-2024-dashboard/pkg/render"
	"Stage-2024-dashboard/pkg/view"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) QueryHome(c echo.Context) error {
	columns, err := h.Q.GetIndexColumns(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.QueryHome(columns))
}

func (h *Handler) QuerySearch(c echo.Context) error {
	var cs []string
	var ss []string
	for name, p := range c.QueryParams() {
		switch name {
		case "column":
			cs = p
		case "search":
			ss = p
		}
	}
	if len(cs) != len(ss) {
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return nil
	}
	ps := []database.QueryParams{}
	for i, c := range cs {
		ps = append(ps, database.QueryParams{
			Column: c,
			Search: ss[i],
		})
	}

	nerdStr := c.QueryParam("nerd_mode")
	nerd := false
	if nerdStr == "on" {
		nerd = true
	}

	layout := "2006-01-02T15:04"
	startStr := c.QueryParam("start")
	start, err := time.Parse(layout, startStr)
	if err != nil {
		return err
	}
	endStr := c.QueryParam("end")
	end, err := time.Parse(layout, endStr)
	if err != nil {
		return err
	}

	slog.Info("aaa", "start", start, "end", end)

	e, err := h.Q.QuearySearch(c.Request().Context(), ps, start, end)
	if err != nil {
		slog.Warn(err.Error())
		return err
	}

	configs, err := h.Q.ListEventIndexConfigs(c.Request().Context())
	if err != nil {
		return err
	}

	byTopic := map[string][]database.EventIndexConfig{}
	for _, config := range configs {
		byTopic[config.TopicName] = append(byTopic[config.TopicName], config)
	}

	events := []view.EventShow{}
	prev := time.Unix(0, 0).Format("2006-01-02")
	for _, event := range e {
		x := false
		d := event.Event.EventTimestamp.Time.Format("2006-01-02")
		if prev != d {
			x = true
			prev = d
		}
		events = append(events, view.EventShow{
			Event:    event.Event,
			ShowDate: x,
			Column:   event.Selected + 2,
			Json:     renderer.FormatJson(event.Event.EventValue, byTopic[event.Event.TopicName]),
		})
	}

	return render(c, view.ListEvents(events, nerd))
}
