package renderer

import (
	"Stage-2024-dashboard/pkg/database"
	"Stage-2024-dashboard/pkg/helper"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
)

type linkNode struct {
	node  any
	index string
}

func FormatJson(in []byte, configs []database.EventIndexConfig) string {
	var data map[string]any
	err := json.Unmarshal(in, &data)
	if err != nil {
		slog.Warn("Failed to unmarshal event", "error", err)
		return string(in)
	}

	for _, config := range configs {
		replaceObjectByKeySelectors(data, config.KeySelector, linkNode{node: getObjectByKeySelectors(data, config.KeySelector), index: "index_" + config.IndexColumn})
	}

	return renderNode(data)
}

func RenderJson(in []byte) string {
	var data map[string]any
	err := json.Unmarshal(in, &data)
	if err != nil {
		slog.Warn("Failed to unmarshal event", "error", err)
		return string(in)
	}

	return renderNode(data)
}

func replaceObjectByKeySelectors(data interface{}, keySelectors []string, replacement linkNode) {
	if len(keySelectors) == 0 {
		return
	}

	if dataMap, ok := data.(map[string]interface{}); ok {
		if len(keySelectors) == 1 {
			dataMap[keySelectors[0]] = replacement
			return
		}

		nextLevel := dataMap[keySelectors[0]]
		replaceObjectByKeySelectors(nextLevel, keySelectors[1:], replacement)
	}
}

func getObjectByKeySelectors(data interface{}, keySelectors []string) interface{} {
	current := data
	for _, key := range keySelectors {
		if dataMap, ok := current.(map[string]interface{}); ok {
			if val, exists := dataMap[key]; exists {
				current = val
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
	return current
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
			//TODO: add config to turn map on or off
			loc := false
			if key == "location" && loc {
				html.WriteString(locationmap(val))
			}
		}
		html.WriteString("</ul>")
		return html.String()
	case string:
		return fmt.Sprintf("<span class='json-string'>\"%s\"</span>", value)
	case float64:
		return fmt.Sprintf("<span class='json-number'>%f</span>", value)
	case bool:
		return fmt.Sprintf("<span class='json-boolean'>%v</span>", value)
	case linkNode:
		html := strings.Builder{}
		html.WriteString("<a href='#' onclick='addQuery(`")
		html.WriteString(value.index)
		html.WriteString("`,`")
		w, ok := value.node.(string)
		if ok {
			html.WriteString(w)
		} else {
			slog.Warn("linknode content was not a string", "node", value.node)
			loc, err := json.Marshal(value.node)
			helper.MaybeDieErr(err)
			stringloc := strings.ReplaceAll(string(loc), ":", ": ")
			stringloc = strings.ReplaceAll(stringloc, ",", ", ")
			html.WriteString(stringloc)
			renderNode(value.node)
		}
		html.WriteString("`)'>")
		html.WriteString(renderNode(value.node))
		html.WriteString("</a>")
		return html.String()
	default:
		return fmt.Sprintf("%v", value)
	}
}

func locationmap(val any) string {
	var loc [2]string
	for key, val := range val.(map[string]interface{}) {
		if key == "latitude" {
			loc[0] = fmt.Sprintf("%v", val)
		} else {
			loc[1] = fmt.Sprintf("%v", val)
		}
	}

	return fmt.Sprintf("<div><iframe width=\"100%%<\" height=\"200\" frameborder=\"0\" scrolling=\"no\" marginheight=\"0<\" marginwidth=\"0\"  loading=\"lazy\" src=\"https://maps.google.com/maps?width=100%%25&amp;height=400&amp;hl=en&amp;q=%s,%s+(123)&amp;t=&amp;z=14&amp;ie=UTF8&amp;iwloc=B&amp;output=embed\"></iframe></div>", loc[0], loc[1])
}
