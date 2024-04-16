package handler

import (
	"Stage-2024-dashboard/pkg/view"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Home(c echo.Context) error {
	return render(c, view.Home())
}

func (h *Handler) ConfigStats(c echo.Context) error {
	stats, err := h.Q.GetConfigStats(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.ConfigStats(stats))
}
