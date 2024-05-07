package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestServer(t *testing.T) {
	req := httptest.NewRequest("GET", "/employees", nil)
	rr := httptest.NewRecorder()

	handler := http.Handler(NewServer().Handler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, strings.TrimSpace(rr.Body.String()), `{"employees: []", "type": "GETALL"}`)
}
