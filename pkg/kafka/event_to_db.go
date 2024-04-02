package kafka

import (
	"Stage-2024-dashboard/pkg/database"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redpanda-data/console/backend/pkg/serde"
)

func EventImporter() {
	q := database.GetQueries()
	cl := GetClient()
	s := CreateSerde()

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
				hb, err = json.Marshal(record.Headers)
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

			e, err := q.CreateEvent(ctx, database.CreateEventParams{
				EventhubTimestamp: pgtype.Timestamp{Time: record.Timestamp, Valid: true},
				TopicName:         record.Topic,
				TopicOffset:       record.Offset,
				TopicPartition:    record.Partition,
				EventHeaders:      hb,
				EventKey:          kb,
				EventValue:        vb,
			})
			if err != nil {
				slog.Warn("Failed to write event to database", "err", err, "topic", record.Topic)
				continue
			}
			slog.Debug("Event saved", "id", e.ID)
		}

	}

}
