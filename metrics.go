// SPDX-License-Identifier: BSD-3-Clause

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	namespace = ""
	registry  = prometheus.NewRegistry()

	DefaultRegisterer prometheus.Registerer = registry
	DefaultGatherer   prometheus.Gatherer   = registry
)

func Register(c prometheus.Collector) error {
	return DefaultRegisterer.Register(c)
}

func MustRegister(cs ...prometheus.Collector) {
	DefaultRegisterer.MustRegister(cs...)
}

func Unregister(c prometheus.Collector) bool {
	return DefaultRegisterer.Unregister(c)
}

func NewCounter(name, help string) prometheus.Counter {
	return NewCounterFor(DefaultRegisterer, name, help)
}

func NewCounterFor(registerer prometheus.Registerer, name, help string) prometheus.Counter {
	return promauto.With(registerer).NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func NewGauge(name, help string) prometheus.Gauge {
	return NewGaugeFor(DefaultRegisterer, name, help)
}

func NewGaugeFor(registerer prometheus.Registerer, name, help string) prometheus.Gauge {
	return promauto.With(registerer).NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func NewHistogram(name, help string, buckets []float64) prometheus.Histogram {
	return NewHistogramFor(DefaultRegisterer, name, help, buckets)
}

func NewHistogramFor(registerer prometheus.Registerer, name, help string, buckets []float64) prometheus.Histogram {
	return promauto.With(registerer).NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	})
}

func NewSummary(name, help string, objectives map[float64]float64) prometheus.Summary {
	return NewSummaryFor(DefaultRegisterer, name, help, objectives)
}

func NewSummaryFor(registerer prometheus.Registerer, name, help string, objectives map[float64]float64) prometheus.Summary {
	return promauto.With(registerer).NewSummary(prometheus.SummaryOpts{
		Namespace:  namespace,
		Name:       name,
		Help:       help,
		Objectives: objectives,
	})
}

func NewCounterVec(name, help string, labelValues []string) *prometheus.CounterVec {
	return NewCounterVecFor(DefaultRegisterer, name, help, labelValues)
}

func NewCounterVecFor(registerer prometheus.Registerer, name, help string, labelValues []string) *prometheus.CounterVec {
	return promauto.With(registerer).NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labelValues)
}

func NewGaugeVec(name, help string, labelValues []string) *prometheus.GaugeVec {
	return NewGaugeVecFor(DefaultRegisterer, name, help, labelValues)
}

func NewGaugeVecFor(registerer prometheus.Registerer, name, help string, labelValues []string) *prometheus.GaugeVec {
	return promauto.With(registerer).NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labelValues)
}

func NewHistogramVec(name, help string, buckets []float64, labelValues []string) *prometheus.HistogramVec {
	return NewHistogramVecFor(DefaultRegisterer, name, help, buckets, labelValues)
}

func NewHistogramVecFor(registerer prometheus.Registerer, name, help string, buckets []float64, labelValues []string) *prometheus.HistogramVec {
	return promauto.With(registerer).NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labelValues)
}

func NewSummaryVec(name, help string, objectives map[float64]float64, labelValues []string) *prometheus.SummaryVec {
	return NewSummaryVecFor(DefaultRegisterer, name, help, objectives, labelValues)
}

func NewSummaryVecFor(registerer prometheus.Registerer, name, help string, objectives map[float64]float64, labelValues []string) *prometheus.SummaryVec {
	return promauto.With(registerer).NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  namespace,
		Name:       name,
		Help:       help,
		Objectives: objectives,
	}, labelValues)
}
