package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create a test gauge with name and help string
	testGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test_gauge",
		Help: "Test gauge. Its value is always 42.",
	})

	// Register the gauge in the global metrics registry
	prometheus.MustRegister(testGauge)

	// Set gauge to 42
	testGauge.Set(42)

	//Expose all metrics in the default registry on /metrics
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server listening")
	// Listen for HTTP requests on port 8090
	log.Fatal(http.ListenAndServe(":8090", nil))
}
