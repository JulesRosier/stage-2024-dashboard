package kafka

import (
	"Stage-2024-dashboard/pkg/helper"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
)

func GetClient() *kgo.Client {
	seed := os.Getenv("SEED_BROKER")

	slog.Info("Starting kafka client", "seedbrokers", seed)
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seed),
		kgo.ConsumeTopics(".*-locations$"),
		kgo.ConsumeRegex(),
		kgo.ConsumerGroup("Testing"),
	)
	if err != nil {
		panic(err)
	}

	return cl
}

func GetRepoClient() *sr.Client {
	registry := os.Getenv("REGISTRY")

	slog.Info("starting schema registry client", "host", registry)
	rcl, err := sr.NewClient(sr.URLs(registry))
	helper.MaybeDieErr(err)
	return rcl
}
