package handler

import (
	"Stage-2024-dashboard/pkg/view"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return render(c, view.Home())
}
