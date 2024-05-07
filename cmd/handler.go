package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/th3g3ntl3m3n/emplyee_db/internal/db"
)

type Handler struct {
	db db.Database
}

func NewHandler() *Handler {
	return &Handler{
		db: db.NewDB(),
	}
}

func (h Handler) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	skipValue := r.URL.Query().Get("skip")
	limitValue := r.URL.Query().Get("limit")

	skip, _ := strconv.Atoi(skipValue)
	limit, _ := strconv.Atoi(limitValue)

	res, err := h.db.GetAllEmployee(skip, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var req db.Employee
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w)
		return
	}

	// simple validation to check if all fields are set or not
	// if no then we are returning error
	// I am assuming all fields must be filled
	if req.Name == "" || req.Position == "" || req.Salary == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w)
		return
	}

	emp, err := h.db.CreateEmployee(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
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
