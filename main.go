package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
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
			go updateGaugeWithEntity(gauge, configEntity, &wg)
		}

		wg.Wait()
	}
}
