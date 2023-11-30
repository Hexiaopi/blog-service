package redis

import (
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr     string
	UserName string
	PassWord string
	DB       int
}

func New(conf *Config, hooks ...redis.Hook) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Username: conf.UserName,
		Password: conf.PassWord,
		DB:       conf.DB,
	})
	for _, hook := range hooks {
		rdb.AddHook(hook)
	}
	return rdb, nil
}
