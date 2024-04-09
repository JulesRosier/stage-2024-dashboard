package view

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
)

func FormatJson(in []byte) string {
	var data map[string]interface{}
	err := json.Unmarshal(in, &data)
	if err != nil {
		slog.Warn("Failed to unmarshal event", "error", err)
		return "ERROR"
	}
	return renderNode(data)
}

func renderNode(data any) string {
	switch value := data.(type) {
	case map[string]interface{}:
		html := strings.Builder{}
		html.WriteString("<ul>")
		for key, val := range value {
			html.WriteString("<li><span class='json-key'>")
			html.WriteString(key)
			html.WriteString("</span>: ")
			html.WriteString(renderNode(val))
			html.WriteString("</li>")
		}
		html.WriteString("</ul>")
		return html.String()
	case string:
		return fmt.Sprintf("<span class='json-string'>\"%s\"</span>", value)
	case float64:
		return fmt.Sprintf("<span class='json-number'>%f</span>", value)
	case bool:
		return fmt.Sprintf("<span class='json-boolean'>%v</span>", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}
