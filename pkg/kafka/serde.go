package kafka

import (
	"Stage-2024-dashboard/pkg/helper"
	"log/slog"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/redpanda-data/console/backend/pkg/config"
	"github.com/redpanda-data/console/backend/pkg/msgpack"
	"github.com/redpanda-data/console/backend/pkg/proto"
	"github.com/redpanda-data/console/backend/pkg/schema"
	"github.com/redpanda-data/console/backend/pkg/serde"
	"go.uber.org/zap"
)

// Creates  serde.Service and all other services that are need to start it.
func CreateSerde() *serde.Service {
	registry := "http://" + os.Getenv("REGISTRY")

	slog.Info("Creating serde service", "registry", registry)

	logger := zap.L()

	schemaService, err := schema.NewService(
		config.Schema{
			Enabled: true,
			URLs:    []string{registry},
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
