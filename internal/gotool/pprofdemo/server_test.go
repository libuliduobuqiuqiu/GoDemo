package pprofdemo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestCPUHandler(t *testing.T) {
	handler := NewHandler()
	for _, duration := range []string{"bad", "0s", "-1s", "121s"} {
		r := httptest.NewRecorder()
		handler.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/work/cpu?duration="+duration, nil))
		if r.Code != http.StatusBadRequest {
			t.Fatalf("duration %q: status %d", duration, r.Code)
		}
	}
	r := httptest.NewRecorder()
	handler.ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/work/cpu?duration=20ms", nil))
	var result cpuResult
	if err := json.Unmarshal(r.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if r.Code != http.StatusOK || result.Iterations == 0 || result.Elapsed < 20*time.Millisecond || result.Checksum == [32]byte{} {
		t.Fatalf("unexpected response: status=%d result=%+v", r.Code, result)
	}
}

func TestCPUWorkCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if result := burnCPU(ctx, time.Minute); result.Iterations != 0 {
		t.Fatalf("canceled work performed %d iterations", result.Iterations)
	}
}

// TestHTTPProfileHasCPUActivity verifies samples, rather than just a nonempty gzip file.
func TestHTTPProfileHasCPUActivity(t *testing.T) {
	if testing.Short() {
		t.Skip("CPU profiling takes two seconds")
	}
	server := httptest.NewServer(NewHandler())
	defer server.Close()
	client := &http.Client{Timeout: 10 * time.Second}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	started := make(chan struct{})
	go func() {
		defer close(done)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/work/cpu?duration=30s", nil)
		if err != nil {
			t.Error(err)
			close(started)
			return
		}
		close(started)
		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			t.Error("CPU request returned before cancellation")
		} else if ctx.Err() == nil {
			t.Errorf("CPU request failed: %v", err)
		}
	}()
	defer func() { cancel(); <-done }()
	<-started
	profile, err := client.Get(server.URL + "/debug/pprof/profile?seconds=2")
	if err != nil {
		t.Fatal(err)
	}
	defer profile.Body.Close()
	data, err := io.ReadAll(profile.Body)
	if err != nil {
		t.Fatal(err)
	}
	if profile.StatusCode != http.StatusOK {
		t.Fatalf("profile status=%d: %s", profile.StatusCode, data)
	}
	path := filepath.Join(t.TempDir(), "cpu.pprof")
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	cancel()
	cmd := exec.Command("go", "tool", "pprof", "-top", "-nodecount=100", path)
	if binary := os.Getenv("PPROF_BINARY"); binary != "" {
		cmd = exec.Command(binary, "-top", "-nodecount=100", path)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("pprof: %v\n%s", err, output)
	}
	if !strings.Contains(string(output), "burnCPU") {
		t.Fatalf("CPU workload missing from profile:\n%s", output)
	}
	t.Logf("CPU profile contains workload samples:\n%s", output)
}
