package kafka

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/settings"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redpanda-data/console/backend/pkg/serde"
)

type stringHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func EventImporter(q *database.Queries, eventStream chan database.Event, set settings.Kafka) {
	cl := GetClient(set)
	s := CreateSerde(set)

	ctx := context.Background()

	slog.Info("Waiting for events...")
	for {
		fetches := cl.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			panic(fmt.Sprint(errs))
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()

			sRecord := s.DeserializeRecord(
				ctx,
				record,
				serde.DeserializationOptions{
					KeyEncoding:        serde.PayloadEncodingUnspecified,
					ValueEncoding:      serde.PayloadEncodingUnspecified,
					IgnoreMaxSizeLimit: true,
				},
			)
			if _, ok := sRecord.Value.DeserializedPayload.(map[string]any); !ok {
				slog.Warn("Bad DeserializedPayload",
					"topic", record.Topic,
					"type", reflect.TypeOf(sRecord.Value.DeserializedPayload),
					"payload", fmt.Sprintf("%v", sRecord.Value.DeserializedPayload))
			}

			var err error
			var vb []byte
			if sRecord.Value.DeserializedPayload == nil {
				vb = nil
			} else {
				vb, err = json.Marshal(sRecord.Value.DeserializedPayload)
				if err != nil {
					slog.Warn("Failed to marshal record value to bjson", "err", err)
					continue
				}
			}

			var hb []byte
			if len(record.Headers) == 0 {
				hb = nil
			} else {
				stringHeaders := make([]stringHeader, len(record.Headers)-1)
				for _, h := range record.Headers {
					stringHeaders = append(stringHeaders, stringHeader{h.Key, string(h.Value)})
				}
				hb, err = json.Marshal(stringHeaders)
				if err != nil {
					slog.Warn("Failed to marshal record headers to bjson", "err", err)
					continue
				}
			}

			var kb []byte
			if sRecord.Key.DeserializedPayload == nil {
				kb = nil
			} else {
				kb, err = json.Marshal(sRecord.Key.DeserializedPayload)
				if err != nil {
					slog.Warn("Failed to marshal record key to bjson", "err", err)
					continue
				}
			}
			eventType := "EVENT_TYPE_NOT_FOUND"
			for _, h := range sRecord.Headers {
				if h.Key == "EVENT_TYPE" {
					eventType = string(h.Value)
				}
			}
			id := pgtype.Int4{}
			if sRecord.Value.SchemaID == nil {
				id.Valid = false
			} else {
				x := *sRecord.Value.SchemaID
				id.Int32 = int32(x)
				id.Valid = true
			}

			e, err := q.CreateEvent(ctx, database.CreateEventParams{
				EventhubTimestamp: pgtype.Timestamptz{Time: record.Timestamp, Valid: true},
				TopicName:         record.Topic,
				TopicOffset:       record.Offset,
				TopicPartition:    record.Partition,
				EventHeaders:      hb,
				EventKey:          kb,
				EventValue:        vb,
				SchemaID:          id,
				SchemaFormat:      string(sRecord.Value.Encoding),
				EventType:         eventType,
			})
			eventStream <- e
			if err != nil {
				slog.Warn("Failed to write event to database", "err", err, "topic", record.Topic)
				continue
			}
			slog.Debug("Event saved", "id", e.ID)
		}

	}

}
