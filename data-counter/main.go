package main

import (
	"io"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	dataIngested = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "otel_prw_compare_data_ingested_bytes_total",
		Help: "The total bytes ingested by the app.",
	}, []string{"route"})
)

func countData(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Body.Close()

	dataIngested.WithLabelValues(req.RequestURI).Add(float64(len(body)))

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", countData)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8090", nil)
}
