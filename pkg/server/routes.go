package server

import "Stage-2024-dashboard/pkg/handler"

func (s *Server) RegisterRoutes(hdlr *handler.Handler) {
	e := s.e

	e.Static("/static", "./static")

	e.GET("/", hdlr.Home)

	e.POST("/index/full", hdlr.FullIndex)
	e.POST("/index/new", hdlr.IndexNewEvents)

	e.GET("/config_stats", hdlr.ConfigStats)

	e.GET("/config", hdlr.EventIndexConfigHome)
	e.POST("/event_index_config", hdlr.EventIndexConfigCreate)
	e.GET("/event_index_config/:id", hdlr.EventIndexConfig)
	e.DELETE("/event_index_config/:id", hdlr.EventIndexConfigDelete)
	e.PUT("/event_index_config/:id", hdlr.EventIndexConfigEdit)
	e.GET("/event_index_config/:id/edit", hdlr.EventIndexConfigEditForm)
	e.GET("/h/event_index_config/list", hdlr.EventIndexConfigList)

	e.POST("/timestamp_config", hdlr.TimestampConfigCreate)
	e.GET("/timestamp_config/:id", hdlr.TimestampConfig)
	e.DELETE("/timestamp_config/:id", hdlr.TimestampConfigDelete)
	e.PUT("/timestamp_config/:id", hdlr.TimestampConfigEdit)
	e.GET("/timestamp_config/:id/edit", hdlr.TimestampConfigEditForm)
	e.GET("/h/timestamp_config/list", hdlr.TimestampConfigList)

	e.POST("/timestamp_config/auto", hdlr.TimestampConfigAuto)
	e.POST("/event_index_config/auto", hdlr.EventIndexConfigAuto)

	e.GET("/query", hdlr.QueryHome)
	e.GET("/query/search", hdlr.QuerySearch)
}
