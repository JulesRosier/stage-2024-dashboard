package settings

type Alert struct {
	Slack struct {
		WebhookURL string `yaml:"webhookURL"`
	} `yaml:"slack"`
}

func (m *Alert) SetDefault() {

}
