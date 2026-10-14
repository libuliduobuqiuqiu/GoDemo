package app

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func multipartRequest(t *testing.T, path, field, filename, content string) *http.Request {
	t.Helper()
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile(field, filename)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.WriteString(part, content); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"username", "password", "address"} {
		if err := writer.WriteField(name, "demo"); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, path, &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}

func TestLegacyBindingAndUploads(t *testing.T) {
	router := gin.New()
	jsonGroupInit(router)
	formGroupInit(router)
	request := func(req *http.Request, want int) string {
		t.Helper()
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		if recorder.Code != want {
			t.Fatalf("%s: got %d, want %d, body=%s", req.URL.Path, recorder.Code, want, recorder.Body.String())
		}
		return recorder.Body.String()
	}
	badJSON := httptest.NewRequest(http.MethodPost, "/json-group/json", strings.NewReader("{"))
	badJSON.Header.Set("Content-Type", "application/json")
	if body := request(badJSON, http.StatusBadRequest); strings.Contains(body, "posting ok") {
		t.Fatalf("invalid JSON also returned success: %s", body)
	}
	form := func(values url.Values) *http.Request {
		req := httptest.NewRequest(http.MethodPost, "/form-group/login", strings.NewReader(values.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		return req
	}
	request(form(url.Values{"user": {"user"}}), http.StatusBadRequest)
	request(form(url.Values{"user": {"user"}, "password": {"wrong"}}), http.StatusUnauthorized)
	request(form(url.Values{"user": {"user"}, "password": {"password"}}), http.StatusOK)
	request(httptest.NewRequest(http.MethodPost, "/form-group/upload", nil), http.StatusBadRequest)
	if body := request(multipartRequest(t, "/form-group/upload", "file", "sample.txt", "hello"), http.StatusOK); !strings.Contains(body, `"size":5`) {
		t.Fatalf("unexpected upload metadata: %s", body)
	}
	request(multipartRequest(t, "/form-group/uploads", "wrong", "sample.txt", "hello"), http.StatusBadRequest)
	if body := request(multipartRequest(t, "/form-group/uploads", "upload[]", "sample.txt", "hello"), http.StatusOK); !strings.Contains(body, `"count":1`) {
		t.Fatalf("unexpected upload count: %s", body)
	}
}

func TestStandaloneUploadIsBoundedAndDoesNotWriteFile(t *testing.T) {
	t.Chdir(t.TempDir())
	router := gin.New()
	router.POST("/upload", UploadFile)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, multipartRequest(t, "/upload", "file", "sample.txt", "hello"))
	if recorder.Code != http.StatusOK {
		t.Fatalf("upload status = %d: %s", recorder.Code, recorder.Body.String())
	}
	if _, err := os.Stat("tmpFile"); !os.IsNotExist(err) {
		t.Fatalf("unexpected output file: %v", err)
	}
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, multipartRequest(t, "/upload", "file", "large.bin", strings.Repeat("x", maxUploadBytes)))
	if recorder.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("large upload status = %d: %s", recorder.Code, recorder.Body.String())
	}
}

func TestLegacyRemoteReaderClosesResponseAndReportsFailure(t *testing.T) {
	closed := false
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: http.StatusBadGateway, Body: closeTracker{Reader: strings.NewReader("upstream error"), closed: &closed}, Header: make(http.Header)}, nil
	})}
	router := gin.New()
	othersGroupInit(router, "https://example.test/image.png", client)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/others/someDataFromReader", nil))
	if recorder.Code != http.StatusServiceUnavailable || !closed {
		t.Fatalf("remote failure: status=%d closed=%t", recorder.Code, closed)
	}
	closed = false
	client.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode:    http.StatusOK,
			Body:          closeTracker{Reader: strings.NewReader("image bytes"), closed: &closed},
			Header:        http.Header{"Content-Type": []string{"image/png"}},
			ContentLength: int64(len("image bytes")),
		}, nil
	})
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/others/someDataFromReader", nil))
	if recorder.Code != http.StatusOK || recorder.Body.String() != "image bytes" || !closed {
		t.Fatalf("remote success: status=%d body=%q closed=%t", recorder.Code, recorder.Body.String(), closed)
	}
}

type closeTracker struct {
	io.Reader
	closed *bool
}

func (r closeTracker) Close() error { *r.closed = true; return nil }
