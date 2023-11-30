package config

import (
	"github.com/spf13/pflag"

	"github.com/hexiaopi/blog-service/pkg/log"
)

var Logger *log.Logger

type LogConfig struct {
	FileName  string `yaml:"file-name"`
	LogLevel  string `yaml:"log-level"`
	MaxSize   int    `yaml:"max-size"`
	MaxBackup int    `yaml:"max-backup"`
	MaxAge    int    `yaml:"max-age"`
	Compress  bool   `yaml:"compress"`
	Encoding  string `yaml:"encoding"`
	Env       string `yaml:"env"`
}

func (o *LogConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.FileName, "log.file-name", o.FileName, "Log file name")
	fs.StringVar(&o.LogLevel, "log.log-level", o.LogLevel, "Log level")
	fs.IntVar(&o.MaxSize, "log.max-size", o.MaxSize, "Log file max size")
	fs.IntVar(&o.MaxBackup, "log.max-backup", o.MaxBackup, "Log file max backup")
	fs.IntVar(&o.MaxAge, "log.max-age", o.MaxAge, "Log file max age")
	fs.BoolVar(&o.Compress, "log.compress", o.Compress, "Log file compress")
	fs.StringVar(&o.Encoding, "log.encoding", o.Encoding, "Log file encoding")
	fs.StringVar(&o.Env, "log.env", o.Env, "Log env")
}

func (o *LogConfig) NewLog() *log.Logger {
	conf := &log.Config{
		FileName:  o.FileName,
		LogLevel:  o.LogLevel,
		MaxSize:   o.MaxSize,
		MaxBackup: o.MaxBackup,
		MaxAge:    o.MaxAge,
		Compress:  o.Compress,
		Encoding:  o.Encoding,
		Env:       o.Env,
	}
	return log.New(conf)
}
