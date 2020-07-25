package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type EnvConfig struct {
	ConfigFilePath string `env:"FILES_CONTENT_EXPORTER_CONFIG_FILE_PATH,required"`
	Host           string `env:"FILES_CONTENT_EXPORTER_HOST" envDefault:"127.0.0.1"`
	Port           uint16 `env:"FILES_CONTENT_EXPORTER_PORT" envDefault:"9457"`
}

func main() {
	envConfig := EnvConfig{}
	if err := env.Parse(&envConfig); err != nil {
		fmt.Printf("%+v\n", err)
	}

	configFile := *ReadConfigFile(envConfig.ConfigFilePath)

	gauges := *initializeGauges(&configFile)

	http.Handle("/metrics", NewHandlerWrapper(handlerWrapper(&gauges), promhttp.Handler()))

	addr := fmt.Sprintf("%s:%d", envConfig.Host, envConfig.Port)
	log.Printf("Starting prometheus exporter at %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Printf("http.ListenAndServer: %v\n", err)
	}
}

func handlerWrapper(ceg *EntityWithGauge) func() {
	return func() {
		var wg sync.WaitGroup
		wg.Add(len(*ceg))

		for configEntity, gauge := range *ceg {
			go updateMetric(configEntity, gauge, &wg)
		}

		wg.Wait()
	}
}

func updateMetric(ce *ConfigFileEntity, ga *prometheus.Gauge, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := os.Stat(ce.File)
	if err != nil {
		log.Printf("File is not existing %v", ce.File)
		return
	}

	content, err := ioutil.ReadFile(ce.File)
	if err != nil {
		log.Printf("Cannot read file %v", ce.File)
		return
	}

	value, err := strconv.ParseFloat(string(content), 64)
	if err != nil {
		log.Printf("Cannot convert '%v' from %v to Float64", string(content), ce.File)
		return
	}

	(*ga).Set(value)
}
