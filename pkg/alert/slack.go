package alert

import (
	"Stage-2024-dashboard/pkg/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type Payload struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type string   `json:"type"`
	Text Markdown `json:"text"`
}

type Markdown struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// SendSlackNotification sends a POST request with a JSON payload to the specified Slack webhook URL.
func SendSlackNotification(set settings.Alert, message string) error {
	payload := Payload{Blocks: []Block{
		{
			Type: "section",
			Text: Markdown{
				Type: "mrkdwn",
				Text: message,
			},
		},
	}}

	jsonData, err := json.Marshal(payload)
	slog.Debug("Slack msg", "payload", jsonData)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req, err := http.NewRequest("POST", set.Slack.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	return nil
}

// createASCIITable creates an ASCII table from a 2D array of strings.
func createASCIITable(headers []string, data [][]string) string {
	if len(data) == 0 {
		return ""
	}

	data = append([][]string{headers}, data...)

	colWidths := make([]int, len(headers))
	for _, row := range data {
		for colIdx, cell := range row {
			if len(cell) > colWidths[colIdx] {
				colWidths[colIdx] = len(cell)
			}
		}
	}

	createSeparator := func() string {
		var separator strings.Builder
		for _, width := range colWidths {
			separator.WriteString("|-")
			separator.WriteString(strings.Repeat("-", width))
			separator.WriteString("-")
		}
		separator.WriteString("|\n")
		return separator.String()
	}

	createRow := func(row []string) string {
		var result strings.Builder
		for colIdx, cell := range row {
			result.WriteString("| ")
			result.WriteString(cell)
			padding := colWidths[colIdx] - len(cell)
			result.WriteString(strings.Repeat(" ", padding))
			result.WriteString(" ")
		}
		result.WriteString("|\n")
		return result.String()
	}

	var table strings.Builder

	table.WriteString(createRow(headers))
	table.WriteString(createSeparator())

	for _, row := range data[1:] {
		table.WriteString(createRow(row))
	}

	return table.String()
}