package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCPUWorkRouteOverHTTP(t *testing.T) {
	router := NewRouter(filepath.Join("..", "..", "..", "..", "web", "templates", "index.html"))
	server := httptest.NewServer(router)
	defer server.Close()
	client := &http.Client{Timeout: 5 * time.Second}
	start := time.Now()
	response, err := client.Get(server.URL + "/work/cpu?duration=1s")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	var result struct {
		Iterations uint64 `json:"iterations"`
		Elapsed    int64  `json:"elapsed_ns"`
	}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	if response.StatusCode != http.StatusOK || time.Since(start) < time.Second || result.Elapsed < int64(time.Second) || result.Iterations == 0 {
		t.Fatalf("CPU request returned too early or failed: status=%d elapsed=%s result=%+v", response.StatusCode, time.Since(start), result)
	}
	t.Logf("HTTP CPU request took %s, iterations=%d", time.Since(start), result.Iterations)
	response, err = client.Get(server.URL + "/debug/pprof/")
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Fatalf("pprof status=%d", response.StatusCode)
	}
}

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
