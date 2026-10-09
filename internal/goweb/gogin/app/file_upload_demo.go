package app

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func StartFileUploadServer() {
	r := gin.Default()
	r.POST("/upload", UploadFile)

	r.Run(":8090")
}

func UploadFile(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		HandleErrResponse(c, err)
		return
	}

	fmt.Println(string(body))

	file, err := c.FormFile("file")
	if err != nil {
		HandleErrResponse(c, err)
		return
	}

	uploadFile, err := file.Open()
	if err != nil {
		HandleErrResponse(c, err)
		return
	}

	defer uploadFile.Close()

	fileContent, err := io.ReadAll(uploadFile)
	if err != nil {
		HandleErrResponse(c, err)
		return
	}

	if err := os.WriteFile("tmpFile", fileContent, 0644); err != nil {
		HandleErrResponse(c, err)
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	address := c.PostForm("address")

	fmt.Println(username, password, address)
	errmsg := CheckParam(username, password, address)
	if errmsg != "" {
		HandleErrResponse(c, fmt.Errorf("%s", errmsg))
		return
	}

	HandleSuccessResponse(c, struct{}{})
	return
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
