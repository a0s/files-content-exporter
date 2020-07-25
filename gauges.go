package main

import "github.com/prometheus/client_golang/prometheus"

type gaugedEntities map[*entity]*prometheus.Gauge

func newGaugedEntities(cf *yamlConfig) *gaugedEntities {
	gauges := make(gaugedEntities)

	for _, configEntity := range cf.Entities {
		tempLabels := make(map[string]string)
		for k, v := range configEntity.Labels {
			tempLabels[k] = v
		}
		if cf.PathAsLabelEnabled == true {
			tempLabels["path"] = configEntity.File
		}
		newGauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Name:        configEntity.Name,
			Help:        configEntity.Help,
			ConstLabels: tempLabels,
		})
		prometheus.MustRegister(newGauge)
		gauges[&configEntity] = &newGauge
	}

	return &gauges
}
