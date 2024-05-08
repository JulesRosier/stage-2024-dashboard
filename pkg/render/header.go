package renderer

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func HeaderTableFormat(in []byte) string {
	h := []header{}
	var out string
	err := json.Unmarshal(in, &h)
	if err != nil {
		slog.Warn("Failed to unmarshal event", "error", err)
		return fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", "ERR", err)
	}
	for _, header := range h {
		out = fmt.Sprintf("%s <tr><td>%s</td><td>%s</td></tr>", out, header.Key, header.Value)
	}

	return out
}
