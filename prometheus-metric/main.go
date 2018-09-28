package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	log.Println("Start ..")

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/metrics", prometheus.Handler())
	http.ListenAndServe(":8080", nil)
}
