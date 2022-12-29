package GinDemo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"Sex"`
	Age      int    `json:"Age"`
}

func GinInit() *gin.Engine {
	r := gin.Default()
	r.POST("/user/create", CreateUser)
	return r
}

func CreateUser(c *gin.Context) {
	var user []User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, v := range user {
		fmt.Println(v)
	}
	c.JSON(http.StatusOK, gin.H{"ret_info": "create user successfully."})
}
