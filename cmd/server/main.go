package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

var healthRequests = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "health_requests_total",
		Help: "Total number of requests to health endpoint",
	},
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	healthRequests.Inc()

	resp := HealthResponse{
		Status:  "ok",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	resp := VersionResponse{
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(resp)
}

func main() {

	prometheus.MustRegister(healthRequests)

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("server started on :8080")
	log.Println("version is 1.0.0")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
