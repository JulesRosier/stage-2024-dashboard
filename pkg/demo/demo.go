package demo

import (
	"Stage-2024-dashboard/pkg/settings"
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type User struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Bike struct {
	Company string `json:"company"`
	Id      string `json:"id"`
}

type UserCreated struct {
	User      User      `json:"user"`
	Timestamp time.Time `json:"timeStamp"`
}

type BikePickedUp struct {
	User      User      `json:"user"`
	Bike      Bike      `json:"bike"`
	Timestamp time.Time `json:"timeStamp"`
}

type BikeReturned struct {
	User      User      `json:"user"`
	Bike      Bike      `json:"bike"`
	Timestamp time.Time `json:"timeStamp"`
}

var client *kgo.Client

func Init(set settings.Kafka) {
	slog.Info("Starting kafka client", "seedbrokers", set.Brokers)

	cl, err := kgo.NewClient(
		kgo.SeedBrokers(set.Brokers...),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		panic(err)
	}

	client = cl
}

func ProduceCreateUser(ctx context.Context, user User) error {
	u := UserCreated{
		User:      user,
		Timestamp: time.Now(),
	}
	eventBtyes, err := json.Marshal(u)
	if err != nil {
		return err
	}
	client.ProduceSync(ctx, &kgo.Record{
		Key:   []byte("demo"),
		Value: eventBtyes,
		Topic: "demo_user_created",
	})
	return nil
}

func ProduceBikePickedUp(ctx context.Context, user User, bike Bike) error {
	b := BikePickedUp{
		User:      user,
		Bike:      bike,
		Timestamp: time.Now(),
	}
	eventBtyes, err := json.Marshal(b)
	if err != nil {
		return err
	}
	client.ProduceSync(ctx, &kgo.Record{
		Key:   []byte("demo"),
		Value: eventBtyes,
		Topic: "demo_bike_picked_up",
	})
	return nil
}

func ProduceBikeReturned(ctx context.Context, user User, bike Bike) error {
	e := BikeReturned{
		User:      user,
		Bike:      bike,
		Timestamp: time.Now(),
	}
	eventBtyes, err := json.Marshal(e)
	if err != nil {
		return err
	}
	client.ProduceSync(ctx, &kgo.Record{
		Key:   []byte("demo"),
		Value: eventBtyes,
		Topic: "demo_bike_returned",
	})
	return nil

}
