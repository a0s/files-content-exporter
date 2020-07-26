package main

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type envsSettings struct {
	ConfigFilePath string `env:"FILES_CONTENT_EXPORTER_CONFIG_FILE_PATH" envDefault:"/config.yml"`
	Host           string `env:"FILES_CONTENT_EXPORTER_HOST" envDefault:"127.0.0.1"`
	Port           uint16 `env:"FILES_CONTENT_EXPORTER_PORT" envDefault:"9457"`
	LogLevel       string `env:"FILES_CONTENT_EXPORTER_LOG_LEVEL" envDefault:"INFO"`
}

var envs envsSettings

func init() {
	envs = envsSettings{}
	if err := env.Parse(&envs); err != nil {
		log.Fatalf("env parsing %v", err)
	}

	switch envs.LogLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	default:
		log.Fatalf("unknown log level `%v'", envs.LogLevel)
	}
}
