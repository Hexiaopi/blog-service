package config

import (
	"fmt"
	"time"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DatabaseSetting struct {
	DBType          string        `yaml:"dbType"`
	UserName        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	DBName          string        `yaml:"dbname"`
	Host            string        `yaml:"host"`
	Port            string        `yaml:"port"`
	Charset         string        `yaml:"charset"`
	ParseTime       bool          `yaml:"parseTime"`
	MaxIdleConn     int           `yaml:"maxIdleConn"`
	MaxOpenConn     int           `yaml:"maxOpenConn"`
	ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
}

var TestDataBase = DatabaseSetting{
	DBType:          "mysql",
	UserName:        "root",
	Password:        "",
	Host:            "127.0.0.1",
	Port:            "3306",
	DBName:          "blog_service",
	Charset:         "utf8",
	ParseTime:       true,
	MaxIdleConn:     10,
	MaxOpenConn:     30,
	ConnMaxLifeTime: time.Minute * 60,
}

func NewDBEngine(database *DatabaseSetting, runMode string) (*gorm.DB, error) {
	db, err := gorm.Open(database.DBType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
		database.UserName,
		database.Password,
		database.Host,
		database.Port,
		database.DBName,
		database.Charset,
		database.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if runMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(database.MaxIdleConn)
	db.DB().SetMaxOpenConns(database.MaxOpenConn)
	db.DB().SetConnMaxLifetime(database.ConnMaxLifeTime)
	otgorm.AddGormCallbacks(db)
	return db, nil
}
