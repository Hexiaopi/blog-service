package redis

import (
	"time"
)

type Option func(*Config)

func WithAddr(addr string) Option {
	return func(conf *Config) {
		conf.Addr = addr
	}
}

func WithUserName(username string) Option {
	return func(conf *Config) {
		conf.UserName = username
	}
}

func WithPassWord(password string) Option {
	return func(conf *Config) {
		conf.PassWord = password
	}
}

func WithDB(db int) Option {
	return func(conf *Config) {
		conf.DB = db
	}
}

func WithPoolSize(size int) Option {
	return func(conf *Config) {
		conf.PoolSize = size
	}
}

func WithMinIdleConn(conn int) Option {
	return func(conf *Config) {
		conf.MinIdleConn = conn
	}
}

func WithMinIdleTime(duration time.Duration) Option {
	return func(conf *Config) {
		conf.MinIdleTime = duration
	}
}

func WithMaxConnTime(duration time.Duration) Option {
	return func(conf *Config) {
		conf.MaxConnTime = duration
	}
}
