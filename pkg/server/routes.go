package server

import "Stage-2024-dashboard/pkg/handler"

func (s *Server) RegisterRoutes() {
	e := s.e

	e.Static("/static", "./static")

	e.GET("/", handler.Home)

	e.POST("/index/full", handler.FullIndex)

	e.GET("/event_index_config", handler.EventIndexConfigHome)
	e.POST("/event_index_config", handler.EventIndexConfigCreate)
	e.GET("/event_index_config/:id", handler.EventIndexConfig)
	e.DELETE("/event_index_config/:id", handler.EventIndexConfigDelete)
	e.PUT("/event_index_config/:id", handler.EventIndexConfigEdit)
	e.GET("/event_index_config/:id/edit", handler.EventIndexConfigEditForm)
	e.GET("/h/event_index_config/list", handler.EventIndexConfigList)
}
