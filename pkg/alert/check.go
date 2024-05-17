package alert

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/settings"
	"context"
	"fmt"
	"log/slog"
	"strings"
)

const layout = "2006-01-02 15:04:05"

func CheckDeltas(set settings.Alert, db *database.Queries) error {
	ctx := context.Background()
	msg := strings.Builder{}
	msg.WriteString("Violations:\n")
	for _, deltaCnf := range set.EventDeltas {
		slog.Info("checking delta", "delta_config", deltaCnf)
		deltas, err := db.CheckDeltas(ctx, deltaCnf)
		if err != nil {
			return err
		}
		if len(deltas) == 0 {
			continue
		}
		msg.WriteString(fmt.Sprintf("\t%s -> %s:", deltaCnf.TopicA, deltaCnf.TopicB))
		msg.WriteString("```")
		data := [][]string{}
		var headers []string
		if len(deltas) > 0 {
			headers = []string{"Delta", "Event ID", "Timestamp"}
		}
		for _, delta := range deltas {
			data = append(data, []string{delta.Delta.String(), delta.Id, delta.Timestamp.Format(layout)})
		}
		msg.WriteString(createASCIITable(headers, data))
		msg.WriteString("```")
	}
	m := msg.String()
	endMsg := " ...```"
	m = m[:3000-len(endMsg)] + endMsg
	err := SendSlackNotification(set, m)
	if err != nil {
		return err
	}
	return nil
}
