package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// RuleInput ...
type RuleInput struct {
	Type    string
	Content string
}

// FormattedContent ...
func (ri *RuleInput) FormattedContent() string {
	if ri.Type == "COMMAND" {
		return "~" + ri.Content
	}
	return ri.Content
}

// RuleOutputType
const (
	RuleOutputTypeMessage = "MESSAGE"
	RuleOutputTypeTime    = "TIME"

	RuleOutputTypeTimeWeatherCurrent = "WEATHER_CURRENT"
)

// RuleOutput ...
type RuleOutput struct {
	Type string

	// MESSAGE
	Content string

	// Time
	TimeFormat string `yaml:"time-format"`

	// Weather
	DefaultLocation string `yaml:"default-location"`
	Units           string
}

// Rule ...
type Rule struct {
	Input  RuleInput
	Output RuleOutput
}

// Scheduler ...
type Scheduler struct {
	Interval int
	Output   RuleOutput
}

// Config ...
type Config struct {
	Server struct {
		Username       string
		Password       string
		Channel        string
		LimitSendCount float64 `yaml:"limit-send-count"`
	}
	External struct {
		OpenWeatherMap struct {
			APIKey string `yaml:"api-key"`
		} `yaml:"open-weather-map"`
	}
	Rules      []Rule
	Schedulers []Scheduler
}

var config Config

// Load ...
func Load(confPath string) error {
	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &config)
}

// Get ...
func Get() Config {
	return config
}
