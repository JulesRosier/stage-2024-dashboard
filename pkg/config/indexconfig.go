package config

import (
	"Stage-2024-dashboard/pkg/database"
	"context"
	"encoding/json"
	"log/slog"
	"regexp"
)

func AutoEventIndexConfig(ctx context.Context, q *database.Queries) error {
	events, err := q.ListAllTopics(ctx)
	if err != nil {
		return err
	}
	configs, err := q.ListEventIndexConfigs(ctx)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	for _, event := range events {
		err := json.Unmarshal(event.EventValue, &data)
		if err != nil {
			slog.Warn("Failed to unmarshal event", "error", err)
			return err
		}
		uuidKeys := make(map[string][]string)
		findUUIDKeys(data, []string{}, &uuidKeys)

		for _, path := range uuidKeys {
			allow := true
			for _, config := range configs {
				if config.TopicName == event.TopicName && compareSlice(config.KeySelector, path) {
					allow = false
				}
			}
			if allow {
				indexName := makeIndexName(path)
				q.CreateEventIndexConfig(ctx, database.CreateEventIndexConfigParams{
					TopicName:   event.TopicName,
					IndexColumn: indexName,
					KeySelector: path,
				})
			}
		}
	}
	return nil
}

var uuidPattern = regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`)

func findUUIDKeys(data any, path []string, uuidKeys *map[string][]string) {
	switch value := data.(type) {
	case map[string]any:
		for key, val := range value {
			newPath := append(path, key)
			findUUIDKeys(val, newPath, uuidKeys)
		}
	default:
		if checkUUID(data) {
			(*uuidKeys)[data.(string)] = path
		}
	}
}

func checkUUID(value any) bool {
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	return uuidPattern.MatchString(strValue)
}

func makeIndexName(path []string) string {
	l := len(path)
	if l > 1 {
		return path[l-2] + "_" + path[l-1]
	} else {
		return "unknown_id"
	}
}

func compareSlice(a, b []string) bool {
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}
