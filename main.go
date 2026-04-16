package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bcgov/go-sample-api/config"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, OpenShift!!!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/openapi.yaml")
	})
	// serve the swagger-ui index and assets (we use a small index.html + CDN for assets)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs/swagger-ui"))))

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
