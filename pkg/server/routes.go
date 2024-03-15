package server

import (
	"Stage-2024-dashboard/pkg/handlers"
)

func (s *Server) RegisterRoutes() {
	e := &s.e

	e.Static("/static", "./static")
	e.GET("/", handlers.HelloWebHandler)

}
