package config

import (
	"time"

	"github.com/spf13/pflag"
	"gorm.io/gorm"

	"github.com/hexiaopi/blog-service/pkg/mysql"
)

var DBEngine *gorm.DB

type MySQLConfig struct {
	Host                  string        `mapstructure:"host"`
	Port                  string        `mapstructure:"port"`
	UserName              string        `mapstructure:"username"`
	PassWord              string        `mapstructure:"password"`
	DataBase              string        `mapstructure:"database"`
	Charset               string        `mapstructure:"charset"`
	MaxIdleConnections    int           `mapstructure:"max-idle-connections"`
	MaxOpenConnections    int           `mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `mapstructure:"max-connection-life-time"`
	LogLevel              int           `mapstructure:"log-level"`
}

func (o *MySQLConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.Host, "mysql.host", o.Host, "MySQL service host")
	fs.StringVar(&o.Port, "mysql.port", o.Port, "MySQL service port")
	fs.StringVar(&o.UserName, "mysql.username", o.UserName, "MySQL service username")
	fs.StringVar(&o.PassWord, "mysql.password", o.PassWord, "MySQL service password")
	fs.StringVar(&o.Charset, "mysql.charset", o.Charset, "MySQL service charset")
	fs.IntVar(&o.MaxIdleConnections, "mysql.max-idle-connections", o.MaxIdleConnections, "MySQL max idle connections allowed to connect")
	fs.IntVar(&o.MaxOpenConnections, "mysql.max-open-connections", o.MaxOpenConnections, "MySQL max open connections allowed to connect")
	fs.DurationVar(&o.MaxConnectionLifeTime, "mysql.max-open-connection-life-time", o.MaxConnectionLifeTime, "MySQL max connection life time allowed to connect")
	fs.IntVar(&o.LogLevel, "mysql.log-level", o.LogLevel, "gorm log level")
}

func (o *MySQLConfig) NewClient() (*gorm.DB, error) {
	conf := &mysql.Config{
		Host:                  o.Host,
		Port:                  o.Port,
		UserName:              o.UserName,
		PassWord:              o.PassWord,
		DataBase:              o.DataBase,
		Charset:               o.Charset,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: o.MaxConnectionLifeTime,
		LogLevel:              o.LogLevel,
	}
	return mysql.New(conf, &mysql.OpentracingPlugin{})
}
