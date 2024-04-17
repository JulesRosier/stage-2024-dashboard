package handler

import (
	"Stage-2024-dashboard/pkg/broadcast"
	"Stage-2024-dashboard/pkg/database"
)

type Handler struct {
	Q       *database.Queries
	EventBr *broadcast.BroadcastServer[database.Event]
}

func NewHandler(q *database.Queries, eventBr *broadcast.BroadcastServer[database.Event]) *Handler {
	return &Handler{
		Q:       q,
		EventBr: eventBr,
	}
}
