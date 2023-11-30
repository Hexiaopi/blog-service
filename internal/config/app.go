package config

import (
	"github.com/spf13/pflag"
)

var AppEngine AppConfig

type AppConfig struct {
	MySQL      MySQLConfig `yaml:"mysql"`
	Redis      RedisConfig `yaml:"redis"`
	Log        LogConfig   `yaml:"log"`
	JWT        JWTConfig   `yaml:"jwt"`
	HTTP       HttpConfig  `yaml:"http"`
	TraceAgent string      `yaml:"traceAgent"`
}

func (o *AppConfig) Flags(fs *pflag.FlagSet) {
	o.MySQL.AddFlags(fs)
	o.Redis.AddFlags(fs)
	o.Log.AddFlags(fs)
}
