package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metriken für erfolgreiche und fehlgeschlagene Anfragen
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

// Funktion zur Initialisierung und Registrierung der Metriken bei Prometheus
func InitializeMetrics() {
	prometheus.MustRegister(successfulRequests)
	prometheus.MustRegister(failedRequests)
}

// Funktion zum Inkrementieren des Zählers für erfolgreiche Anfragen
func IncrementSuccessfulRequests(status string) {
	successfulRequests.WithLabelValues(status).Inc()
}

// Funktion zum Inkrementieren des Zählers für fehlgeschlagene Anfragen
func IncrementFailedRequests(status string) {
	failedRequests.WithLabelValues(status).Inc()
}
