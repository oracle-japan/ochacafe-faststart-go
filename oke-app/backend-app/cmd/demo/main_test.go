package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/internal/pkg/db"
	"github.com/oracle-japan/ochacafe-faststart-go/oke-app/backend-app/internal/pkg/repo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	db.SetupDB()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetAllItems(t *testing.T) {
	if shutdown := retryInitTracer(); shutdown != nil {
		defer shutdown()
	}

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w, req)

	var itemsList []repo.Items
	err := json.Unmarshal(w.Body.Bytes(), &itemsList)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	// サボって件数だけ見る
	assert.Equal(t, 6, len(itemsList))
}

func TestGetById(t *testing.T) {
	if shutdown := retryInitTracer(); shutdown != nil {
		defer shutdown()
	}
	router := setupRouter()

	// 1件だけ適当に取る
	w_test := httptest.NewRecorder()
	req_test, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w_test, req_test)
	var itemsList []repo.Items
	err_test := json.Unmarshal(w_test.Body.Bytes(), &itemsList)
	if err_test != nil {
		log.Fatal(err_test)
	}

	// 適当に取得したデータでテスト
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items/"+strconv.FormatUint(uint64(itemsList[0].ID), 10), nil)
	router.ServeHTTP(w, req)

	var items repo.Items
	err := json.Unmarshal(w.Body.Bytes(), &items)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, itemsList[0].ID, items.ID)
}

func TestPostNewItem(t *testing.T) {
	if shutdown := retryInitTracer(); shutdown != nil {
		defer shutdown()
	}
	router := setupRouter()
	body := strings.NewReader(`{"Name": "速習Golang", "Date": "20240207190000", "Topics": "Golang", "Presenters": "Takuya Niita"}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/items", body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// 更新したレコード数
	assert.Equal(t, "1", w.Body.String())
}

func TestDeleteNewItem(t *testing.T) {
	if shutdown := retryInitTracer(); shutdown != nil {
		defer shutdown()
	}

	router := setupRouter()

	// 1件だけ適当に取る
	w_test := httptest.NewRecorder()
	req_test, _ := http.NewRequest("GET", "/items", nil)
	router.ServeHTTP(w_test, req_test)
	var itemsList []repo.Items
	err_test := json.Unmarshal(w_test.Body.Bytes(), &itemsList)
	if err_test != nil {
		log.Fatal(err_test)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/items/"+strconv.FormatUint(uint64(itemsList[0].ID), 10), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// 削除したレコード数
	assert.Equal(t, "1", w.Body.String())
}
