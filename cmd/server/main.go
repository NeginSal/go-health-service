package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	log.Println("server started on :8080")
	log.Println("version is 1.0.0")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
