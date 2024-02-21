package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
)

type Config struct {
	Addr        string
	UserName    string
	PassWord    string
	DB          int
	PoolSize    int
	MinIdleConn int
	MinIdleTime time.Duration
	MaxConnTime time.Duration
	Tracer      opentracing.Tracer
}

func NewClient(conf *Config, hooks ...redis.Hook) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Username:     conf.UserName,
		Password:     conf.PassWord,
		DB:           conf.DB,
		PoolSize:     conf.PoolSize,
		MinIdleConns: conf.MinIdleConn,
		IdleTimeout:  conf.MinIdleTime,
		MaxConnAge:   conf.MaxConnTime,
	})

	for _, hook := range hooks {
		rdb.AddHook(hook)
	}
	return rdb, nil
}

func New(opts ...Option) (*redis.Client, error) {
	conf := &Config{}
	for _, opt := range opts {
		opt(conf)
	}
	hooks := make([]redis.Hook, 0)
	if conf.Tracer != nil {
		hooks = append(hooks, NewHook(conf.Tracer))
	}
	return NewClient(conf, hooks...)
}
