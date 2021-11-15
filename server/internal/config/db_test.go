package config

import (
	"testing"
	"time"
)

var testDataBase = DatabaseSetting{
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

func TestNewDBEngine(t *testing.T) {
	db, err := NewDBEngine(&testDataBase, "debug")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", db.DB().Stats())
}
