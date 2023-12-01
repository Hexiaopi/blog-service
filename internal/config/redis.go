package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/pflag"

	redisPkg "github.com/hexiaopi/blog-service/pkg/redis"
)

var RedisEngine *redis.Client

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	UserName string `mapstructure:"username"`
	PassWord string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func (o *RedisConfig) AddFlags(fs *pflag.FlagSet) {
	pflag.StringVar(&o.Addr, "redis.addr", o.Addr, "Redis server addr")
	pflag.StringVar(&o.UserName, "redis.username", o.UserName, "Redis server username")
	pflag.StringVar(&o.PassWord, "redis.password", o.PassWord, "Redis server password")
	pflag.IntVar(&o.DB, "redis.db", o.DB, "Redis server db")
}

func (o *RedisConfig) NewClient() (*redis.Client, error) {
	conf := &redisPkg.Config{
		Addr:     o.Addr,
		UserName: o.UserName,
		PassWord: o.PassWord,
		DB:       o.DB,
	}
	return redisPkg.New(conf)
}
