package handler

import "Stage-2024-dashboard/pkg/database"

type Handler struct {
	Q *database.Queries
}

func NewHandler(q *database.Queries) *Handler {
	return &Handler{
		Q: q,
	}
}
