package routers

import (
	"godemo/internal/goweb/gogin/proxy/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSshRouters(group *gin.RouterGroup) {
	sshGroup := group.Group("")
	sshGroup.POST("ssh", handlers.ExecSshCommand)
	sshGroup.GET("terminal", handlers.Terminal)
}
