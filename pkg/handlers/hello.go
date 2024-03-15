package handlers

import (
	"Stage-2024-dashboard/pkg/view"

	"github.com/labstack/echo/v4"
)

func HelloWebHandler(c echo.Context) error {
	name := "test"
	return render(c, view.HelloPost(name))
}
