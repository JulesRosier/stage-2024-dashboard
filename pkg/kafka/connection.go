package kafka

import (
	"Stage-2024-dashboard/pkg/helper"
	"Stage-2024-dashboard/pkg/settings"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/kversion"
	"github.com/twmb/franz-go/pkg/sasl/scram"
	"github.com/twmb/franz-go/pkg/sr"
)

func GetClient(set settings.Kafka) *kgo.Client {
	seed := set.Brokers
	group := set.ConsumeGroup
	user := set.Auth.User
	pw := set.Auth.Password

	opts := []kgo.Opt{
		kgo.SeedBrokers(seed...),
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
	helper.MaybeDie(err, "Failed to create client")

	return cl
}

func GetRepoClient(set settings.Kafka) *sr.Client {
	registry := set.SchemaRgistry.Urls
	user := set.Auth.User
	pw := set.Auth.Password

	slog.Info("starting schema registry client", "host", registry)
	opts := []sr.Opt{
		sr.URLs(registry...),
	}
	if user != "" && pw != "" {
		opts = append(opts, sr.BasicAuth(user, pw))
	}
	rcl, err := sr.NewClient(opts...)
	helper.MaybeDieErr(err)
	return rcl
}
