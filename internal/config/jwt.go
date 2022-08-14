package config

import "time"

type JWTSetting struct {
	Secret string        `yaml:"Secret"`
	Issuer string        `yaml:"Issuer"`
	Expire time.Duration `yaml:"Expire"`
}
