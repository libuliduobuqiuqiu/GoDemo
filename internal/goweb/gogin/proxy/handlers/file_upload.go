package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"godemo/internal/goweb/gogin/proxy/public"
	"io"
	"mime"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

const maxFormBytes = 8 << 20

// ParseFormData checks a bounded multipart body without storing or logging its contents.
func ParseFormData(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxFormBytes)
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "form exceeds 8 MiB"})
		} else {
			public.HandleErrJson(c, err)
		}
		return
	}
	if err := GetFormData(body, c.GetHeader("Content-Type")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	public.HandleSuccessJson(c, struct{}{})
}

func GetFormData(requestBody []byte, contentType string) error {
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil || mediaType != "multipart/form-data" || params["boundary"] == "" {
		return fmt.Errorf("multipart/form-data with boundary required")
	}
	reader := multipart.NewReader(bytes.NewReader(requestBody), params["boundary"])
	parts := 0
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if _, err := io.Copy(io.Discard, part); err != nil {
			part.Close()
			return err
		}
		if err := part.Close(); err != nil {
			return err
		}
		parts++
	}
	if parts == 0 {
		return fmt.Errorf("multipart form is empty")
	}
	return nil
}
