package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/oklog/ulid/v2"
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

	if skipValue == "" || limitValue == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "bad request, skip and limit can't be empty")
		return
	}

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
		fmt.Fprintln(w, "bad request")
		return
	}

	emp, _ := h.db.CreateEmployee(req)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (h Handler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	employeeID := r.PathValue("id")
	_, err := ulid.Parse(employeeID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "bad request path value id is missing")

		return
	}

	emp, err := h.db.GetEmployeeByID(employeeID)
	if errors.Is(err, db.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "employee not found")

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (h Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var req db.Employee
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w)

		return
	}

	employeeID := r.PathValue("id")
	_, err := ulid.Parse(employeeID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "bad request path value id is missing")

		return
	}

	// simple validation to check if all fields are set or not
	// if no then we are returning error
	// I am assuming all fields must be filled
	if req.Name == "" || req.Position == "" || req.Salary == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "bad request, you can't set empty values")

		return
	}

	req.ID = employeeID
	emp, err := h.db.UpdateEmployee(req)
	if errors.Is(err, db.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "employee not found")

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func (h Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := r.PathValue("id")
	_, err := ulid.Parse(employeeID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "bad request path value id is missing")

		return
	}

	err = h.db.DeleteEmployee(employeeID)
	if errors.Is(err, db.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "employee not found")

		return
	}

	fmt.Fprintf(w, `{"ID": "%s"}`, employeeID)
}
