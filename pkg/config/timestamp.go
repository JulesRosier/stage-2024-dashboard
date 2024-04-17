package config

import (
	"Stage-2024-dashboard/pkg/database"
	"context"
	"encoding/json"
	"log/slog"
	"reflect"
	"time"
)

func AutoTimestampConfig(ctx context.Context, q *database.Queries) error {
	events, err := q.GetEachEventTypeWithNoTimestampConfig(ctx)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	for _, event := range events {

		err := json.Unmarshal(event.EventValue, &data)
		if err != nil {
			slog.Warn("Failed to unmarshal event", "error", err)
			continue
		}
		path := checkTimestamps(data, []string{})
		if path == nil {
			slog.Info("No timestamp found", "topic", event.TopicName, "event_id", event.ID)
		}
		slog.Debug("Path found", "topic", event.TopicName, "path", path)
		q.CreateTimestampConfig(ctx, database.CreateTimestampConfigParams{
			TopicName:   event.TopicName,
			KeySelector: path,
		})
	}
	return nil
}

func checkTimestamps(data map[string]interface{}, path []string) []string {
	for key, value := range data {
		a := append(path, key)
		if isTimestamp(value) {
			return a
		} else if reflect.TypeOf(value).Kind() == reflect.Map {
			v := checkTimestamps(value.(map[string]interface{}), a)
			if v != nil {
				return v
			}
		}
	}
	return nil
}

func isTimestamp(value interface{}) bool {
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		time.ANSIC,
		time.DateTime,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RubyDate,
		time.Stamp,
		time.StampMicro,
		time.StampMilli,
		time.StampNano,
		time.Layout,
		time.UnixDate,
	}
	for _, layout := range layouts {
		_, err := time.Parse(layout, strValue)
		if err == nil {
			return true
		}
	}
	return false
}
