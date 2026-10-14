package app

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func TestIndexRoute(t *testing.T) {
	router := NewRouter(filepath.Join("..", "..", "..", "..", "web", "templates", "index.html"))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/index", nil))
	if recorder.Code != http.StatusOK || !strings.Contains(recorder.Body.String(), "Main WebSite") {
		t.Fatalf("index route: %d %s", recorder.Code, recorder.Body.String())
	}
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/json-group/print-json", nil))
	if recorder.Code != http.StatusOK || !strings.Contains(recorder.Body.String(), `"test":"2022-11-1"`) {
		t.Fatalf("JSON route: %d %s", recorder.Code, recorder.Body.String())
	}
}
