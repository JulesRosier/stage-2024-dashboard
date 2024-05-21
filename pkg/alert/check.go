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
const maxMsgLen = 3000

func CheckDeltas(set settings.Alert, db *database.Queries) error {
	ctx := context.Background()
	for _, deltaCnf := range set.EventDeltas {
		msg := strings.Builder{}
		slog.Info("checking delta", "delta_config", deltaCnf)
		deltas, err := db.CheckDeltas(ctx, deltaCnf, set.Interval)
		if err != nil {
			return err
		}
		if len(deltas) == 0 {
			continue
		}
		msg.WriteString(fmt.Sprintf("\t%s -> %s:", deltaCnf.TopicA, deltaCnf.TopicB))
		msg.WriteString("```")
		data := [][]string{}
		headers := []string{"Delta", "Event ID", "Timestamp"}
		for _, delta := range deltas {
			data = append(data, []string{delta.Delta.String(), delta.Id, delta.Timestamp.Format(layout)})
		}
		msg.WriteString(createASCIITable(headers, data))
		msg.WriteString("```")
		m := msg.String()
		endMsg := "\n...```"
		i := strings.LastIndex(m[:maxMsgLen], "\n")
		room := maxMsgLen - i
		if room < len(endMsg) {
			i = strings.LastIndex(m[:i-len(endMsg)], "\n")
		}
		m = m[:i] + endMsg
		err = SendSlackNotification(set, "Violations", m)
		if err != nil {
			return err
		}
	}
	return nil
}
