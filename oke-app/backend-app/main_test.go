package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/db"
	"github.com/stretchr/testify/assert"
)

func TestGetAllItems(t *testing.T) {
	db.SetupDB()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetId1(t *testing.T) {
	db.SetupDB()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestPostNewItem(t *testing.T) {
	router := setupRouter()
	body := strings.NewReader(`{"Name": "速習Golang", "Date": "20240207190000", "Topics": "Golang", "Presenters": "Takuya Niita"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteNewItem(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
