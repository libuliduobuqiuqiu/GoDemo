package app

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123423423"},
	"austin": gin.H{"email": "asutil@example.com", "phone": "6666"},
	"lena":   gin.H{"emial": "lena@example.coim", "phone": "5234423"},
}

func postJSON(c *gin.Context) {
	var user struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
		return
	}
	name := user.Name
	age := user.Age
	c.String(http.StatusOK, name+","+age+", posting ok.")
}

func jsonGroupInit(engine *gin.Engine) {
	address := map[string]interface{}{
		"name":    "golang",
		"country": "china",
		"city":    "guangzhou",
	}

	group := engine.Group("json-group")
	group.GET("/print-json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "2022-11-1",
		})
	})

	// unicode ascii json
	group.GET("/print-assciijson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "1.19",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	// javascript
	group.GET("/jsonp", func(c *gin.Context) {
		data2 := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data2)
	})

	// json
	group.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hell, world</b>",
		})
	})
	group.POST("json", postJSON)

	// pure json
	group.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world</b>",
		})
	})

	// message
	group.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name     string `json:"name"`
			Messages string `json:"message"`
			Number   int    `json:"number"`
		}
		msg.Name = "Lena"
		msg.Messages = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	group.GET("/someJSON", func(c *gin.Context) {
		// userList := []string{"wangwu", "zhangsna", "foo"}
		c.SecureJSON(http.StatusOK, address)
	})

}

func formGroupInit(engine *gin.Engine) {
	group := engine.Group("form-group")
	group.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user and password are required"})
			return
		}
		if form.User == "user" && form.Password == "password" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	})

	group.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	group.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		c.JSON(200, gin.H{
			"code":    200,
			"info":    "success",
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	engine.MaxMultipartMemory = 8 << 20
	group.POST("/upload", func(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{"filename": file.Filename, "size": file.Size})
	})

	group.POST("/uploads", func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadBytes)
		form, err := c.MultipartForm()
		if err != nil {
			var tooLarge *http.MaxBytesError
			if errors.As(err, &tooLarge) {
				c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "upload exceeds 8 MiB"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "multipart form is required"})
			}
			return
		}
		files := form.File["upload[]"]
		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "upload[] is required"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": len(files)})
	})

}

func othersGroupInit(engine *gin.Engine, imageURL string, client *http.Client) {
	if client == nil {
		client = newLegacyHTTPClient()
	}
	group := engine.Group("others")
	group.GET("/someDataFromReader", func(c *gin.Context) {
		// 从指定地址获取图片
		request, err := http.NewRequestWithContext(c.Request.Context(), http.MethodGet, imageURL, nil)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		response, err := client.Do(request)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		// 获取响应头属性
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		// 将读取的图片，响应头属性
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	// xml
	group.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})

	})

	// yaml
	group.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// 使用BasicAuth组件
	// gin.Accounts是map[string]string的一种快捷方式
	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由BasicAuth中间设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "No SECRET: ("})
		}
	})

}

func newLegacyHTTPClient() *http.Client {
	return &http.Client{Timeout: 5 * time.Second}
}

func NewRouter(templatePath string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.LoadHTMLFiles(templatePath)
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main WebSite",
		})
	})
	jsonGroupInit(r)
	return r
}

func Start() error {
	return NewRouter("web/templates/index.html").Run("127.0.0.1:8999")
}
