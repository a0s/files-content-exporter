package main

import (
	"github.com/prometheus/client_golang/prometheus"
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
