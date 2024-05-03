package handler

import (
	"Stage-2024-dashboard/pkg/database"
	renderer "Stage-2024-dashboard/pkg/render"
	"Stage-2024-dashboard/pkg/view"
	"log/slog"
	"net/http"
	"strconv"
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
	if len(cs) != len(ss) || len(cs) == 0 {
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return nil
	}
	headers := []view.EventHeaders{}
	for i, c := range cs {
		headers = append(headers, view.EventHeaders{
			Qp: database.QueryParams{
				Column: c,
				Search: ss[i],
			},
			Color: colorClasses[i],
		})
	}
	qp := []database.QueryParams{}
	for i, c := range cs {
		qp = append(qp, database.QueryParams{
			Column: c,
			Search: ss[i],
		})
	}

	nerdStr := c.QueryParam("nerd_mode")
	nerd := false
	if nerdStr == "on" {
		nerd = true
	}
	offsetStr := c.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}
	if limit > 500 {
		limit = 500
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

	e, err := h.Q.QuearySearch(c.Request().Context(), qp, start, end, offset, limit)
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
	var query string
	if len(e) == 0 {
		query = ""
	} else {
		q := c.Request().URL.Query()
		q.Set("offset", strconv.Itoa(offset+limit))
		q.Set("limit", strconv.Itoa(100))
		query = q.Encode()
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
		cs := []string{}
		for _, c := range event.Selects[1:] {
			cs = append(cs, headers[c].Color)
		}
		events = append(events, view.EventShow{
			Event:    event.Event,
			ShowDate: x,
			Columns:  event.Selects,
			Json:     renderer.FormatJson(event.Event.EventValue, byTopic[event.Event.TopicName]),
			Colors:   cs,
		})
	}
	return render(c, view.ListEvents(events, headers, nerd, query, offset))
}

var colorClasses = []string{
	"pico-background-pink-450",
	"pico-background-cyan-300",
	"pico-background-violet-450",
	"pico-background-lime-200",
	"pico-background-slate-450",
}
