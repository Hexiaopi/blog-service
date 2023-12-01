package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Init() error {
	v := viper.New()
	v.SetConfigFile("configs/app.yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("read config file err:%v", err)
		}
	}
	if err := v.Unmarshal(&AppEngine); err != nil {
		return fmt.Errorf("unmarshal config err:%v", err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&AppEngine); err != nil {
			panic(fmt.Errorf("unmarshal config err:%v", err))
		}
	})

	AppEngine.Flags(pflag.CommandLine)

	pflag.Parse()
	v.BindPFlags(pflag.CommandLine)

	Logger = AppEngine.Log.NewLog()
	var err error
	Tracer, _, err = NewJaegerTracer("blog-service", AppEngine.TraceAgent)
	if err != nil {
		return fmt.Errorf("init jaeger client err:%v", err)
	}

	DBEngine, err = AppEngine.MySQL.NewClient()
	if err != nil {
		return fmt.Errorf("init db engine fail err:%v", err)
	}
	RedisEngine, err = AppEngine.Redis.NewClient()
	if err != nil {
		return fmt.Errorf("init redis engine fail err:%v", err)
	}
	return nil
}
