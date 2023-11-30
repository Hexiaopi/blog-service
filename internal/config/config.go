package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Init(cmd *cobra.Command) {
	var err error
	v := viper.New()
	v.SetConfigFile("configs/app.yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("read config file err:%v", err))
		}
	}
	if err := v.Unmarshal(&AppEngine); err != nil {
		panic(fmt.Errorf("unmarshal config err:%v", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := v.Unmarshal(&AppEngine); err != nil {
			panic(fmt.Errorf("unmarshal config err:%v", err))
		}
	})
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		name := f.Name
		if !f.Changed && v.IsSet(name) {
			val := v.Get(name)
			cmd.Flags().Set(name, fmt.Sprintf("%v", val))
		}
	})

	AppEngine.Flags(pflag.CommandLine)

	pflag.Parse()
	v.BindPFlags(pflag.CommandLine)

	Logger = AppEngine.Log.NewLog()
	Tracer, _, err = NewJaegerTracer("blog-service", AppEngine.TraceAgent)
	if err != nil {
		panic(fmt.Errorf("init jaeger client err:%v", err))
	}

	DBEngine, err = AppEngine.MySQL.NewClient()
	if err != nil {
		panic(fmt.Errorf("init db engine fail err:%v", err))
	}
	RedisEngine, err = AppEngine.Redis.NewClient()
	if err != nil {
		panic(fmt.Errorf("init redis engine fail err:%v", err))
	}
}
