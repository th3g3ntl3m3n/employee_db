package cmd

import (
	"net/http"

	"go.uber.org/zap"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()
	log := zap.NewExample().Sugar()
	defer log.Sync()
	handler := NewHandler()

	mux.HandleFunc("GET /employees", handler.GetAllEmployee)
	mux.HandleFunc("POST /employees", handler.CreateEmployee)
	mux.HandleFunc("GET /employee/{id}", handler.GetEmployeeByID)
	mux.HandleFunc("PATCH /employee/{id}", handler.UpdateEmployee)
	mux.HandleFunc("DELETE /employee/{id}", handler.DeleteEmployee)

	return &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
}
