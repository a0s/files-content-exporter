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

type envs struct {
	ConfigFilePath string `env:"FILES_CONTENT_EXPORTER_CONFIG_FILE_PATH,required"`
	Host           string `env:"FILES_CONTENT_EXPORTER_HOST" envDefault:"127.0.0.1"`
	Port           uint16 `env:"FILES_CONTENT_EXPORTER_PORT" envDefault:"9457"`
}

func main() {
	envs := &envs{}
	if err := env.Parse(envs); err != nil {
		log.Fatal(err)
	}

	configFile := readYamlConfig(envs.ConfigFilePath)

	gauges := newGaugedEntities(configFile)

	http.Handle("/metrics", newWrapper(wrapperFunc(gauges), promhttp.Handler()))

	addr := fmt.Sprintf("%s:%d", envs.Host, envs.Port)
	log.Printf("starting prometheus exporter at %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func wrapperFunc(ceg *gaugedEntities) func() {
	return func() {
		var wg sync.WaitGroup
		wg.Add(len(*ceg))

		for configEntity, gauge := range *ceg {
			go updateMetric(configEntity, gauge, &wg)
		}

		wg.Wait()
	}
}

func updateMetric(ce *entity, ga *prometheus.Gauge, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := os.Stat(ce.File)
	if err != nil {
		log.Printf("file is not existing %v", ce.File)
		return
	}

	content, err := ioutil.ReadFile(ce.File)
	if err != nil {
		log.Printf("can't read file %v", ce.File)
		return
	}

	value, err := strconv.ParseFloat(string(content), 64)
	if err != nil {
		log.Printf("can't convert '%v' from %v to Float64", string(content), ce.File)
		return
	}

	(*ga).Set(value)
}
