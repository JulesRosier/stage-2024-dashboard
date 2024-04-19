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

	e.GET("/events/live", hdlr.EventsLiveHome)
	e.GET("/events/live/sse", hdlr.EventsLiveSSE)

	d := e.Group("/demo")
	d.GET("/home", hdlr.DemoHome)
	d.GET("/step/1", hdlr.DemoStep1)
	d.POST("/step/1", hdlr.DemoStep1Post)
	d.POST("/step/2", hdlr.DemoStep2Post)
	d.POST("/step/3", hdlr.DemoStep3Post)
}
