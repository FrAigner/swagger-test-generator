package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	successfulRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "swagger_test_generator_successful_requests_total",
			Help: "Anzahl der erfolgreichen Anfragen zur Testgenerierung",
		},
		[]string{"status"},
	)
	failedRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "swagger_test_generator_failed_requests_total",
			Help: "Anzahl der fehlgeschlagenen Anfragen",
		},
		[]string{"status"},
	)
)

func InitializeMetrics() {
	prometheus.MustRegister(successfulRequests)
	prometheus.MustRegister(failedRequests)
}

func IncrementSuccessfulRequests(status string) {
	successfulRequests.WithLabelValues(status).Inc()
}

func IncrementFailedRequests(status string) {
	failedRequests.WithLabelValues(status).Inc()
}
