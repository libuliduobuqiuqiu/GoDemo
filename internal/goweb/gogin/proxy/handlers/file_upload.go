package handlers

import (
	"bytes"
	"fmt"
	"godemo/internal/goweb/gogin/proxy/public"
	"io"
	"mime"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logrus.WithError(err).Error()
		public.HandleErrJson(c, err)
		return
	}

	uploadFile, err := file.Open()
	if err != nil {
		logrus.WithError(err).Error()
		public.HandleErrJson(c, err)
		return
	}

	defer uploadFile.Close()

	fileContent, err := io.ReadAll(uploadFile)
	if err != nil {
		logrus.WithError(err).Error()
		public.HandleErrJson(c, err)
		return
	}

	if err := os.WriteFile("tmpFile", fileContent, 0644); err != nil {
		logrus.WithError(err).Error()
		public.HandleErrJson(c, err)
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	address := c.PostForm("address")

	fmt.Println(username, password, address)
	errmsg := CheckParam(username, password, address)
	if errmsg != "" {
		public.HandleErrJson(c, fmt.Errorf("%s", errmsg))
		return
	}

	public.HandleSuccessJson(c, struct{}{})
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

func ParseFormData(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.WithError(err).Error()
		public.HandleErrJson(c, err)
		return
	}

	contentType := c.GetHeader("Content-Type")
	fmt.Println(contentType)
	if err = GetFormData(body, contentType); err != nil {
		public.HandleErrJson(c, err)
		return
	}

	public.HandleSuccessJson(c, struct{}{})
}

func GetFormData(requestBody []byte, contentType string) (err error) {
	fmt.Println(contentType)
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		logrus.WithError(err).Error()
		return
	}

	boundary := params["boundary"]
	reader := multipart.NewReader(bytes.NewReader(requestBody), boundary)
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			logrus.WithError(err).Error()
			return err
		}

		name := part.FormName()
		fileName := part.FileName()
		data, err := io.ReadAll(part)
		if err != nil {
			logrus.WithError(err).Error()
			return err
		}

		if fileName == "" {
			value := string(data)
			fmt.Printf("FormName: %s, FormValue: %s \n", name, value)
		} else {
			fmt.Printf("file field: %s filename=%s size=%d\n", name, fileName, len(data))
			fmt.Println(string(data))
		}
	}

	return
}
