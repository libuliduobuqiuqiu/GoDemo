package app

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const maxUploadBytes = 8 << 20

// StartFileUploadServer runs the isolated upload example on localhost.
func StartFileUploadServer() error {
	r := gin.Default()
	r.POST("/upload", UploadFile)
	return r.Run("127.0.0.1:8090")
}

// UploadFile validates a multipart upload without writing client data to disk.
func UploadFile(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadBytes)
	file, err := c.FormFile("file")
	if err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "upload exceeds 8 MiB"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		}
		return
	}
	if errmsg := CheckParam(c.PostForm("username"), c.PostForm("password"), c.PostForm("address")); errmsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errmsg})
		return
	}

	input, err := file.Open()
	if err != nil {
		HandleErrResponse(c, err)
		return
	}
	defer input.Close()
	bytesRead, err := io.Copy(io.Discard, input)
	if err != nil {
		HandleErrResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"filename": file.Filename, "size": bytesRead})
}

func CheckParam(username, password, address string) (errmsg string) {
	if username == "" {
		errmsg += "username 不能为空;"
	}
	if password == "" {
		errmsg += "password 不能为空;"
	}
	if address == "" {
		errmsg += "address 不能为空;"
	}
	return
}
