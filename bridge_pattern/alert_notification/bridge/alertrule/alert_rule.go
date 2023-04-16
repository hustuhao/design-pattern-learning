package alertrule

type Recipients struct {
	Email     []string `yaml:"email,omitempty"`
	SMS       []string `yaml:"sms,omitempty"`
	Slack     []string `yaml:"slack,omitempty"`
	PagerDuty []string `yaml:"pagerduty,omitempty"`
}

type AlertRule struct {
	Name       string              `yaml:"name"`
	Level      string              `yaml:"level"`
	Message    string              `yaml:"message"`
	Recipients map[string][]string `yaml:"recipients"`
}

type Config struct {
	Alerts []AlertRule `yaml:"alerts"`
}

var config Config
