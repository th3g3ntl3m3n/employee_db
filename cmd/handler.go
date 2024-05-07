package cmd

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h Handler) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, `{"employees: []", "type": "GETALL"}`)
}

func (h Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, `{"employee: {"id": "newDataId"}", "type": "ADDNEW"}`)
}

func (h Handler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	pathValue := r.PathValue("id")
	fmt.Fprintf(w, `{"employee: {"id": "%s"}", "type": "GETBYID"}`, pathValue)
}

func (h Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	pathValue := r.PathValue("id")
	fmt.Fprintf(w, `{"employee: {"id": "%s"}", "type": "PATCH"}`, pathValue)
}

func (h Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	pathValue := r.PathValue("id")
	fmt.Fprintf(w, `{"employee: {"id": "%s"}", "type": "DELETE"}`, pathValue)
}
