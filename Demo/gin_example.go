package Demo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form: "User" binding: "required"`
	Password string `form: "Password" binding: "required"`
}

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123423423"},
	"austin": gin.H{"email": "asutil@example.com", "phone": "6666"},
	"lena":   gin.H{"emial": "lena@example.coim", "phone": "5234423"},
}

func Gin_new() *gin.Engine {
	user := make(map[string]string)
	user["name"] = "linshukai"
	user["age"] = "22"
	user["sex"] = "main"

	address := make(map[string]interface{})
	address["name"] = "boning"
	address["country"] = "china"
	address["city"] = "guanghzou"

	r := gin.New()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("../templates/*")

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "2022-11-1",
		})
	})

	r.GET("/test2", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "1.19",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main WebSite",
		})
	})

	r.GET("/JSONP", func(c *gin.Context) {
		data2 := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data2)
	})

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind((&form)) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in "})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hell, world</b>",
		})
	})

	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world</b>",
		})
	})

	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s, page: %s, name: %s, message: %s", id, page, name, message)
		c.JSON(200, gin.H{
			"code": 200,
			"info": "success",
		})
	})

	r.GET("/someJSON", func(c *gin.Context) {
		// userList := []string{"wangwu", "zhangsna", "foo"}
		c.SecureJSON(http.StatusOK, address)
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name     string `json:user`
			Messages string
			Number   int
		}
		msg.Name = "Lena"
		msg.Messages = "hey"
		msg.Number = 123
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {

		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})

	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
	})

	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			dst := "./" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d file uploaded", len(files)))
	})

	r.GET("/someDataFromReader", func(c *gin.Context) {
		// 从指定地址获取图片
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")

		if err != nil || response.StatusCode != http.StatusOK {
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

	// 使用BasicAuth组件
	// gin.Accounts是map[string]string的一种快捷方式
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由BasicAuth中间设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Println(user)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "No SECRET: ("})
		}
	})

	r.GET("/someGet", getting)
	r.POST("/somePost", posting)
	return r

}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "getting ok.")
}

func posting(c *gin.Context) {
	var user struct {
		Name string `json:name`
		Age  string `json:age`
	}
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusUnauthorized, "err")
	}
	name := user.Name
	age := user.Age
	c.String(http.StatusOK, name+","+age+", posting ok.")
}
