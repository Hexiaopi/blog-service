package config

import "time"

type JWTConfig struct {
	Secret string        `yaml:"secret"`
	Issuer string        `yaml:"issuer"`
	Expire time.Duration `yaml:"expire"`
}
