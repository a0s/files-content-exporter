package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	configFile := readYamlConfig(envs.ConfigFilePath)

	gauges := newGaugedEntities(configFile)

	http.Handle("/metrics", newWrapper(wrapperFunc(gauges), promhttp.Handler()))

	addr := fmt.Sprintf("%s:%d", envs.Host, envs.Port)
	log.Infof("starting prometheus exporter at %s\n", addr)

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

func updateMetric(entity *entity, gauge *prometheus.Gauge, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := os.Stat(entity.File)
	if err != nil {
		log.Warnf("file is not existing %v", entity.File)
		return
	}

	content, err := ioutil.ReadFile(entity.File)
	if err != nil {
		log.Warnf("can't read file %v", entity.File)
		return
	}

	prepared := strings.TrimSpace(string(content))

	value, err := strconv.ParseFloat(prepared, 64)
	if err != nil {
		log.Warnf("can't convert '%v' from %v to Float64", string(content), entity.File)
		return
	}

	(*gauge).Set(value)
}
