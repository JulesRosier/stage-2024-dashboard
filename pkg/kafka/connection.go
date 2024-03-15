package kafka

import (
	"Stage-2024-dashboard/pkg/helper"
	"flag"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sr"
)

func GetClient() *kgo.Client {
	seed := flag.String("seedbroker", "localhost:19092", "brokers port to talk to")
	// topic := flag.String("topic", "bolt-test", "topic to produce to and consume from")

	slog.Info("Starting kafka client", "seedbrokers", *seed)
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(*seed),
		// kgo.ConsumeTopics("baqme-locations"),
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
	registry := flag.String("registry", "localhost:18081", "schema registry port to talk to")
	slog.Info("starting schema registry client", "host", *registry)
	rcl, err := sr.NewClient(sr.URLs(*registry))
	helper.MaybeDieErr(err)
	return rcl
}
