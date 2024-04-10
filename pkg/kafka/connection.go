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
	group := os.Getenv("CONSUMER_GROUP")

	slog.Info("Starting kafka client", "seedbrokers", seed)

	var cl *kgo.Client
	var err error
	if group == "" {
		cl, err = kgo.NewClient(
			kgo.SeedBrokers(seed),
			kgo.ConsumeTopics("^[A-Za-z].*$"),
			kgo.ConsumeRegex(),
		)
	} else {
		cl, err = kgo.NewClient(
			kgo.SeedBrokers(seed),
			kgo.ConsumeTopics("^[A-Za-z].*$"),
			kgo.ConsumeRegex(),
			kgo.ConsumerGroup(group),
		)
	}
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
