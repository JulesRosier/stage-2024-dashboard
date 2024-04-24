package kafka

import (
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/settings"
	"log/slog"
	"time"

	"github.com/redpanda-data/console/backend/pkg/config"
	"github.com/redpanda-data/console/backend/pkg/msgpack"
	"github.com/redpanda-data/console/backend/pkg/proto"
	"github.com/redpanda-data/console/backend/pkg/schema"
	"github.com/redpanda-data/console/backend/pkg/serde"
	"go.uber.org/zap"
)

// Creates  serde.Service and all other services that are need to start it.
func CreateSerde(set settings.Kafka) *serde.Service {
	slog.Info("Creating serde service", "registry", set.SchemaRgistry.Urls)

	logger := zap.L()

	urls := []string{}
	for _, url := range set.SchemaRgistry.Urls {
		urls = append(urls, "http://"+url)
	}
	schemaService, err := schema.NewService(
		config.Schema{
			Enabled:  true,
			URLs:     urls,
			Username: set.Auth.User,
			Password: set.Auth.Password,
		},
		logger,
	)
	helper.MaybeDie(err, "Failed to create schema service")

	protoService, err := proto.NewService(
		config.Proto{
			Enabled: true,
			SchemaRegistry: config.ProtoSchemaRegistry{
				Enabled:         true,
				RefreshInterval: time.Minute * 1,
			},
		},
		logger,
		schemaService,
	)
	helper.MaybeDie(err, "Failed to create proto service")

	err = protoService.Start()
	helper.MaybeDie(err, "Failed to start proto service")

	seserv := serde.NewService(
		schemaService,
		protoService,
		&msgpack.Service{},
	)

	return seserv
}
