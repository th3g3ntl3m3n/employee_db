package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
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

func TestServerGetEmployeeByID(t *testing.T) {
	req := httptest.NewRequest("GET", "/employee/12HDHD34", nil)
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, strings.TrimSpace(rr.Body.String()), `{"employee: {"id": "12HDHD34"}", "type": "GETBYID"}`)
}
func TestServerGetCreateEmployee(t *testing.T) {
	req := httptest.NewRequest("POST", "/employees", nil)
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, strings.TrimSpace(rr.Body.String()), `{"employee: {"id": "newDataId"}", "type": "ADDNEW"}`)
}
func TestServerUpdateEmployee(t *testing.T) {
	req := httptest.NewRequest("PATCH", "/employee/123HHD", nil)
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, strings.TrimSpace(rr.Body.String()), `{"employee: {"id": "123HHD"}", "type": "PATCH"}`)
}
func TestServerDeleteEmployee(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/employee/123HHD", nil)
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, strings.TrimSpace(rr.Body.String()), `{"employee: {"id": "123HHD"}", "type": "DELETE"}`)
}
