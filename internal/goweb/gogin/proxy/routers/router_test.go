package routers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRouterAndProxy(t *testing.T) {
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" || r.URL.Query().Get("name") != "go" {
			t.Errorf("unexpected upstream URL: %s", r.URL)
		}
		fmt.Fprint(w, "upstream ok")
	}))
	defer upstream.Close()

	router, err := NewRouter([]string{upstream.URL}, false)
	if err != nil {
		t.Fatal(err)
	}
	request := func(path string) *httptest.ResponseRecorder {
		t.Helper()
		recorder := httptest.NewRecorder()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, path, nil).WithContext(ctx))
		return recorder
	}
	if got := request("/healthz"); got.Code != http.StatusOK || !strings.Contains(got.Body.String(), `"ok"`) {
		t.Fatalf("health route: %d %s", got.Code, got.Body.String())
	}
	if got := request("/netac/base/ssh"); got.Code != http.StatusNotFound {
		t.Fatalf("SSH must be disabled by default: %d", got.Code)
	}
	proxyPath := "/netac/base/proxy?proxy_pass="
	if got := request(proxyPath + url.QueryEscape(upstream.URL+"/hello?name=go")); got.Code != http.StatusOK || got.Body.String() != "upstream ok" {
		t.Fatalf("proxy route: %d %s", got.Code, got.Body.String())
	}
	for _, tc := range []struct {
		target string
		status int
	}{
		{"", http.StatusBadRequest},
		{"file:///etc/passwd", http.StatusBadRequest},
		{"http://127.0.0.1:1/hello", http.StatusForbidden},
	} {
		got := request(proxyPath + url.QueryEscape(tc.target))
		if got.Code != tc.status {
			t.Errorf("target %q: got %d, want %d", tc.target, got.Code, tc.status)
		}
	}
}

func TestProxyVerifiesTLS(t *testing.T) {
	upstream := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "unexpected success")
	}))
	defer upstream.Close()
	router, err := NewRouter([]string{upstream.URL}, false)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/netac/base/proxy?proxy_pass="+url.QueryEscape(upstream.URL), nil).WithContext(ctx))
	if recorder.Code != http.StatusBadGateway {
		t.Fatalf("untrusted TLS certificate: got %d", recorder.Code)
	}
}

func TestSSHRequiresKnownHosts(t *testing.T) {
	t.Setenv("GO_DEMO_SSH_KNOWN_HOSTS", "")
	if _, err := NewRouter(nil, true); err == nil {
		t.Fatal("SSH routes started without known_hosts")
	}
}

func TestTerminalRejectsCrossOrigin(t *testing.T) {
	path := filepath.Join(t.TempDir(), "known_hosts")
	if err := os.WriteFile(path, nil, 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GO_DEMO_SSH_KNOWN_HOSTS", path)
	router, err := NewRouter(nil, true)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/netac/base/terminal", nil)
	req.Header.Set("Origin", "http://other.example")
	req.Header.Set("Connection", "Upgrade")
	req.Header.Set("Upgrade", "websocket")
	req.Header.Set("Sec-WebSocket-Version", "13")
	req.Header.Set("Sec-WebSocket-Key", "dGhlIHNhbXBsZSBub25jZQ==")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusForbidden {
		t.Fatalf("cross-origin WebSocket: got %d, want 403", recorder.Code)
	}
}
