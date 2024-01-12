package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/db"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	db.SetupDB()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
