package config

import log "github.com/sirupsen/logrus"

func InitLogger(logLevel string) error {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
