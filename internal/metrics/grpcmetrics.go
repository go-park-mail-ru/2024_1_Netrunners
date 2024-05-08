package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type GrpcMetrics struct {
	service          string
	methodCallsTotal *prometheus.CounterVec
}

func InitGrpcMetrics(service string) *GrpcMetrics {
	return &GrpcMetrics{
		service: service,
		methodCallsTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "method_Calls_total",
				Help: "Total amount of grpc methods calls",
			},
			[]string{"service", "method"},
		),
	}
}

func (grpcMetrics *GrpcMetrics) Register() {
	prometheus.MustRegister(grpcMetrics.methodCallsTotal)
	prometheus.MustRegister(grpcMetrics.methodCallsTotal)
}

func (grpcMetrics *GrpcMetrics) IncRequestsTotal(method string) {
	grpcMetrics.methodCallsTotal.WithLabelValues(grpcMetrics.service, method).Inc()
}
