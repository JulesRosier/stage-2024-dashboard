package kafka

import (
	"Stage-2024-dashboard/pkg/helper"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kversion"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"github.com/twmb/franz-go/pkg/sr"
)

func GetClient() *kgo.Client {
	seed := os.Getenv("SEED_BROKER")
	group := os.Getenv("CONSUMER_GROUP")
	user := os.Getenv("EH_AUTH_USER")
	pw := os.Getenv("EH_AUTH_PASSWORD")

	opts := []kgo.Opt{
		kgo.SeedBrokers(seed),
		kgo.ConsumeTopics("^[A-Za-z].*$"),
		kgo.ConsumeRegex(),
		kgo.MaxConcurrentFetches(12),
		kgo.MaxVersions(kversion.V2_6_0()),
	}

	if user != "" && pw != "" {
		opts = append(opts, kgo.SASL(scram.Auth{
			User: user,
			Pass: pw,
		}.AsSha512Mechanism()))
	}

	if group != "" {
		opts = append(opts, kgo.ConsumerGroup(group))
	}

	slog.Info("Starting kafka client", "seedbrokers", seed)

	cl, err := kgo.NewClient(
		opts...,
	)
	helper.DieMsg("Failed to create client", err)

	return cl
}

func GetRepoClient() *sr.Client {
	registry := os.Getenv("REGISTRY")
	user := os.Getenv("EH_AUTH_USER")
	pw := os.Getenv("EH_AUTH_PASSWORD")

	slog.Info("starting schema registry client", "host", registry)
	opts := []sr.Opt{
		sr.URLs(registry),
	}
	if user != "" && pw != "" {
		opts = append(opts, sr.BasicAuth(user, pw))
	}
	rcl, err := sr.NewClient(opts...)
	helper.MaybeDieErr(err)
	return rcl
}
