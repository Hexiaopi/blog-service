package config

import "time"

type JWTConfig struct {
	Secret string        `mapstructure:"secret"`
	Issuer string        `mapstructure:"issuer"`
	Expire time.Duration `mapstructure:"expire"`
}
