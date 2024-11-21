package main

import (
	"log"
	"net/http"
	"swagger-test-generator/handlers"
	"swagger-test-generator/metrics"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Initialisiere Metriken
	metrics.InitializeMetrics()

	// Erstelle einen neuen Router
	r := mux.NewRouter()

	// API Endpunkt für die Testgenerierung
	r.HandleFunc("/generate-tests", handlers.HandleGenerateTests).Methods("POST")

	// Füge Swagger UI hinzu
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Füge den /metrics Endpunkt für Prometheus hinzu
	r.Handle("/metrics", promhttp.Handler())

	// Erstelle einen HTTP-Server mit Timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Starte den Server
	log.Println("Server läuft auf https://localhost:8080...")
	log.Fatal(server.ListenAndServeTLS("scripts/server.crt", "scripts/server.key"))
}
