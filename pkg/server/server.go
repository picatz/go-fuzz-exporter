package server

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/picatz/go-fuzz-exporter/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// TODO(kent): make server more configurable
func Start(caseName string, src io.Reader) {
	go func() {
		addr := "0.0.0.0:4022"
		mux := http.NewServeMux()
		mux.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

		ln, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}

		srv := &http.Server{
			Handler:      mux,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  30 * time.Second,
		}

		log.Printf("starting go-fuzz-exporter server for %q available at %s", caseName, addr)

		srv.Serve(ln)
	}()

	metrics.Collect(caseName, src)
}
