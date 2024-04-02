package handler

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/view"
	"log/slog"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func QueryHome(c echo.Context) error {
	q := database.GetQueries()
	columns, err := q.GetIndexColumns(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.QueryHome(columns))
}

func QuerySearch(c echo.Context) error {
	column := strings.TrimSpace(c.FormValue("column"))
	search := strings.TrimSpace(c.FormValue("search"))

	q := database.GetQueries()
	e, err := q.QuearySearch(c.Request().Context(), column, search, 20)
	if err != nil {
		slog.Warn(err.Error())
		return err
	}
	ewd := []view.EventWithDate{}
	prev := time.Unix(0, 0).Format("2006-01-02")
	for _, event := range e {
		x := false
		d := event.EventTimestamp.Time.Format("2006-01-02")
		if prev != d {
			x = true
			prev = d
		}
		ewd = append(ewd, view.EventWithDate{
			Event:    event,
			ShowDate: x,
		})
	}

	return render(c, view.ListEvents(ewd))
}
