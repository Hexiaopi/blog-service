package config

import (
	log "github.com/sirupsen/logrus"
)

func Init() {
	var err error
	AppEngine, err = NewAppConfig("configs/app.yaml")
	if err != nil {
		log.Fatalf("init config fail %v", err)
	}
	log.Infof("config:%v", AppEngine)
	if err = InitLogger(AppEngine.LogLevel); err != nil {
		log.Fatalf("init log fail err:%v", err)
	}
	Tracer, _, err = NewJaegerTracer(AppEngine.ServiceName, AppEngine.TraceAgent)
	if err != nil {
		log.Fatal("init jaeger client fail.")
	}
	DBEngine, err = NewDBEngine(&AppEngine.DataBase, AppEngine.LogLevel)
	if err != nil {
		log.Fatalf("init db engine fail err:%v", err)
	}
	RedisEngine, err = NewCacheEngine(&AppEngine.Redis)
	if err != nil {
		log.Fatalf("init redis engine fail err:%v", err)
	}
}
