package handler

import (
	"Stage-2024-dashboard/pkg/view"
	"bytes"
	"fmt"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EventsLiveHome(c echo.Context) error {
	return render(c, view.EventsLiveHome())
}

func (h *Handler) EventsLiveSSE(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")

	ch := (*h.EventBr).Subscribe()

	ctx := c.Request().Context()
	for {
		select {
		case result, ok := <-ch:
			if !ok {
				return nil
			}
			buf := bytes.NewBufferString("")
			view.LiveEvent(result).Render(ctx, buf)
			fmt.Fprint(c.Response(), buildSSE("message", buf.String()))
			c.Response().Flush()
		case <-ctx.Done():
			(*h.EventBr).CancelSubscription(ch)
			return nil
		}
	}
}
