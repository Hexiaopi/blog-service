package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	LogLevel        string          `yaml:"LogLevel"`
	ServiceName     string          `yaml:"ServiceName"`
	ServicePort     int             `yaml:"ServicePort"`
	ContextTimeout  time.Duration   `yaml:"ContextTimeout"`
	DefaultPageSize int             `yaml:"DefaultPageSize"`
	MaxPageSize     int             `yaml:"MaxPageSize"`
	DataBase        DatabaseSetting `yaml:"DataBase"`
	JWT             JWTSetting      `yaml:"JWT"`
	TraceAgent      string          `yaml:"TraceAgent"`
}

func NewAppConfig(file string) (*AppConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var app AppConfig
	if err := yaml.Unmarshal(data, &app); err != nil {
		return nil, err
	}
	return &app, nil
}
