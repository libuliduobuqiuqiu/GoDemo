package pprofdemo

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestPprofDefaultMuxRegistration(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/debug/pprof/", nil)
	defaultResponse := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(defaultResponse, req)
	if defaultResponse.Code != http.StatusOK {
		t.Fatalf("default mux status = %d", defaultResponse.Code)
	}
	customResponse := httptest.NewRecorder()
	http.NewServeMux().ServeHTTP(customResponse, req)
	if customResponse.Code != http.StatusNotFound {
		t.Fatalf("custom mux status = %d", customResponse.Code)
	}
}

func TestTraceToExplicitPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "fib.trace")
	if err := AnalysisFibByTrace(path); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() == 0 {
		t.Fatal("trace output is empty")
	}
}
