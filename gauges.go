package main

import "github.com/prometheus/client_golang/prometheus"

type EntityWithGauge map[*ConfigFileEntity]*prometheus.Gauge

func initializeGauges(cf *ConfigFile) *EntityWithGauge {
	gauges := make(EntityWithGauge)
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
			ConstLabels: tempLabels,
		})
		prometheus.MustRegister(newGauge)
		gauges[&configEntity] = &newGauge
	}
	return &gauges
}
