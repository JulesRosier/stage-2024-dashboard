package alert

import (
	"Stage-2024-dashboard/pkg/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Text string `json:"text"`
}

// SendSlackNotification sends a POST request with a JSON payload to the specified Slack webhook URL.
func SendSlackNotification(set settings.Alert, message string) error {
	// Create the payload
	payload := Payload{Text: message}

	// Marshal the payload into JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", set.Slack.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK response status: %s", resp.Status)
	}

	return nil
}
