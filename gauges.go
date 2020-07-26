package main

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

type gaugedEntities map[*entity]*prometheus.Gauge

func newGaugedEntities(cf *yamlConfig) *gaugedEntities {
	gauges := make(gaugedEntities, len(cf.Entities))

	for _, e := range cf.Entities {
		entity := e

		tempLabels := make(map[string]string)
		for k, v := range entity.Labels {
			tempLabels[k] = v
		}
		if cf.PathAsLabelEnabled == true {
			tempLabels["path"] = entity.File
		}
		gauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Name:        entity.Name,
			Help:        entity.Help,
			ConstLabels: tempLabels,
		})

		prometheus.MustRegister(gauge)
		gauges[&entity] = &gauge
	}

	return &gauges
}

func updateGaugeWithEntity(gauge *prometheus.Gauge, entity *entity, wg *sync.WaitGroup) {
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
