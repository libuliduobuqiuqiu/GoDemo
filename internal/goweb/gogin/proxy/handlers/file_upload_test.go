package handlers

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestParseFormDataResponses(t *testing.T) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if err := writer.WriteField("username", "demo"); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	router := gin.New()
	router.POST("/upload", ParseFormData)
	for _, tc := range []struct {
		name, contentType, body string
		want                    int
	}{
		{"valid multipart", writer.FormDataContentType(), body.String(), http.StatusOK},
		{"missing boundary", "multipart/form-data", body.String(), http.StatusBadRequest},
		{"wrong media type", "text/plain", "hello", http.StatusBadRequest},
		{"too large", writer.FormDataContentType(), strings.Repeat("x", maxFormBytes+1), http.StatusRequestEntityTooLarge},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/upload", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", tc.contentType)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			if recorder.Code != tc.want {
				t.Fatalf("got %d, want %d: %s", recorder.Code, tc.want, recorder.Body.String())
			}
		})
	}
}
