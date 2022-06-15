package main

import "github.com/prometheus/client_golang/prometheus"

type IGrillExporter struct {
	up      prometheus.Gauge
	metrics struct {
		tempGuage      *prometheus.GaugeVec
		thresholdGuage *prometheus.GaugeVec
		systemInfo     *prometheus.GaugeVec
		up             prometheus.Gauge
	}
}

var metricNamespace = "igrill"

func NewIGrillExporter() *IGrillExporter {
	ig := &IGrillExporter{}
	ig.metrics.tempGuage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "temperature",
		Help:      "",
	}, []string{"probe"})

	ig.metrics.thresholdGuage = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "threshold",
		Help:      "",
	}, []string{"probe"})

	ig.metrics.systemInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "system_info",
		Help:      "",
	}, []string{"probe"})

	ig.metrics.up = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricNamespace,
		Name:      "up",
		Help:      "",
	})

	return ig
}

func (e *IGrillExporter) Describe(ch chan<- *prometheus.Desc) {
	e.metrics.tempGuage.Describe(ch)
	e.metrics.thresholdGuage.Describe(ch)
	e.metrics.systemInfo.Describe(ch)
	e.metrics.up.Describe(ch)

}
func (e *IGrillExporter) Collect(ch chan<- prometheus.Metric) {
	e.metrics.tempGuage.Collect(ch)
	e.metrics.thresholdGuage.Collect(ch)
	e.metrics.systemInfo.Collect(ch)
	e.metrics.tempGuage.Collect(ch)
}
