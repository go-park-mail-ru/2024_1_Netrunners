package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

type HttpMetrics struct {
	requestsTotal   *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

func InitHttpMetrics() *HttpMetrics {
	return &HttpMetrics{
		requestsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total amount of http requests with status codes",
			},
			[]string{"endpoint", "method", "status"},
		),
		requestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_requests_duration",
				Help:    "Duration of http requests with status codes",
				Buckets: []float64{0.001, 0.01, 0.1, 1, 10},
			},
			[]string{"method", "status"},
		),
	}
}

func (httpMetrics *HttpMetrics) Register() {
	prometheus.MustRegister(httpMetrics.requestsTotal)
	prometheus.MustRegister(httpMetrics.requestDuration)
}

func (httpMetrics *HttpMetrics) IncRequestsTotal(endpoint, method string, status int) {
	httpMetrics.requestsTotal.WithLabelValues(endpoint, method, fmt.Sprintf("%s", status)).Inc()
}

func (httpMetrics *HttpMetrics) IncRequestDuration(method string, status int, duration float64) {
	httpMetrics.requestDuration.WithLabelValues(method, fmt.Sprintf("%s", status)).Observe(duration)
}
