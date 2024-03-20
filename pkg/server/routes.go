package server

import "Stage-2024-dashboard/pkg/handler"

func (s *Server) RegisterRoutes() {
	e := s.e

	e.Static("/static", "./static")
	e.GET("/", handler.HelloWebHandler)

	e.GET("/event_index_config", handler.EventIndexConfig)
	e.POST("/event_index_config", handler.EventIndexConfigCreate)
	e.DELETE("/event_index_config/:id", handler.EventIndexConfigDelete)
	e.GET("/h/event_index_config/list", handler.EventIndexConfigList)
}
