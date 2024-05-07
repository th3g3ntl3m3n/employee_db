package cmd

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()
	log := zap.NewExample().Sugar()
	defer log.Sync()

	mux.HandleFunc("GET /employees", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintln(w, `{"employees: []", "type": "GETALL"}`)
	})

	return &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
}
