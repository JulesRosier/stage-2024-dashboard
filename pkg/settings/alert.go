package settings

import "time"

type Alert struct {
	Slack struct {
		WebhookURL string `yaml:"webhookURL"`
	} `yaml:"slack"`
	EventDeltas []EventDelta  `yaml:"eventDeltas"`
	Interval    time.Duration `yaml:"interval"`
}

type EventDelta struct {
	TopicA   string        `yaml:"topicA"`
	TopicB   string        `yaml:"topicB"`
	Index    string        `yaml:"index"`
	MaxDelta time.Duration `yaml:"maxDelta"`
}

func (a *Alert) SetDefault() {
	a.Interval = time.Hour
}
