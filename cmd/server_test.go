package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/th3g3ntl3m3n/emplyee_db/internal/db"
)

func TestServerGetAllEmployee(t *testing.T) {
	req := httptest.NewRequest("GET", "/employees?skip=0&limit=5", nil)
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)

	resp := []db.Employee{}
	json.NewDecoder(rr.Body).Decode(&resp)

	assert.Equal(t, len(resp), 0)
}
func TestServerGetAllEmployeeBadRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/employees", nil)
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestServerGetEmployeeByID(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)

	getReq := httptest.NewRequest("GET", "/employee/"+empResp.ID, nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, getReq)

	assert.Equal(t, rr.Code, http.StatusOK)

	empRespGet := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empRespGet)

	assert.Equal(t, empRespGet, empResp)
}

func TestServerGetEmployeeByIDBadRequest(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)

	getReq := httptest.NewRequest("GET", "/employee/abcd", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, getReq)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

// Test for create employee
func TestServerCreateEmployee(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)
	assert.NotEmpty(t, empResp.ID)
	assert.Equal(t, empResp.Name, "Vikas")
	assert.Equal(t, empResp.Salary, int32(100))
	assert.Equal(t, empResp.Position, "Dev")

	getReq := httptest.NewRequest("GET", "/employee/"+empResp.ID, nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, getReq)
	assert.Equal(t, rr.Code, http.StatusOK)
}

func TestServerCreateEmployeeBadRequest(t *testing.T) {
	emp := `{"Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)

}

func TestServerCreateEmployeeBadRequest2(t *testing.T) {
	emp := `{"Position": 1029}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)

}

// Test for update employee
func TestServerUpdateEmployee(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)

	empResp.Name = "NewName"
	empResp.Position = "EM"
	empResp.Salary = 10e6

	jsonBytes, _ := json.Marshal(empResp)

	updateReq := httptest.NewRequest("PATCH", "/employee/"+empResp.ID, bytes.NewBuffer(jsonBytes))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, updateReq)

	assert.Equal(t, rr.Code, http.StatusOK)

	empRespUpdate := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empRespUpdate)

	assert.Equal(t, empRespUpdate, empResp)
}

func TestServerUpdateEmployeeBadRequest(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)

	empResp.Name = "NewName"
	empResp.Position = ""
	empResp.Salary = 0

	jsonBytes, _ := json.Marshal(empResp)

	updateReq := httptest.NewRequest("PATCH", "/employee/"+empResp.ID, bytes.NewBuffer(jsonBytes))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, updateReq)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestServerUpdateEmployeeBadRequest2(t *testing.T) {
	emp := `{"Name": "12345", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)

	emp = `{"Name": 12345, "Salary": 100, "Position": "Dev"}`

	updateReq := httptest.NewRequest("PATCH", "/employee/"+empResp.ID, strings.NewReader(emp))
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, updateReq)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestServerUpdateEmployeeBadRequest3(t *testing.T) {
	emp := `{"Name": "12345", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("PATCH", "/employee/abcde", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest)
}

func TestServerUpdateEmployeeNotFound(t *testing.T) {
	emp := `{"Name": "12345", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("PATCH", "/employees/01HX9YP3J2PB9RMJXVWKPCWTN0", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusNotFound)
}

// Test for Delete Employee
func TestServerDeleteEmployee(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)

	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)

	deleteReq := httptest.NewRequest("DELETE", "/employee/"+empResp.ID, nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, deleteReq)

	assert.Equal(t, rr.Code, http.StatusOK)

	empRespDelete := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empRespDelete)

	assert.Equal(t, empRespDelete.ID, empResp.ID)

	getReq := httptest.NewRequest("GET", "/employee/"+empResp.ID, nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, getReq)

	assert.Equal(t, rr.Code, http.StatusNotFound)
}

func TestServerDeleteEmployeeBadRequest(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)
	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)
	log.Println(empResp.ID)

	deleteReq := httptest.NewRequest("DELETE", "/employee/abcdef", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, deleteReq)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestServerDeleteEmployeeNotFound(t *testing.T) {
	emp := `{"Name": "Vikas", "Salary": 100, "Position": "Dev"}`

	req := httptest.NewRequest("POST", "/employees", strings.NewReader(emp))
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusCreated)
	empResp := db.Employee{}
	json.NewDecoder(rr.Body).Decode(&empResp)
	log.Println(empResp.ID)

	deleteReq := httptest.NewRequest("DELETE", "/employee/01HX9YP3J2PB9RMJXVWKPCWTN0", nil)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, deleteReq)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
